// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserDataSourceModel struct {
	Username   types.String `tfsdk:"username" path:"username,required"`
	Email      types.String `tfsdk:"email" json:"email,computed"`
	FirstName  types.String `tfsdk:"first_name" json:"firstName,computed"`
	ID         types.Int64  `tfsdk:"id" json:"id,computed"`
	LastName   types.String `tfsdk:"last_name" json:"lastName,computed"`
	Password   types.String `tfsdk:"password" json:"password,computed"`
	Phone      types.String `tfsdk:"phone" json:"phone,computed"`
	UserStatus types.Int64  `tfsdk:"user_status" json:"userStatus,computed"`
}
