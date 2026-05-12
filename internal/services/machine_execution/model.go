// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_execution

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineExecutionModel struct {
	ID              types.String                                                 `tfsdk:"id" json:"-,computed"`
	ExecutionID     types.String                                                 `tfsdk:"execution_id" json:"execution_id,computed"`
	MachineID       types.String                                                 `tfsdk:"machine_id" path:"machine_id,required"`
	Command         *[]types.String                                              `tfsdk:"command" json:"command,required"`
	Cwd             types.String                                                 `tfsdk:"cwd" json:"cwd,optional"`
	Stdin           types.String                                                 `tfsdk:"stdin" json:"stdin,optional,no_refresh"`
	TimeoutMs       types.Int64                                                  `tfsdk:"timeout_ms" json:"timeout_ms,optional,no_refresh"`
	Env             *map[string]types.String                                     `tfsdk:"env" json:"env,optional,no_refresh"`
	CompletedAt     timetypes.RFC3339                                            `tfsdk:"completed_at" json:"completed_at,computed" format:"date-time"`
	CreatedAt       timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode       types.String                                                 `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage    types.String                                                 `tfsdk:"error_message" json:"error_message,computed"`
	ExitCode        types.Int64                                                  `tfsdk:"exit_code" json:"exit_code,computed"`
	ExpiresAt       timetypes.RFC3339                                            `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	RetryAfterMs    types.Int64                                                  `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Signal          types.Int64                                                  `tfsdk:"signal" json:"signal,computed"`
	StartedAt       timetypes.RFC3339                                            `tfsdk:"started_at" json:"started_at,computed" format:"date-time"`
	Status          types.String                                                 `tfsdk:"status" json:"status,computed"`
	StderrBytes     types.Int64                                                  `tfsdk:"stderr_bytes" json:"stderr_bytes,computed"`
	StderrTruncated types.Bool                                                   `tfsdk:"stderr_truncated" json:"stderr_truncated,computed"`
	StdoutBytes     types.Int64                                                  `tfsdk:"stdout_bytes" json:"stdout_bytes,computed"`
	StdoutTruncated types.Bool                                                   `tfsdk:"stdout_truncated" json:"stdout_truncated,computed"`
	EnvKeys         customfield.List[types.String]                               `tfsdk:"env_keys" json:"env_keys,computed"`
	Artifacts       customfield.NestedObjectList[MachineExecutionArtifactsModel] `tfsdk:"artifacts" json:"artifacts,computed"`
	Timeouts        timeouts.Value                                               `tfsdk:"timeouts"`
}

func (m MachineExecutionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MachineExecutionModel) MarshalJSONForUpdate(state MachineExecutionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MachineExecutionArtifactsModel struct {
	ArtifactID types.String `tfsdk:"artifact_id" json:"artifact_id,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
}
