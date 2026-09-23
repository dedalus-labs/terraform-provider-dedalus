// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachinesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MachinesItemsDataSourceModel] `json:"items,computed"`
}

type MachinesDataSourceModel struct {
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MachinesItemsDataSourceModel] `tfsdk:"items"`
	Timeouts timeouts.Value                                             `tfsdk:"timeouts"`
}

func (m *MachinesDataSourceModel) toListParams(_ context.Context) (params dedalus.MachineListParams, diags diag.Diagnostics) {
	params = dedalus.MachineListParams{}

	return
}

type MachinesItemsDataSourceModel struct {
	ID               types.String      `tfsdk:"id" json:"machine_id,computed"`
	AutosleepSeconds types.Int64       `tfsdk:"autosleep_seconds" json:"autosleep_seconds,computed"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DesiredState     types.String      `tfsdk:"desired_state" json:"desired_state,computed"`
	MachineID        types.String      `tfsdk:"machine_id" json:"machine_id,computed"`
	MemoryMiB        types.Int64       `tfsdk:"memory_mib" json:"memory_mib,computed"`
	Phase            types.String      `tfsdk:"phase" json:"phase,computed"`
	StorageGiB       types.Int64       `tfsdk:"storage_gib" json:"storage_gib,computed"`
	VCPU             types.Float64     `tfsdk:"vcpu" json:"vcpu,computed"`
}
