// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StoreOrderDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"order_id": schema.Int64Attribute{
				Required: true,
			},
			"complete": schema.BoolAttribute{
				Computed: true,
			},
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"pet_id": schema.Int64Attribute{
				Computed: true,
			},
			"quantity": schema.Int64Attribute{
				Computed: true,
			},
			"ship_date": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Order Status\nAvailable values: \"placed\", \"approved\", \"delivered\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"placed",
						"approved",
						"delivered",
					),
				},
			},
		},
	}
}

func (d *StoreOrderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StoreOrderDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
