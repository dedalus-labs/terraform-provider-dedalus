// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_execution

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineExecutionsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MachineExecutionsItemsDataSourceModel] `json:"items,computed"`
}

type MachineExecutionsDataSourceModel struct {
	MachineID types.String                                                        `tfsdk:"machine_id" path:"machine_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[MachineExecutionsItemsDataSourceModel] `tfsdk:"items"`
	Timeouts  timeouts.Value                                                      `tfsdk:"timeouts"`
}

func (m *MachineExecutionsDataSourceModel) toListParams(_ context.Context) (params dedalus.MachineExecutionListParams, diags diag.Diagnostics) {
	params = dedalus.MachineExecutionListParams{
		MachineID: m.MachineID.ValueString(),
	}

	return
}

type MachineExecutionsItemsDataSourceModel struct {
	ID              types.String                                                            `tfsdk:"id" json:"execution_id,computed"`
	Command         customfield.List[types.String]                                          `tfsdk:"command" json:"command,computed"`
	CreatedAt       timetypes.RFC3339                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExecutionID     types.String                                                            `tfsdk:"execution_id" json:"execution_id,computed"`
	MachineID       types.String                                                            `tfsdk:"machine_id" json:"machine_id,computed"`
	Status          types.String                                                            `tfsdk:"status" json:"status,computed"`
	Artifacts       customfield.NestedObjectList[MachineExecutionsArtifactsDataSourceModel] `tfsdk:"artifacts" json:"artifacts,computed"`
	CompletedAt     timetypes.RFC3339                                                       `tfsdk:"completed_at" json:"completed_at,computed" format:"date-time"`
	Cwd             types.String                                                            `tfsdk:"cwd" json:"cwd,computed"`
	EnvKeys         customfield.List[types.String]                                          `tfsdk:"env_keys" json:"env_keys,computed"`
	ErrorCode       types.String                                                            `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage    types.String                                                            `tfsdk:"error_message" json:"error_message,computed"`
	ExitCode        types.Int64                                                             `tfsdk:"exit_code" json:"exit_code,computed"`
	ExpiresAt       timetypes.RFC3339                                                       `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	RetryAfterMs    types.Int64                                                             `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Signal          types.Int64                                                             `tfsdk:"signal" json:"signal,computed"`
	StartedAt       timetypes.RFC3339                                                       `tfsdk:"started_at" json:"started_at,computed" format:"date-time"`
	StderrBytes     types.Int64                                                             `tfsdk:"stderr_bytes" json:"stderr_bytes,computed"`
	StderrTruncated types.Bool                                                              `tfsdk:"stderr_truncated" json:"stderr_truncated,computed"`
	StdoutBytes     types.Int64                                                             `tfsdk:"stdout_bytes" json:"stdout_bytes,computed"`
	StdoutTruncated types.Bool                                                              `tfsdk:"stdout_truncated" json:"stdout_truncated,computed"`
}

type MachineExecutionsArtifactsDataSourceModel struct {
	ArtifactID types.String `tfsdk:"artifact_id" json:"artifact_id,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
}
