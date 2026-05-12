// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_ssh_session

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

var _ datasource.DataSourceWithConfigValidators = (*MachineSSHSessionsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[MachineSSHSessionsItemsDataSourceModel](ctx),
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
						"session_id": schema.StringAttribute{
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
							CustomType: customfield.NewNestedObjectType[MachineSSHSessionsMachineSSHSessionConnectionDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[MachineSSHSessionsMachineSSHSessionConnectionHostTrustDataSourceModel](ctx),
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
					},
				},
			},
			"timeouts": timeouts.AttributesWithOpts(ctx, timeouts.Opts{
				Read:            true,
				ReadDescription: "The timeout for the operation, default: 120 seconds",
			}),
		},
	}
}

func (d *MachineSSHSessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MachineSSHSessionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
