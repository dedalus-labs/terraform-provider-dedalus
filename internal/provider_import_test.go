// Copyright (c) 2026 Dedalus Labs, Inc. All rights reserved.

package internal_test

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestAccMachineImportPreservesRemoteAutosleep(t *testing.T) {
	terraform, bin := buildTerraformProvider(t)
	for _, test := range []struct {
		name    string
		seconds int64
		config  string
	}{
		{name: "disabled", seconds: 0, config: "never"},
		{name: "seconds", seconds: 90, config: "90s"},
		{name: "default", seconds: 300, config: "300s"},
	} {
		t.Run(test.name, func(t *testing.T) {
			api := &machineContract{
				t: t, exists: true,
				machine: machineReply{
					MachineID: "00000000-0000-4000-8000-000000000001",
					VCPU:      1, MemoryMiB: 2048, StorageGiB: 10,
					AutosleepSeconds: test.seconds, DesiredState: "running", Phase: "running",
				},
			}
			server := httptest.NewServer(api)
			t.Cleanup(server.Close)
			cli := newTerraformCLI(t, terraform, bin, server.URL)
			cli.resource(fmt.Sprintf(`resource "dedalus_machine" "test" {
  vcpu = 1
  memory_mib = 2048
  storage_gib = 10
  autosleep = %q
}
`, test.config))
			cli.run("import", "-input=false", "dedalus_machine.test", api.machine.MachineID)
			cli.noChanges()
		})
	}
}
