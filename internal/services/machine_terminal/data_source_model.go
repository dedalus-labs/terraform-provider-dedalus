// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineTerminalDataSourceModel struct {
	ID           types.String      `tfsdk:"id" path:"terminal_id,computed"`
	TerminalID   types.String      `tfsdk:"terminal_id" path:"terminal_id,required"`
	MachineID    types.String      `tfsdk:"machine_id" path:"machine_id,required"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode    types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Height       types.Int64       `tfsdk:"height" json:"height,computed"`
	Protocol     types.String      `tfsdk:"protocol" json:"protocol,computed"`
	ReadyAt      timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status       types.String      `tfsdk:"status" json:"status,computed"`
	StreamURL    types.String      `tfsdk:"stream_url" json:"stream_url,computed"`
	Width        types.Int64       `tfsdk:"width" json:"width,computed"`
	Timeouts     timeouts.Value    `tfsdk:"timeouts"`
}

func (m *MachineTerminalDataSourceModel) toReadParams(_ context.Context) (params dedalus.MachineTerminalGetParams, diags diag.Diagnostics) {
	params = dedalus.MachineTerminalGetParams{
		MachineID:  m.MachineID.ValueString(),
		TerminalID: m.TerminalID.ValueString(),
	}

	return
}
