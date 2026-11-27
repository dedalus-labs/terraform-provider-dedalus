// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine

import (
	"context"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MachineResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"machine_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"image_id": schema.StringAttribute{
				Description:   "Organization-owned ready custom image ID. Requires Pro. Omit for the platform default.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"memory_mib": schema.Int64Attribute{
				Description:   "Memory in MiB.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(4096),
			},
			"vcpu": schema.Float64Attribute{
				Description:   "CPU in vCPUs.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(1),
			},
			"autosleep": schema.StringAttribute{
				Description: `Idle window before autosleep. Accepts fixed duration units like 30s, 30m, 2h, 7d3h4s, or 1w3d, raw seconds ("1800"), or never to disable.`,
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("300s"),
			},
			"storage_gib": schema.Int64Attribute{
				Description: "Storage in GiB.",
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(10),
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
			"status": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MachineStatusModel](ctx),
				Attributes: map[string]schema.Attribute{
					"autoresize": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"cpu": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeCPUModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[MachineStatusAutoresizeMemoryModel](ctx),
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
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create:            true,
				CreateDescription: "The timeout for the operation, default: 600 seconds",
				Read:              true,
				ReadDescription:   "The timeout for the operation, default: 120 seconds",
				Update:            true,
				UpdateDescription: "The timeout for the operation, default: 600 seconds",
				Delete:            true,
				DeleteDescription: "The timeout for the operation, default: 600 seconds",
			}),
		},
	}
}

func (r *MachineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MachineResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
