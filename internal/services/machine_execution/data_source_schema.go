// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_execution

import (
	"context"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MachineExecutionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"execution_id": schema.StringAttribute{
				Required: true,
			},
			"machine_id": schema.StringAttribute{
				Required: true,
			},
			"completed_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"cwd": schema.StringAttribute{
				Computed: true,
			},
			"error_code": schema.StringAttribute{
				Computed: true,
			},
			"error_message": schema.StringAttribute{
				Computed: true,
			},
			"exit_code": schema.Int64Attribute{
				Computed: true,
			},
			"expires_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"retry_after_ms": schema.Int64Attribute{
				Computed: true,
			},
			"signal": schema.Int64Attribute{
				Computed: true,
			},
			"started_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "wake_in_progress", "queued", "running", "succeeded", "failed", "cancelled", "expired".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"wake_in_progress",
						"queued",
						"running",
						"succeeded",
						"failed",
						"cancelled",
						"expired",
					),
				},
			},
			"stderr_bytes": schema.Int64Attribute{
				Computed: true,
			},
			"stderr_truncated": schema.BoolAttribute{
				Computed: true,
			},
			"stdout_bytes": schema.Int64Attribute{
				Computed: true,
			},
			"stdout_truncated": schema.BoolAttribute{
				Computed: true,
			},
			"command": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"env_keys": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"artifacts": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[MachineExecutionArtifactsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"artifact_id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				Read:            true,
				ReadDescription: "The timeout for the operation, default: 600 seconds",
			}),
		},
	}
}

func (d *MachineExecutionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MachineExecutionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
