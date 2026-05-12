// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineTerminalsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MachineTerminalsItemsDataSourceModel] `json:"items,computed"`
}

type MachineTerminalsDataSourceModel struct {
	MachineID types.String                                                       `tfsdk:"machine_id" path:"machine_id,required"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[MachineTerminalsItemsDataSourceModel] `tfsdk:"items"`
	Timeouts  timeouts.Value                                                     `tfsdk:"timeouts"`
}

func (m *MachineTerminalsDataSourceModel) toListParams(_ context.Context) (params dedalus.MachineTerminalListParams, diags diag.Diagnostics) {
	params = dedalus.MachineTerminalListParams{
		MachineID: m.MachineID.ValueString(),
	}

	return
}

type MachineTerminalsItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"terminal_id,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Height       types.Int64       `tfsdk:"height" json:"height,computed"`
	MachineID    types.String      `tfsdk:"machine_id" json:"machine_id,computed"`
	Status       types.String      `tfsdk:"status" json:"status,computed"`
	TerminalID   types.String      `tfsdk:"terminal_id" json:"terminal_id,computed"`
	Width        types.Int64       `tfsdk:"width" json:"width,computed"`
	ErrorCode    types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Protocol     types.String      `tfsdk:"protocol" json:"protocol,computed"`
	ReadyAt      timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	StreamURL    types.String      `tfsdk:"stream_url" json:"stream_url,computed"`
}
