// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id,required"`
	ExistingUsername types.String `tfsdk:"existing_username" path:"username,optional"`
	Email            types.String `tfsdk:"email" json:"email,optional"`
	FirstName        types.String `tfsdk:"first_name" json:"firstName,optional"`
	LastName         types.String `tfsdk:"last_name" json:"lastName,optional"`
	Password         types.String `tfsdk:"password" json:"password,optional"`
	Phone            types.String `tfsdk:"phone" json:"phone,optional"`
	UserStatus       types.Int64  `tfsdk:"user_status" json:"userStatus,optional"`
	Username         types.String `tfsdk:"username" json:"username,optional"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
