// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine

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

var _ datasource.DataSourceWithConfigValidators = (*MachineDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Bare, lowercase, hyphenated Machine UUID. Pass the returned machine_id unchanged.",
				Computed:    true,
			},
			"machine_id": schema.StringAttribute{
				Description: "Bare, lowercase, hyphenated Machine UUID. Pass the returned machine_id unchanged.",
				Required:    true,
			},
			"autosleep_seconds": schema.Int64Attribute{
				Description: "Seconds of inactivity before autosleep. 0 disables autosleep.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9223372036),
				},
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
			"storage_gib": schema.Int64Attribute{
				Computed: true,
			},
			"vcpu": schema.Float64Attribute{
				Description: "CPU in vCPUs.",
				Computed:    true,
			},
			"status": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MachineStatusDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"autoresize": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"cpu": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeCPUDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"blocked_at": schema.StringAttribute{
										Description: "Time the current refusal was first recorded.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"blocked_reason": schema.StringAttribute{
										Description: "Why the latest automatic increase could not proceed, as the error code an explicit resize would return. Absent once an automatic increase is accepted.\nAvailable values: \"MACHINE_GROWTH_LIMIT_REACHED\", \"MACHINE_TIER_SHAPE_EXCEEDED\", \"MACHINE_QUOTA_EXCEEDED\", \"MACHINE_MEMORY_ABOVE_LIMIT\", \"ASSIGNED_HOST_CAPACITY_CONFLICT\", \"MACHINE_PLACEMENT_EXHAUSTED\", \"MACHINE_RESIZE_UNAVAILABLE\", \"MACHINE_SHAPE_UNSUPPORTED\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"MACHINE_GROWTH_LIMIT_REACHED",
												"MACHINE_TIER_SHAPE_EXCEEDED",
												"MACHINE_QUOTA_EXCEEDED",
												"MACHINE_MEMORY_ABOVE_LIMIT",
												"ASSIGNED_HOST_CAPACITY_CONFLICT",
												"MACHINE_PLACEMENT_EXHAUSTED",
												"MACHINE_RESIZE_UNAVAILABLE",
												"MACHINE_SHAPE_UNSUPPORTED",
											),
										},
									},
								},
							},
							"memory": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeMemoryDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"blocked_at": schema.StringAttribute{
										Description: "Time the current refusal was first recorded.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"blocked_reason": schema.StringAttribute{
										Description: "Why the latest automatic increase could not proceed, as the error code an explicit resize would return. Absent once an automatic increase is accepted.\nAvailable values: \"MACHINE_GROWTH_LIMIT_REACHED\", \"MACHINE_TIER_SHAPE_EXCEEDED\", \"MACHINE_QUOTA_EXCEEDED\", \"MACHINE_MEMORY_ABOVE_LIMIT\", \"ASSIGNED_HOST_CAPACITY_CONFLICT\", \"MACHINE_PLACEMENT_EXHAUSTED\", \"MACHINE_RESIZE_UNAVAILABLE\", \"MACHINE_SHAPE_UNSUPPORTED\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"MACHINE_GROWTH_LIMIT_REACHED",
												"MACHINE_TIER_SHAPE_EXCEEDED",
												"MACHINE_QUOTA_EXCEEDED",
												"MACHINE_MEMORY_ABOVE_LIMIT",
												"ASSIGNED_HOST_CAPACITY_CONFLICT",
												"MACHINE_PLACEMENT_EXHAUSTED",
												"MACHINE_RESIZE_UNAVAILABLE",
												"MACHINE_SHAPE_UNSUPPORTED",
											),
										},
									},
								},
							},
						},
					},
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
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 600 seconds",
			}),
		},
	}
}

func (d *MachineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MachineDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
