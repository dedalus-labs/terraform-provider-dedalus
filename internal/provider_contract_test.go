// @custom start
// Check the registered provider surface against the public Machines contract.

package internal_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestInvariantProviderSurfaceMatchesPublicContract(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	provider := internal.NewProvider("test")()

	resources := []string{}
	for _, newResource := range provider.Resources(ctx) {
		var metadata resource.MetadataResponse
		newResource().Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "dedalus"}, &metadata)
		resources = append(resources, metadata.TypeName)
	}
	slices.Sort(resources)
	wantResources := []string{"dedalus_machine", "dedalus_machine_execution", "dedalus_machine_ssh_session"}
	if !reflect.DeepEqual(resources, wantResources) {
		t.Errorf("resources: got %v, want %v", resources, wantResources)
	}

	sources := []string{}
	for _, newSource := range provider.DataSources(ctx) {
		var metadata datasource.MetadataResponse
		newSource().Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: "dedalus"}, &metadata)
		sources = append(sources, metadata.TypeName)
	}
	slices.Sort(sources)
	wantSources := []string{
		"dedalus_machine", "dedalus_machine_execution", "dedalus_machine_executions",
		"dedalus_machine_ssh_session", "dedalus_machine_ssh_sessions", "dedalus_machines",
	}
	if !reflect.DeepEqual(sources, wantSources) {
		t.Errorf("data sources: got %v, want %v", sources, wantSources)
	}
}

func TestInvariantMachineCreationAcceptsServerDefaults(t *testing.T) {
	t.Parallel()
	schema := machine.ResourceSchema(t.Context())
	for _, name := range []string{"autosleep", "memory_mib", "storage_gib", "vcpu"} {
		attribute := schema.Attributes[name]
		if attribute == nil {
			t.Fatalf("missing machine attribute %q", name)
		}
		if !attribute.IsOptional() || !attribute.IsComputed() || attribute.IsRequired() {
			t.Errorf("%s must be optional and computed", name)
		}
	}
}

func TestInvariantMachineListPreservesIdentityAndPhase(t *testing.T) {
	t.Parallel()
	var item machine.MachinesItemsDataSourceModel
	if err := apijson.UnmarshalComputed([]byte(`{"machine_id":"dm-test","phase":"running"}`), &item); err != nil {
		t.Fatal(err)
	}
	if item.ID.ValueString() != "dm-test" || item.MachineID.ValueString() != "dm-test" {
		t.Fatalf("machine identity: id=%q machine_id=%q", item.ID.ValueString(), item.MachineID.ValueString())
	}
	if item.Phase.ValueString() != "running" {
		t.Fatalf("phase: got %q, want running", item.Phase.ValueString())
	}
}

// @custom end
