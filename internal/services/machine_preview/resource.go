// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_preview

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/option"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/importpath"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*MachinePreviewResource)(nil)
var _ resource.ResourceWithModifyPlan = (*MachinePreviewResource)(nil)
var _ resource.ResourceWithImportState = (*MachinePreviewResource)(nil)

func NewResource() resource.Resource {
	return &MachinePreviewResource{}
}

// MachinePreviewResource defines the resource implementation.
type MachinePreviewResource struct {
	client *dedalus.Client
}

func (r *MachinePreviewResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_machine_preview"
}

func (r *MachinePreviewResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (r *MachinePreviewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MachinePreviewModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, diags := data.Timeouts.Create(ctx, time.Duration(5*time.Minute))
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if operationTimeout != time.Duration(0) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, operationTimeout)
		defer cancel()
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Machines.Previews.New(
		ctx,
		dedalus.MachinePreviewNewParams{
			MachineID: data.MachineID.ValueString(),
		},
		option.WithRequestBody("application/json", dataBytes),
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
	data.ID = data.PreviewID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MachinePreviewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *MachinePreviewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MachinePreviewModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, diags := data.Timeouts.Read(ctx, time.Duration(2*time.Minute))
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
	_, err := r.client.Machines.Previews.Get(
		ctx,
		dedalus.MachinePreviewGetParams{
			MachineID: data.MachineID.ValueString(),
			PreviewID: data.PreviewID.ValueString(),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.PreviewID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MachinePreviewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MachinePreviewModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	operationTimeout, diags := data.Timeouts.Delete(ctx, time.Duration(5*time.Minute))
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if operationTimeout != time.Duration(0) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, operationTimeout)
		defer cancel()
	}

	_, err := r.client.Machines.Previews.Delete(
		ctx,
		dedalus.MachinePreviewDeleteParams{
			MachineID: data.MachineID.ValueString(),
			PreviewID: data.PreviewID.ValueString(),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.PreviewID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MachinePreviewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(MachinePreviewModel)

	path_machine_id := ""
	path_preview_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<machine_id>/<preview_id>",
		&path_machine_id,
		&path_preview_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.MachineID = types.StringValue(path_machine_id)
	data.PreviewID = types.StringValue(path_preview_id)

	operationTimeout, diags := data.Timeouts.Read(ctx, time.Duration(2*time.Minute))
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
	_, err := r.client.Machines.Previews.Get(
		ctx,
		dedalus.MachinePreviewGetParams{
			MachineID: path_machine_id,
			PreviewID: path_preview_id,
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.PreviewID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MachinePreviewResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
