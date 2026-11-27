// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MachinePortResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"port_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"machine_id": schema.StringAttribute{
				Description:   "Bare, lowercase, hyphenated Machine UUID. Pass the returned machine_id unchanged.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"port": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"expires": schema.Int64Attribute{
				Description: "Port lifetime in integer seconds. Omit for no expiry.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"protocol": schema.StringAttribute{
				Description: `Available values: "http".`,
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("http"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("http"),
			},
			"password": schema.StringAttribute{
				Description: "Required when visibility is key_required. Choose exactly 8 ASCII letters or digits. Forbidden for private and public ports. No access key is generated.",
				Optional:    true,
				Sensitive:   true,
			},
			"visibility": schema.StringAttribute{
				Description: "Who may reach the port. private requires an owning-organization API key. key_required requires an owner-chosen password at creation. public accepts anonymous requests, which can wake the machine and incur compute charges.\nAvailable values: \"public\", \"key_required\", \"private\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"public",
						"key_required",
						"private",
					),
				},
				Default: stringdefault.StaticString("private"),
			},
			"application_url": schema.StringAttribute{
				Description: "Isolated guest application origin. Direct clients preserve their HTTP method and body and send any required Dedalus credential in headers.",
				Computed:    true,
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
				Description: "Unavailable is a recoverable failure reaching the guest service. See error_code and error_message for the diagnosis.\nAvailable values: \"wake_in_progress\", \"ready\", \"unavailable\", \"closed\", \"expired\", \"failed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"wake_in_progress",
						"ready",
						"unavailable",
						"closed",
						"expired",
						"failed",
					),
				},
			},
			"url": schema.StringAttribute{
				Description: "Stable browser share URL. Browser navigation only; use application_url for APIs, webhooks, and WebSockets.",
				Computed:    true,
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create:            true,
				CreateDescription: "The timeout for the operation, default: 300 seconds",
				Read:              true,
				ReadDescription:   "The timeout for the operation, default: 120 seconds",
				Update:            true,
				UpdateDescription: "The timeout for the operation, default: 120 seconds",
				Delete:            true,
				DeleteDescription: "The timeout for the operation, default: 300 seconds",
			}),
		},
	}
}

func (r *MachinePortResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MachinePortResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
