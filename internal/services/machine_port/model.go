// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachinePortModel struct {
	ID             types.String      `tfsdk:"id" json:"-,computed"`
	PortID         types.String      `tfsdk:"port_id" json:"port_id,computed"`
	MachineID      types.String      `tfsdk:"machine_id" path:"machine_id,required"`
	Port           types.Int64       `tfsdk:"port" json:"port,required"`
	Expires        types.Int64       `tfsdk:"expires" json:"expires,optional,no_refresh"`
	Protocol       types.String      `tfsdk:"protocol" json:"protocol,computed_optional"`
	Password       types.String      `tfsdk:"password" json:"password,optional,no_refresh"`
	Visibility     types.String      `tfsdk:"visibility" json:"visibility,computed_optional"`
	ApplicationURL types.String      `tfsdk:"application_url" json:"application_url,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode      types.String      `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage   types.String      `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt      timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	ReadyAt        timetypes.RFC3339 `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs   types.Int64       `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	URL            types.String      `tfsdk:"url" json:"url,computed"`
	Timeouts       timeouts.Value    `tfsdk:"timeouts"`
}

func (m MachinePortModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MachinePortModel) MarshalJSONForUpdate(state MachinePortModel) (data []byte, err error) {
	// @custom start
	// Port updates require visibility and omit an unset password. This endpoint
	// accepts application/json, not JSON Merge Patch deletion semantics.
	return apijson.MarshalRoot(struct {
		Visibility types.String `json:"visibility,required"`
		Password   types.String `json:"password,optional"`
	}{Visibility: m.Visibility, Password: m.Password})
	// @custom end
}
