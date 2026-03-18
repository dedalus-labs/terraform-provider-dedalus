// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"context"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkspaceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
			"desired_state": schema.StringAttribute{
				Description: `Available values: "active", "inactive", "destroyed".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"inactive",
						"destroyed",
					),
				},
			},
			"image_version": schema.StringAttribute{
				Computed: true,
			},
			"memory_mib": schema.Int64Attribute{
				Description: "Memory in MiB.",
				Computed:    true,
			},
			"schema": schema.StringAttribute{
				Description: "A URL to the JSON Schema for this object.",
				Computed:    true,
			},
			"storage_gib": schema.Int64Attribute{
				Computed: true,
			},
			"vcpu": schema.Float64Attribute{
				Description: "CPU in vCPUs.",
				Computed:    true,
			},
			"status": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[WorkspaceStatusDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"last_progress_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"last_transition_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"observed_revision": schema.StringAttribute{
						Computed: true,
					},
					"phase": schema.StringAttribute{
						Description: `Available values: "accepted", "placement_pending", "starting", "running", "stopping", "sleeping", "destroying", "destroyed", "failed".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"accepted",
								"placement_pending",
								"starting",
								"running",
								"stopping",
								"sleeping",
								"destroying",
								"destroyed",
								"failed",
							),
						},
					},
					"reason": schema.StringAttribute{
						Computed: true,
					},
					"retryable": schema.BoolAttribute{
						Computed: true,
					},
					"revision": schema.StringAttribute{
						Computed: true,
					},
					"assigned_host": schema.StringAttribute{
						Computed: true,
					},
					"last_error": schema.StringAttribute{
						Computed: true,
					},
					"memory_assigned_mib": schema.Int64Attribute{
						Computed: true,
					},
					"memory_resize_state": schema.StringAttribute{
						Computed: true,
					},
					"memory_target_mib": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 600 seconds",
			}),
		},
	}
}

func (d *WorkspaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkspaceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
