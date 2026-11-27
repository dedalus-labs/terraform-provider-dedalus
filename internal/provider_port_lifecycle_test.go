// Copyright (c) 2026 Dedalus Labs, Inc. All rights reserved.

package internal_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"
)

const portMachineID = "00000000-0000-4000-8000-000000000001"
const portPath = "/v1/machines/" + portMachineID + "/ports"

var portPasswordPattern = regexp.MustCompile(`^[A-Za-z0-9]{8}$`)

func TestAccPortLifecycle(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	api := &portContract{t: t}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)
	cli := newTerraformCLI(t, terraform, bin, server.URL)
	configurePort(cli, "private", "")
	cli.run("apply", "-auto-approve", "-input=false")
	cli.noChanges()
	configurePort(cli, "public", "")
	cli.run("apply", "-auto-approve", "-input=false")
	cli.noChanges()
	cli.run("state", "rm", "dedalus_machine_port.test")
	cli.run("import", "-input=false", "dedalus_machine_port.test", portMachineID+"/port-contract")
	cli.noChanges()
	cli.run("destroy", "-auto-approve", "-input=false")
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.exists || api.creates != 1 || api.updates != 1 || api.deletes != 1 {
		t.Fatalf("port lifecycle: exists=%t creates=%d updates=%d deletes=%d", api.exists, api.creates, api.updates, api.deletes)
	}
}

func TestAccPortPassword(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	for _, test := range []string{"sensitive", "refresh", "rotation", "private", "import"} {
		t.Run(test, func(t *testing.T) {
			api := &portContract{t: t}
			server := httptest.NewServer(api)
			t.Cleanup(server.Close)
			cli := newTerraformCLI(t, terraform, bin, server.URL)
			const password = "Test7d29"
			configurePort(cli, "key_required", password)
			output := cli.run("apply", "-auto-approve", "-input=false")
			if strings.Contains(output, password) {
				t.Error("apply output contains the port password")
			}
			if test == "sensitive" {
				if strings.Contains(cli.run("show"), password) {
					t.Error("human-readable state contains the port password")
				}
				return
			}
			if test == "import" {
				cli.run("state", "rm", "dedalus_machine_port.test")
				cli.run("import", "-input=false", "dedalus_machine_port.test", portMachineID+"/port-contract")
				if portPasswordFromState(t, cli) != nil {
					t.Fatal("import reconstructed a write-only password")
				}
				cli.run("apply", "-auto-approve", "-input=false")
				cli.noChanges()
				api.mu.Lock()
				defer api.mu.Unlock()
				if api.creates != 1 || api.updates != 1 {
					t.Fatalf("import must reapply the configured password without replacement: creates=%d updates=%d", api.creates, api.updates)
				}
				return
			}
			cli.run("apply", "-refresh-only", "-auto-approve", "-input=false")
			if value := portPasswordFromState(t, cli); value == nil || *value != password {
				t.Fatal("refresh lost the configured write-only password")
			}
			cli.noChanges()
			switch test {
			case "rotation":
				configurePort(cli, "key_required", "Next8c41")
				cli.run("apply", "-auto-approve", "-input=false")
				if value := portPasswordFromState(t, cli); value == nil || *value != "Next8c41" {
					t.Fatal("rotation did not retain the configured password")
				}
			case "private":
				configurePort(cli, "private", "")
				cli.run("apply", "-auto-approve", "-input=false")
				if portPasswordFromState(t, cli) != nil {
					t.Fatal("private port retained an unconfigured password")
				}
			}
			cli.noChanges()
			cli.run("destroy", "-auto-approve", "-input=false")
			api.mu.Lock()
			defer api.mu.Unlock()
			if api.creates != 1 || api.deletes != 1 {
				t.Fatalf("password operation replaced the port: creates=%d deletes=%d", api.creates, api.deletes)
			}
		})
	}
}

func TestAccPortRejectsInvalidPasswords(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	for _, test := range []struct {
		name       string
		visibility string
		password   string
	}{
		{name: "required", visibility: "key_required"},
		{name: "private", visibility: "private", password: "Test7d29"},
		{name: "public", visibility: "public", password: "Test7d29"},
		{name: "length", visibility: "key_required", password: "short"},
		{name: "alphabet", visibility: "key_required", password: "Test7d2!"},
	} {
		t.Run(test.name, func(t *testing.T) {
			api := &portContract{t: t}
			server := httptest.NewServer(api)
			t.Cleanup(server.Close)
			cli := newTerraformCLI(t, terraform, bin, server.URL)
			configurePort(cli, test.visibility, test.password)
			if output, err := cli.command("apply", "-auto-approve", "-input=false"); err == nil {
				t.Fatalf("invalid password configuration was accepted: %s", output)
			}
			api.mu.Lock()
			defer api.mu.Unlock()
			if api.exists || api.creates != 0 {
				t.Fatalf("rejected password created a port: exists=%t creates=%d", api.exists, api.creates)
			}
		})
	}
}

func configurePort(cli *terraformCLI, visibility, password string) {
	cli.t.Helper()
	passwordAttribute := ""
	if password != "" {
		passwordAttribute = fmt.Sprintf("  password = %q\n", password)
	}
	cli.resource(fmt.Sprintf(`resource "dedalus_machine_port" "test" {
  machine_id = %q
  port = 8080
  protocol = "http"
  visibility = %q
%s}
`, portMachineID, visibility, passwordAttribute))
}

func portPasswordFromState(t *testing.T, cli *terraformCLI) *string {
	t.Helper()
	var state struct {
		Resources []struct {
			Type      string `json:"type"`
			Instances []struct {
				Attributes struct {
					Password *string `json:"password"`
				} `json:"attributes"`
			} `json:"instances"`
		} `json:"resources"`
	}
	if err := json.Unmarshal([]byte(cli.run("state", "pull")), &state); err != nil {
		t.Fatal(err)
	}
	if len(state.Resources) != 1 || state.Resources[0].Type != "dedalus_machine_port" || len(state.Resources[0].Instances) != 1 {
		t.Fatal("state must contain exactly one managed port")
	}
	return state.Resources[0].Instances[0].Attributes.Password
}

type portMutation struct {
	Port       *int64          `json:"port"`
	Protocol   string          `json:"protocol"`
	Visibility string          `json:"visibility"`
	Password   json.RawMessage `json:"password"`
}

func (mutation portMutation) validPassword(required bool) bool {
	if mutation.Visibility != "key_required" {
		return len(mutation.Password) == 0
	}
	if !required && len(mutation.Password) == 0 {
		return true
	}
	var password string
	return json.Unmarshal(mutation.Password, &password) == nil && portPasswordPattern.MatchString(password)
}

// The current PortResponse never returns the password, including at creation.
type portReply struct {
	MachineID      string `json:"machine_id"`
	PortID         string `json:"port_id"`
	Port           int64  `json:"port"`
	Protocol       string `json:"protocol"`
	Visibility     string `json:"visibility"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	ApplicationURL string `json:"application_url"`
}

type portContract struct {
	t                         *testing.T
	mu                        sync.Mutex
	exists                    bool
	port                      portReply
	creates, updates, deletes int
}

func (api *portContract) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Authorization") != "Bearer contract-key" {
		http.Error(w, "authentication missing", http.StatusUnauthorized)
		return
	}
	if r.Method == http.MethodPost && r.URL.Path == portPath {
		var create portMutation
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&create); err != nil || create.Port == nil || *create.Port != 8080 || create.Protocol != "http" || !create.validPassword(true) {
			http.Error(w, "invalid create body", http.StatusBadRequest)
			return
		}
		api.port = portReply{
			MachineID: portMachineID, PortID: "port-contract", Port: *create.Port,
			Protocol: "http", Visibility: create.Visibility, Status: "ready",
			CreatedAt: "2026-01-01T00:00:00Z", ApplicationURL: "https://app.example.test",
		}
		api.exists = true
		api.creates++
	} else {
		if r.URL.Path != portPath+"/port-contract" || !api.exists {
			http.Error(w, "port not found", http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
		case http.MethodPatch:
			var update portMutation
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&update); err != nil || update.Port != nil || update.Protocol != "" || (update.Visibility != "public" && update.Visibility != "private" && update.Visibility != "key_required") || !update.validPassword(false) {
				api.t.Errorf("PATCH contract: visibility=%q port_set=%t protocol=%q password_present=%t password_null=%t", update.Visibility, update.Port != nil, update.Protocol, len(update.Password) > 0, string(update.Password) == "null")
				http.Error(w, "update requires visibility and a password only for key_required", http.StatusBadRequest)
				return
			}
			api.port.Visibility = update.Visibility
			api.updates++
		case http.MethodDelete:
			api.port.Status = "closed"
			api.exists = false
			api.deletes++
		default:
			http.Error(w, "unexpected method", http.StatusMethodNotAllowed)
			return
		}
	}
	if err := json.NewEncoder(w).Encode(api.port); err != nil {
		api.t.Errorf("write port response: %v", err)
	}
}
