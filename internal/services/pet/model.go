// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pet

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/dedalus-terraform/internal/apijson"
	"github.com/tidwall/sjson"
)

type PetModel struct {
	ID        types.Int64       `tfsdk:"id" json:"id,required"`
	PetID     types.Int64       `tfsdk:"pet_id" path:"petId,optional"`
	Name      types.String      `tfsdk:"name" json:"name,required"`
	PhotoURLs *[]types.String   `tfsdk:"photo_urls" json:"photoUrls,required"`
	Status    types.String      `tfsdk:"status" json:"status,optional"`
	Category  *PetCategoryModel `tfsdk:"category" json:"category,optional"`
	Tags      *[]*PetTagsModel  `tfsdk:"tags" json:"tags,optional"`
}

func (m PetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PetModel) MarshalJSONForUpdate(state PetModel) (data []byte, err error) {
	data, err = apijson.MarshalForUpdate(m, state)
	if err != nil {
		return nil, err
	}
	return sjson.SetBytes(data, "id", m.ID.ValueString())
}

type PetCategoryModel struct {
	ID   types.Int64  `tfsdk:"id" json:"id,optional"`
	Name types.String `tfsdk:"name" json:"name,optional"`
}

type PetTagsModel struct {
	ID   types.Int64  `tfsdk:"id" json:"id,optional"`
	Name types.String `tfsdk:"name" json:"name,optional"`
}
