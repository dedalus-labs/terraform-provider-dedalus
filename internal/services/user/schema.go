// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ resource.ResourceWithConfigValidators = (*UserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Operations about user",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Optional: true,
			},
			"first_name": schema.StringAttribute{
				Optional: true,
			},
			"last_name": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional: true,
			},
			"phone": schema.StringAttribute{
				Optional: true,
			},
			"user_status": schema.Int64Attribute{
				Description: "User Status",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
