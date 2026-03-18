// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkspaceDataSourceModel struct {
	ID           types.String                                             `tfsdk:"id" path:"workspace_id,computed"`
	WorkspaceID  types.String                                             `tfsdk:"workspace_id" path:"workspace_id,required"`
	DesiredState types.String                                             `tfsdk:"desired_state" json:"desired_state,computed"`
	ImageVersion types.String                                             `tfsdk:"image_version" json:"image_version,computed"`
	MemoryMiB    types.Int64                                              `tfsdk:"memory_mib" json:"memory_mib,computed"`
	Schema       types.String                                             `tfsdk:"schema" json:"$schema,computed"`
	StorageGiB   types.Int64                                              `tfsdk:"storage_gib" json:"storage_gib,computed"`
	VCPU         types.Float64                                            `tfsdk:"vcpu" json:"vcpu,computed"`
	Status       customfield.NestedObject[WorkspaceStatusDataSourceModel] `tfsdk:"status" json:"status,computed"`
	Timeouts     timeouts.Value                                           `tfsdk:"timeouts"`
}

type WorkspaceStatusDataSourceModel struct {
	LastProgressAt    timetypes.RFC3339 `tfsdk:"last_progress_at" json:"last_progress_at,computed" format:"date-time"`
	LastTransitionAt  timetypes.RFC3339 `tfsdk:"last_transition_at" json:"last_transition_at,computed" format:"date-time"`
	ObservedRevision  types.String      `tfsdk:"observed_revision" json:"observed_revision,computed"`
	Phase             types.String      `tfsdk:"phase" json:"phase,computed"`
	Reason            types.String      `tfsdk:"reason" json:"reason,computed"`
	Retryable         types.Bool        `tfsdk:"retryable" json:"retryable,computed"`
	Revision          types.String      `tfsdk:"revision" json:"revision,computed"`
	AssignedHost      types.String      `tfsdk:"assigned_host" json:"assigned_host,computed"`
	LastError         types.String      `tfsdk:"last_error" json:"last_error,computed"`
	MemoryAssignedMiB types.Int64       `tfsdk:"memory_assigned_mib" json:"memory_assigned_mib,computed"`
	MemoryResizeState types.String      `tfsdk:"memory_resize_state" json:"memory_resize_state,computed"`
	MemoryTargetMiB   types.Int64       `tfsdk:"memory_target_mib" json:"memory_target_mib,computed"`
}
