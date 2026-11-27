// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachinePortDataSourceModel struct {
	ID             types.String      `tfsdk:"id" path:"port_id,computed"`
	PortID         types.String      `tfsdk:"port_id" path:"port_id,required"`
	MachineID      types.String      `tfsdk:"machine_id" path:"machine_id,required"`
	ApplicationURL types.String      `tfsdk:"application_url" json:"application_url,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode      types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage   types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt      timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Port           types.Int64       `tfsdk:"port" json:"port,computed"`
	Protocol       types.String      `tfsdk:"protocol" json:"protocol,computed"`
	ReadyAt        timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs   types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	URL            types.String      `tfsdk:"url" json:"url,computed"`
	Visibility     types.String      `tfsdk:"visibility" json:"visibility,computed"`
	Timeouts       timeouts.Value    `tfsdk:"timeouts"`
}

func (m *MachinePortDataSourceModel) toReadParams(_ context.Context) (params dedalus.MachinePortGetParams, diags diag.Diagnostics) {
	params = dedalus.MachinePortGetParams{
		MachineID: m.MachineID.ValueString(),
		PortID:    m.PortID.ValueString(),
	}

	return
}
