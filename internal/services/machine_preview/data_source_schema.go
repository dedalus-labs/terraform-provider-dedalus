// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_preview

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MachinePreviewDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"preview_id": schema.StringAttribute{
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
			"port": schema.Int64Attribute{
				Computed: true,
			},
			"protocol": schema.StringAttribute{
				Description: `Available values: "http", "https".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("http", "https"),
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
			"url": schema.StringAttribute{
				Computed: true,
			},
			"visibility": schema.StringAttribute{
				Description: `Available values: "public", "private", "org".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"public",
						"private",
						"org",
					),
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				ReadDescription: "The timeout for the operation, default: 300 seconds",
			}),
		},
	}
}

func (d *MachinePreviewDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MachinePreviewDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
