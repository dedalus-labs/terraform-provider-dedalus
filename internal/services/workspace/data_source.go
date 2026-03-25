// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WorkspaceDataSource struct {
	client *dedalus.Client
}

var _ datasource.DataSourceWithConfigure = (*WorkspaceDataSource)(nil)

func NewWorkspaceDataSource() datasource.DataSource {
	return &WorkspaceDataSource{}
}

func (d *WorkspaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

func (d *WorkspaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*dedalus.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *dedalus.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *WorkspaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WorkspaceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, diags := data.Timeouts.Read(ctx, time.Duration(10*time.Minute))
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if operationTimeout != time.Duration(0) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, operationTimeout)
		defer cancel()
	}

	res := new(http.Response)
	_, err := d.client.Workspaces.Get(
		ctx,
		data.WorkspaceID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.WorkspaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
