// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineModel struct {
	ID               types.String                                 `tfsdk:"id" json:"-,computed"`
	MachineID        types.String                                 `tfsdk:"machine_id" json:"machine_id,computed"`
	ImageID          types.String                                 `tfsdk:"image_id" json:"image_id,optional,no_refresh"`
	MemoryMiB        types.Int64                                  `tfsdk:"memory_mib" json:"memory_mib,computed_optional"`
	VCPU             types.Float64                                `tfsdk:"vcpu" json:"vcpu,computed_optional"`
	Autosleep        types.String                                 `tfsdk:"autosleep" json:"autosleep,computed_optional,no_refresh"`
	StorageGiB       types.Int64                                  `tfsdk:"storage_gib" json:"storage_gib,computed_optional"`
	AutosleepSeconds types.Int64                                  `tfsdk:"autosleep_seconds" json:"autosleep_seconds,computed"`
	DesiredState     types.String                                 `tfsdk:"desired_state" json:"desired_state,computed"`
	Phase            types.String                                 `tfsdk:"phase" json:"phase,computed,no_refresh"`
	Status           customfield.NestedObject[MachineStatusModel] `tfsdk:"status" json:"status,computed"`
	Timeouts         timeouts.Value                               `tfsdk:"timeouts"`
}

func (m MachineModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MachineModel) MarshalJSONForUpdate(state MachineModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type MachineStatusModel struct {
	Autoresize       customfield.NestedObject[MachineStatusAutoresizeModel] `tfsdk:"autoresize" json:"autoresize,computed"`
	LastProgressAt   timetypes.RFC3339                                      `tfsdk:"last_progress_at" json:"last_progress_at,computed" format:"date-time"`
	LastTransitionAt timetypes.RFC3339                                      `tfsdk:"last_transition_at" json:"last_transition_at,computed" format:"date-time"`
	Phase            types.String                                           `tfsdk:"phase" json:"phase,computed"`
	Reason           types.String                                           `tfsdk:"reason" json:"reason,computed"`
	Retryable        types.Bool                                             `tfsdk:"retryable" json:"retryable,computed"`
	Revision         types.String                                           `tfsdk:"revision" json:"revision,computed"`
	LastError        types.String                                           `tfsdk:"last_error" json:"last_error,computed"`
}

type MachineStatusAutoresizeModel struct {
	CPU    customfield.NestedObject[MachineStatusAutoresizeCPUModel]    `tfsdk:"cpu" json:"cpu,computed"`
	Memory customfield.NestedObject[MachineStatusAutoresizeMemoryModel] `tfsdk:"memory" json:"memory,computed"`
}

type MachineStatusAutoresizeCPUModel struct {
	BlockedAt     timetypes.RFC3339 `tfsdk:"blocked_at" json:"blocked_at,computed" format:"date-time"`
	BlockedReason types.String      `tfsdk:"blocked_reason" json:"blocked_reason,computed"`
}

type MachineStatusAutoresizeMemoryModel struct {
	BlockedAt     timetypes.RFC3339 `tfsdk:"blocked_at" json:"blocked_at,computed" format:"date-time"`
	BlockedReason types.String      `tfsdk:"blocked_reason" json:"blocked_reason,computed"`
}
