// Copyright (c) 2026 Dedalus Labs, Inc. All rights reserved.

package internal_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// These tests use the real Terraform CLI and a contract server, not cloud
// resources. The opt-in keeps ordinary Go tests independent of Terraform.
func TestAccMachineLifecycle(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	for _, test := range []string{"lifecycle", "destroy", "remote_disappearance", "remote_error"} {
		t.Run(test, func(t *testing.T) {
			api := &machineContract{t: t}
			server := httptest.NewServer(api)
			t.Cleanup(server.Close)
			cli := newTerraformCLI(t, terraform, bin, server.URL)
			cli.config(10)
			cli.run("apply", "-auto-approve", "-input=false")
			cli.noChanges()
			if test == "destroy" {
				cli.run("destroy", "-auto-approve", "-input=false")
				if resources := strings.TrimSpace(cli.run("state", "list")); resources != "" {
					t.Fatalf("state after destroy: %s", resources)
				}
				api.mu.Lock()
				defer api.mu.Unlock()
				if api.exists || api.deletes != 1 {
					t.Fatalf("destroy: exists=%t deletes=%d", api.exists, api.deletes)
				}
				return
			}
			if test == "remote_error" {
				api.mu.Lock()
				api.forbidden = true
				api.mu.Unlock()
				output, err := cli.command("plan", "-input=false")
				if err == nil {
					t.Fatalf("remote refusal: err=%v output=%s", err, output)
				}
				api.mu.Lock()
				refusals := api.refusals
				api.mu.Unlock()
				if refusals != 1 {
					t.Fatalf("remote refusals: got %d, want 1", refusals)
				}
				if resources := strings.TrimSpace(cli.run("state", "list")); resources != "dedalus_machine.test" {
					t.Fatalf("state after remote refusal: %s", resources)
				}
				return
			}
			if test == "remote_disappearance" {
				api.mu.Lock()
				api.exists = false
				api.mu.Unlock()
				cli.run("apply", "-refresh-only", "-auto-approve", "-input=false")
				if resources := strings.TrimSpace(cli.run("state", "list")); resources != "" {
					t.Fatalf("state after remote disappearance: %s", resources)
				}
				return
			}
			cli.config(20)
			cli.run("apply", "-auto-approve", "-input=false")
			cli.noChanges()
			cli.run("state", "rm", "dedalus_machine.test")
			cli.run("import", "-input=false", "dedalus_machine.test", "00000000-0000-4000-8000-000000000001")
			cli.noChanges()
			cli.run("destroy", "-auto-approve", "-input=false")
			api.mu.Lock()
			defer api.mu.Unlock()
			if api.exists || api.creates != 1 || api.updates != 1 || api.deletes != 1 {
				t.Fatalf("lifecycle: exists=%t creates=%d updates=%d deletes=%d", api.exists, api.creates, api.updates, api.deletes)
			}
		})
	}
}

func buildTerraformProvider(t *testing.T) (string, string) {
	t.Helper()
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("set TF_ACC=1 to run Terraform CLI contract tests")
	}
	terraform, err := exec.LookPath("terraform")
	if err != nil {
		t.Fatal(err)
	}
	bin := t.TempDir()
	build := exec.CommandContext(t.Context(), "go", "build", "-o", filepath.Join(bin, "terraform-provider-dedalus"), "..")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("provider build: %v\n%s", err, output)
	}
	return terraform, bin
}

type terraformCLI struct {
	t       *testing.T
	binary  string
	dir     string
	env     []string
	baseURL string
}

func newTerraformCLI(t *testing.T, binary, providerDir, baseURL string) *terraformCLI {
	t.Helper()
	dir := t.TempDir()
	rc := filepath.Join(dir, "terraform.rc")
	config := fmt.Sprintf("provider_installation {\n dev_overrides {\n  \"dedalus-labs/dedalus\" = %q\n }\n}\n", filepath.ToSlash(providerDir))
	if err := os.WriteFile(rc, []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	var env []string
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(name, "TF_") || strings.HasPrefix(name, "DEDALUS_") {
			continue
		}
		env = append(env, entry)
	}
	env = append(env, "TF_CLI_CONFIG_FILE="+rc, "TF_IN_AUTOMATION=1", "CHECKPOINT_DISABLE=1")
	return &terraformCLI{t: t, binary: binary, dir: dir, env: env, baseURL: baseURL}
}

func (cli *terraformCLI) config(storage int) {
	cli.t.Helper()
	cli.resource(fmt.Sprintf(`resource "dedalus_machine" "test" {
  vcpu = 1
  memory_mib = 2048
  storage_gib = %d
  autosleep = "300s"
}
`, storage))
}

func (cli *terraformCLI) resource(resource string) {
	cli.t.Helper()
	config := fmt.Sprintf(`terraform {
  required_providers {
    dedalus = { source = "dedalus-labs/dedalus" }
  }
}
provider "dedalus" {
  base_url = %q
  api_key = "contract-key"
  dedalus_org_id = "contract-org"
}
%s
`, cli.baseURL, resource)
	if err := os.WriteFile(filepath.Join(cli.dir, "main.tf"), []byte(config), 0600); err != nil {
		cli.t.Fatal(err)
	}
}

func (cli *terraformCLI) command(args ...string) ([]byte, error) {
	cmd := exec.CommandContext(cli.t.Context(), cli.binary, append(args, "-no-color")...)
	cmd.Dir = cli.dir
	cmd.Env = cli.env
	return cmd.CombinedOutput()
}

func (cli *terraformCLI) run(args ...string) string {
	cli.t.Helper()
	output, err := cli.command(args...)
	if err != nil {
		cli.t.Fatalf("terraform %s: %v\n%s", strings.Join(args, " "), err, output)
	}
	return string(output)
}

func (cli *terraformCLI) noChanges() {
	cli.t.Helper()
	cli.run("plan", "-detailed-exitcode", "-input=false")
}

type machineMutation struct {
	VCPU       *float64 `json:"vcpu"`
	MemoryMiB  *int64   `json:"memory_mib"`
	StorageGiB *int64   `json:"storage_gib"`
	Autosleep  *string  `json:"autosleep"`
}

// Mutation responses and GET responses deliberately differ, as the public
// LifecycleResponse and MachineDetailResponse OpenAPI schemas require.
type machineReply struct {
	MachineID        string         `json:"machine_id"`
	VCPU             float64        `json:"vcpu"`
	MemoryMiB        int64          `json:"memory_mib"`
	StorageGiB       int64          `json:"storage_gib"`
	AutosleepSeconds int64          `json:"autosleep_seconds"`
	DesiredState     string         `json:"desired_state"`
	Phase            string         `json:"phase,omitempty"`
	Status           *machineStatus `json:"status,omitempty"`
}

type machineStatus struct {
	Autoresize struct {
		Memory *struct{} `json:"memory"`
		CPU    *struct{} `json:"cpu"`
	} `json:"autoresize"`
	Phase            string `json:"phase"`
	Reason           string `json:"reason"`
	Retryable        bool   `json:"retryable"`
	Revision         string `json:"revision"`
	LastTransitionAt string `json:"last_transition_at"`
	LastProgressAt   string `json:"last_progress_at"`
}

type machineContract struct {
	t                         *testing.T
	mu                        sync.Mutex
	exists                    bool
	forbidden                 bool
	machine                   machineReply
	creates, updates, deletes int
	refusals                  int
}

func (api *machineContract) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Authorization") != "Bearer contract-key" || r.Header.Get("X-Dedalus-Org-Id") != "contract-org" {
		api.refuse(w, http.StatusUnauthorized, "provider authentication headers missing")
		return
	}
	if r.Method == http.MethodPost && r.URL.Path == "/v1/machines" {
		api.create(w, r)
		return
	}
	if r.URL.Path != "/v1/machines/00000000-0000-4000-8000-000000000001" || !api.exists {
		api.refuse(w, http.StatusNotFound, "machine not found")
		return
	}
	switch r.Method {
	case http.MethodGet:
		if api.forbidden {
			api.refusals++
			api.refuse(w, http.StatusForbidden, "permission denied")
			return
		}
		reply := api.machine
		reply.Phase = ""
		reply.Status = &machineStatus{Phase: "running", Reason: "ready", Revision: strconv.Itoa(api.updates + 1), LastTransitionAt: "2026-01-01T00:00:00Z", LastProgressAt: "2026-01-01T00:00:00Z"}
		api.write(w, reply)
	case http.MethodPatch:
		var update machineMutation
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&update); err != nil || update.StorageGiB == nil || update.VCPU != nil || update.MemoryMiB != nil || update.Autosleep != nil {
			api.refuse(w, http.StatusBadRequest, "update must contain only changed storage_gib")
			return
		}
		api.machine.StorageGiB = *update.StorageGiB
		api.updates++
		api.write(w, api.machine)
	case http.MethodDelete:
		api.exists = false
		api.deletes++
		reply := api.machine
		reply.Phase, reply.DesiredState = "destroyed", "destroyed"
		api.write(w, reply)
	default:
		api.refuse(w, http.StatusMethodNotAllowed, "unexpected method")
	}
}

func (api *machineContract) create(w http.ResponseWriter, r *http.Request) {
	var create machineMutation
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&create); err != nil || create.VCPU == nil || create.MemoryMiB == nil || create.StorageGiB == nil || create.Autosleep == nil {
		api.refuse(w, http.StatusBadRequest, "create must contain configured shape and autosleep")
		return
	}
	api.machine = machineReply{MachineID: "00000000-0000-4000-8000-000000000001", VCPU: *create.VCPU, MemoryMiB: *create.MemoryMiB, StorageGiB: *create.StorageGiB, AutosleepSeconds: 300, DesiredState: "running", Phase: "running"}
	api.exists = true
	api.creates++
	api.write(w, api.machine)
}

func (api *machineContract) write(w http.ResponseWriter, response machineReply) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.t.Errorf("write API response: %v", err)
	}
}

func (api *machineContract) refuse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(struct {
		Code      string `json:"error_code"`
		Message   string `json:"message"`
		Retryable bool   `json:"retryable"`
	}{Code: "CONTRACT_ERROR", Message: message}); err != nil {
		api.t.Errorf("write API error: %v", err)
	}
}
