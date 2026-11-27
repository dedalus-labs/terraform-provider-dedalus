// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachinePortsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MachinePortsItemsDataSourceModel] `json:"items,computed"`
}

type MachinePortsDataSourceModel struct {
	MachineID types.String                                                   `tfsdk:"machine_id" path:"machine_id,required"`
	MaxItems  types.Int64                                                    `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[MachinePortsItemsDataSourceModel] `tfsdk:"items"`
	Timeouts  timeouts.Value                                                 `tfsdk:"timeouts"`
}

func (m *MachinePortsDataSourceModel) toListParams(_ context.Context) (params dedalus.MachinePortListParams, diags diag.Diagnostics) {
	params = dedalus.MachinePortListParams{
		MachineID: m.MachineID.ValueString(),
	}

	return
}

type MachinePortsItemsDataSourceModel struct {
	ID             types.String      `tfsdk:"id" json:"port_id,computed"`
	ApplicationURL types.String      `tfsdk:"application_url" json:"application_url,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	MachineID      types.String      `tfsdk:"machine_id" json:"machine_id,computed"`
	Port           types.Int64       `tfsdk:"port" json:"port,computed"`
	PortID         types.String      `tfsdk:"port_id" json:"port_id,computed"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	Visibility     types.String      `tfsdk:"visibility" json:"visibility,computed"`
	ErrorCode      types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage   types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt      timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Protocol       types.String      `tfsdk:"protocol" json:"protocol,computed"`
	ReadyAt        timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs   types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	URL            types.String      `tfsdk:"url" json:"url,computed"`
}
