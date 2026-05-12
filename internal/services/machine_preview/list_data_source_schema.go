// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_preview

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

var _ datasource.DataSourceWithConfigValidators = (*MachinePreviewsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"machine_id": schema.StringAttribute{
				Required: true,
			},
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
				CustomType:  customfield.NewNestedObjectListType[MachinePreviewsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"machine_id": schema.StringAttribute{
							Computed: true,
						},
						"port": schema.Int64Attribute{
							Computed: true,
						},
						"preview_id": schema.StringAttribute{
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
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				Read:            true,
				ReadDescription: "The timeout for the operation, default: 300 seconds",
			}),
		},
	}
}

func (d *MachinePreviewsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MachinePreviewsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
