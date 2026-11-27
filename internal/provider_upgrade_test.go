// Copyright (c) 2026 Dedalus Labs, Inc. All rights reserved.

package internal_test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The previous provider writes the state before the candidate ever opens it.
// Supplying the previous binary makes the version boundary explicit, and
// avoids changing the checkout or relying on a hand-written state fixture.
func TestAccMachineStateUpgrade(t *testing.T) {
	previous := os.Getenv("TF_ACC_PREVIOUS_PROVIDER")
	if previous == "" {
		t.Skip("set TF_ACC_PREVIOUS_PROVIDER to test upgrading existing state")
	}
	terraform, candidate := buildTerraformProvider(t)
	previousDir := t.TempDir()
	binary, err := os.ReadFile(previous)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(previousDir, "terraform-provider-dedalus"), binary, 0700); err != nil {
		t.Fatal(err)
	}
	api := &machineContract{t: t}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)
	cli := newTerraformCLI(t, terraform, previousDir, server.URL)
	configure := func(storage int) {
		cli.resource(fmt.Sprintf(`resource "dedalus_machine" "test" {
  vcpu = 1
  memory_mib = 2048
  storage_gib = %d
  autosleep = "300s"
}
output "machine_id" { value = dedalus_machine.test.machine_id }
`, storage))
	}
	configure(10)
	cli.run("apply", "-auto-approve", "-input=false")
	cli.noChanges()
	id := strings.TrimSpace(cli.run("output", "-raw", "machine_id"))
	if id != "00000000-0000-4000-8000-000000000001" {
		t.Fatalf("previous provider machine ID: got %q", id)
	}
	config := fmt.Sprintf("provider_installation {\n dev_overrides {\n  \"dedalus-labs/dedalus\" = %q\n }\n}\n", filepath.ToSlash(candidate))
	if err := os.WriteFile(filepath.Join(cli.dir, "terraform.rc"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	cli.noChanges()
	configure(20)
	cli.run("apply", "-auto-approve", "-input=false")
	cli.noChanges()
	if got := strings.TrimSpace(cli.run("output", "-raw", "machine_id")); got != id {
		t.Fatalf("upgraded provider machine ID: got %q, want %q", got, id)
	}
	cli.run("destroy", "-auto-approve", "-input=false")
	api.mu.Lock()
	defer api.mu.Unlock()
	if api.exists || api.creates != 1 || api.updates != 1 || api.deletes != 1 {
		t.Fatalf("upgrade lifecycle: exists=%t creates=%d updates=%d deletes=%d", api.exists, api.creates, api.updates, api.deletes)
	}
}
