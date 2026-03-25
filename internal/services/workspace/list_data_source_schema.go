// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"context"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkspacesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkspacesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"desired_state": schema.StringAttribute{
							Description: `Available values: "running", "sleeping", "destroyed".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"running",
									"sleeping",
									"destroyed",
								),
							},
						},
						"memory_mib": schema.Int64Attribute{
							Description: "Memory in MiB.",
							Computed:    true,
						},
						"status": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[WorkspacesStatusDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"last_progress_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"last_transition_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
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
								"last_error": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"storage_gib": schema.Int64Attribute{
							Computed: true,
						},
						"vcpu": schema.Float64Attribute{
							Description: "CPU in vCPUs.",
							Computed:    true,
						},
						"workspace_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 600 seconds",
			}),
		},
	}
}

func (d *WorkspacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkspacesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
