// Copyright (c) 2026 Dedalus Labs, Inc. All rights reserved.

package internal_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestAccMachineShapeChangesRequireReplacement(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	api := &machineContract{t: t}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)
	cli := newTerraformCLI(t, terraform, bin, server.URL)
	cli.config(10)
	cli.run("apply", "-auto-approve", "-input=false")
	cli.noChanges()
	for _, test := range []struct {
		name    string
		cpu     int
		memory  int
		actions []string
	}{
		{name: "unchanged", cpu: 1, memory: 2048, actions: []string{"no-op"}},
		{name: "cpu", cpu: 2, memory: 2048, actions: []string{"delete", "create"}},
		{name: "memory", cpu: 1, memory: 4096, actions: []string{"delete", "create"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			subCLI := *cli
			subCLI.t = t
			cli := &subCLI
			cli.resource(fmt.Sprintf(`resource "dedalus_machine" "test" {
  vcpu = %d
  memory_mib = %d
  storage_gib = 10
  autosleep = "300s"
}
`, test.cpu, test.memory))
			cli.run("plan", "-input=false", "-out=shape.tfplan")
			var plan struct {
				Changes []struct {
					Address string `json:"address"`
					Change  struct {
						Actions []string `json:"actions"`
					} `json:"change"`
				} `json:"resource_changes"`
			}
			if err := json.Unmarshal([]byte(cli.run("show", "-json", "shape.tfplan")), &plan); err != nil {
				t.Fatal(err)
			}
			if len(plan.Changes) != 1 || plan.Changes[0].Address != "dedalus_machine.test" {
				t.Fatal("plan must contain exactly the test machine")
			}
			if got := plan.Changes[0].Change.Actions; !slices.Equal(got, test.actions) {
				t.Fatalf("shape plan actions: got %v, want %v", got, test.actions)
			}
		})
	}
	api.mu.Lock()
	defer api.mu.Unlock()
	if !api.exists || api.creates != 1 || api.updates != 0 || api.deletes != 0 {
		t.Fatalf("planning mutated the machine: exists=%t creates=%d updates=%d deletes=%d", api.exists, api.creates, api.updates, api.deletes)
	}
}
