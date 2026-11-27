// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MachinePortDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"port_id": schema.StringAttribute{
				Required: true,
			},
			"machine_id": schema.StringAttribute{
				Description: "Bare, lowercase, hyphenated Machine UUID. Pass the returned machine_id unchanged.",
				Required:    true,
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
			"port": schema.Int64Attribute{
				Computed: true,
			},
			"protocol": schema.StringAttribute{
				Description: `Available values: "http".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("http"),
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
			"visibility": schema.StringAttribute{
				Description: `Available values: "public", "key_required", "private".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"public",
						"key_required",
						"private",
					),
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 300 seconds",
			}),
		},
	}
}

func (d *MachinePortDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MachinePortDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
