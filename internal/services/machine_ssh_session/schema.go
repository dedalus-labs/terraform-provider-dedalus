// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_ssh_session

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

var _ resource.ResourceWithConfigValidators = (*MachineSSHSessionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"session_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"machine_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"public_key": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"machine_ssh_session_connection": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MachineSSHSessionMachineSSHSessionConnectionModel](ctx),
				Attributes: map[string]schema.Attribute{
					"endpoint": schema.StringAttribute{
						Computed: true,
					},
					"port": schema.Int64Attribute{
						Computed: true,
					},
					"ssh_username": schema.StringAttribute{
						Computed: true,
					},
					"host_trust": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[MachineSSHSessionMachineSSHSessionConnectionHostTrustModel](ctx),
						Attributes: map[string]schema.Attribute{
							"host_pattern": schema.StringAttribute{
								Computed: true,
							},
							"kind": schema.StringAttribute{
								Description: `Available values: "cert_authority".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("cert_authority"),
								},
							},
							"public_key": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"user_certificate": schema.StringAttribute{
						Computed: true,
					},
				},
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

func (r *MachineSSHSessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MachineSSHSessionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
