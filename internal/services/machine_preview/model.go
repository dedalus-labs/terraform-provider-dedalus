// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_preview

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachinePreviewModel struct {
	ID           types.String      `tfsdk:"id" json:"-,computed"`
	PreviewID    types.String      `tfsdk:"preview_id" json:"preview_id,computed"`
	MachineID    types.String      `tfsdk:"machine_id" path:"machine_id,required"`
	Port         types.Int64       `tfsdk:"port" json:"port,required"`
	Protocol     types.String      `tfsdk:"protocol" json:"protocol,optional"`
	Visibility   types.String      `tfsdk:"visibility" json:"visibility,optional"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode    types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	ReadyAt      timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status       types.String      `tfsdk:"status" json:"status,computed"`
	URL          types.String      `tfsdk:"url" json:"url,computed"`
	Timeouts     timeouts.Value    `tfsdk:"timeouts"`
}

func (m MachinePreviewModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MachinePreviewModel) MarshalJSONForUpdate(state MachinePreviewModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
