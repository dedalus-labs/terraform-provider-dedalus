// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MachineTerminalDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"terminal_id": schema.StringAttribute{
				Required: true,
			},
			"machine_id": schema.StringAttribute{
				Required: true,
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
			"height": schema.Int64Attribute{
				Computed: true,
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
			"width": schema.Int64Attribute{
				Computed: true,
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 120 seconds",
			}),
		},
	}
}

func (d *MachineTerminalDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MachineTerminalDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
