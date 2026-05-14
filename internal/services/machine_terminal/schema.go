// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MachineTerminalResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"terminal_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"machine_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"height": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"width": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"cwd": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"shell": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"env": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"error_code": schema.StringAttribute{
				Computed: true,
			},
			"error_message": schema.StringAttribute{
				Computed: true,
			},
			"expires_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"protocol": schema.StringAttribute{
				Description: `Available values: "websocket".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("websocket"),
				},
			},
			"ready_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"retry_after_ms": schema.Int64Attribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "wake_in_progress", "ready", "closed", "expired", "failed".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"wake_in_progress",
						"ready",
						"closed",
						"expired",
						"failed",
					),
				},
			},
			"stream_url": schema.StringAttribute{
				Computed: true,
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create:            true,
				CreateDescription: "The timeout for the operation, default: 120 seconds",
				Read:              true,
				ReadDescription:   "The timeout for the operation, default: 60 seconds",
				Update:            true,
				UpdateDescription: "The timeout for the operation, default: No timeout",
				Delete:            true,
				DeleteDescription: "The timeout for the operation, default: 120 seconds",
			}),
		},
	}
}

func (r *MachineTerminalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MachineTerminalResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
