// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine

import (
	"context"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MachineResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"machine_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"memory_mib": schema.Int64Attribute{
				Description: "Memory in MiB.",
				Required:    true,
			},
			"storage_gib": schema.Int64Attribute{
				Description: "Storage in GiB.",
				Required:    true,
			},
			"vcpu": schema.Float64Attribute{
				Description: "CPU in vCPUs.",
				Required:    true,
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
			"status": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MachineStatusModel](ctx),
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
