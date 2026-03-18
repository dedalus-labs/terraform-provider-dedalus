// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkspaceModel struct {
	ID           types.String                                   `tfsdk:"id" json:"-,computed"`
	WorkspaceID  types.String                                   `tfsdk:"workspace_id" json:"workspace_id,computed"`
	MemoryMiB    types.Int64                                    `tfsdk:"memory_mib" json:"memory_mib,required"`
	StorageGiB   types.Int64                                    `tfsdk:"storage_gib" json:"storage_gib,required"`
	VCPU         types.Float64                                  `tfsdk:"vcpu" json:"vcpu,required"`
	DesiredState types.String                                   `tfsdk:"desired_state" json:"desired_state,computed"`
	Schema       types.String                                   `tfsdk:"schema" json:"$schema,computed"`
	Status       customfield.NestedObject[WorkspaceStatusModel] `tfsdk:"status" json:"status,computed"`
	Timeouts     timeouts.Value                                 `tfsdk:"timeouts"`
}

func (m WorkspaceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkspaceModel) MarshalJSONForUpdate(state WorkspaceModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type WorkspaceStatusModel struct {
	LastProgressAt   timetypes.RFC3339 `tfsdk:"last_progress_at" json:"last_progress_at,computed" format:"date-time"`
	LastTransitionAt timetypes.RFC3339 `tfsdk:"last_transition_at" json:"last_transition_at,computed" format:"date-time"`
	Phase            types.String      `tfsdk:"phase" json:"phase,computed"`
	Reason           types.String      `tfsdk:"reason" json:"reason,computed"`
	Retryable        types.Bool        `tfsdk:"retryable" json:"retryable,computed"`
	Revision         types.String      `tfsdk:"revision" json:"revision,computed"`
	LastError        types.String      `tfsdk:"last_error" json:"last_error,computed"`
}
