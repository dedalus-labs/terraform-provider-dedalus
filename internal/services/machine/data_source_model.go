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

type MachineDataSourceModel struct {
	ID           types.String                                           `tfsdk:"id" path:"machine_id,computed"`
	MachineID    types.String                                           `tfsdk:"machine_id" path:"machine_id,required"`
	DesiredState types.String                                           `tfsdk:"desired_state" json:"desired_state,computed"`
	MemoryMiB    types.Int64                                            `tfsdk:"memory_mib" json:"memory_mib,computed"`
	StorageGiB   types.Int64                                            `tfsdk:"storage_gib" json:"storage_gib,computed"`
	VCPU         types.Float64                                          `tfsdk:"vcpu" json:"vcpu,computed"`
	Status       customfield.NestedObject[MachineStatusDataSourceModel] `tfsdk:"status" json:"status,computed"`
	Timeouts     timeouts.Value                                         `tfsdk:"timeouts"`
}

func (m *MachineDataSourceModel) toReadParams(_ context.Context) (params dedalus.MachineGetParams, diags diag.Diagnostics) {
	params = dedalus.MachineGetParams{
		MachineID: m.MachineID.ValueString(),
	}

	return
}

type MachineStatusDataSourceModel struct {
	LastProgressAt   timetypes.RFC3339 `tfsdk:"last_progress_at" json:"last_progress_at,computed" format:"date-time"`
	LastTransitionAt timetypes.RFC3339 `tfsdk:"last_transition_at" json:"last_transition_at,computed" format:"date-time"`
	Phase            types.String      `tfsdk:"phase" json:"phase,computed"`
	Reason           types.String      `tfsdk:"reason" json:"reason,computed"`
	Retryable        types.Bool        `tfsdk:"retryable" json:"retryable,computed"`
	Revision         types.String      `tfsdk:"revision" json:"revision,computed"`
	LastError        types.String      `tfsdk:"last_error" json:"last_error,computed"`
}
