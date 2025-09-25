// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pet

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PetDataSourceModel struct {
	PetID     types.Int64                                          `tfsdk:"pet_id" path:"petId,required"`
	ID        types.Int64                                          `tfsdk:"id" json:"id,computed"`
	Name      types.String                                         `tfsdk:"name" json:"name,computed"`
	Status    types.String                                         `tfsdk:"status" json:"status,computed"`
	PhotoURLs customfield.List[types.String]                       `tfsdk:"photo_urls" json:"photoUrls,computed"`
	Category  customfield.NestedObject[PetCategoryDataSourceModel] `tfsdk:"category" json:"category,computed"`
	Tags      customfield.NestedObjectList[PetTagsDataSourceModel] `tfsdk:"tags" json:"tags,computed"`
}

type PetCategoryDataSourceModel struct {
	ID   types.Int64  `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type PetTagsDataSourceModel struct {
	ID   types.Int64  `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
