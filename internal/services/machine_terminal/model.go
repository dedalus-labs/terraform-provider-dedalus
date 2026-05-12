// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineTerminalModel struct {
	ID           types.String             `tfsdk:"id" json:"-,computed"`
	TerminalID   types.String             `tfsdk:"terminal_id" json:"terminal_id,computed"`
	MachineID    types.String             `tfsdk:"machine_id" path:"machine_id,required"`
	Height       types.Int64              `tfsdk:"height" json:"height,required"`
	Width        types.Int64              `tfsdk:"width" json:"width,required"`
	Cwd          types.String             `tfsdk:"cwd" json:"cwd,optional,no_refresh"`
	Shell        types.String             `tfsdk:"shell" json:"shell,optional,no_refresh"`
	Env          *map[string]types.String `tfsdk:"env" json:"env,optional,no_refresh"`
	CreatedAt    timetypes.RFC3339        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode    types.String             `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage types.String             `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt    timetypes.RFC3339        `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Protocol     types.String             `tfsdk:"protocol" json:"protocol,computed"`
	ReadyAt      timetypes.RFC3339        `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs types.Int64              `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status       types.String             `tfsdk:"status" json:"status,computed"`
	StreamURL    types.String             `tfsdk:"stream_url" json:"stream_url,computed"`
	Timeouts     timeouts.Value           `tfsdk:"timeouts"`
}

func (m MachineTerminalModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MachineTerminalModel) MarshalJSONForUpdate(state MachineTerminalModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
