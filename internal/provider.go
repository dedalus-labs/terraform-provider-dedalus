// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*DedalusProvider)(nil)

// DedalusProvider defines the provider implementation.
type DedalusProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// DedalusProviderModel describes the provider data model.
type DedalusProviderModel struct {
	BaseURL      types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey       types.String `tfsdk:"api_key" json:"api_key,optional"`
	XAPIKey      types.String `tfsdk:"x_api_key" json:"x_api_key,optional"`
	DedalusOrgID types.String `tfsdk:"dedalus_org_id" json:"dedalus_org_id,optional"`
}

func (p *DedalusProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dedalus"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"x_api_key": schema.StringAttribute{
				Optional: true,
			},
			"dedalus_org_id": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *DedalusProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *DedalusProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data DedalusProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("DEDALUS_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("DEDALUS_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	}

	if !data.XAPIKey.IsNull() && !data.XAPIKey.IsUnknown() {
		opts = append(opts, option.WithXAPIKey(data.XAPIKey.ValueString()))
	} else if o, ok := os.LookupEnv("DEDALUS_X_API_KEY"); ok {
		opts = append(opts, option.WithXAPIKey(o))
	}

	if !data.DedalusOrgID.IsNull() && !data.DedalusOrgID.IsUnknown() {
		opts = append(opts, option.WithDedalusOrgID(data.DedalusOrgID.ValueString()))
	} else if o, ok := os.LookupEnv("DEDALUS_ORG_ID"); ok {
		opts = append(opts, option.WithDedalusOrgID(o))
	}

	client := dedalus.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *DedalusProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *DedalusProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		machine.NewResource,
	}
}

func (p *DedalusProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		machine.NewMachineDataSource,
		machine.NewMachinesDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DedalusProvider{
			version: version,
		}
	}
}
