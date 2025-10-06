// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StoreOrderDataSourceModel struct {
	ID       types.Int64       `tfsdk:"id" path:"orderId,computed"`
	OrderID  types.Int64       `tfsdk:"order_id" path:"orderId,required"`
	Complete types.Bool        `tfsdk:"complete" json:"complete,computed"`
	PetID    types.Int64       `tfsdk:"pet_id" json:"petId,computed"`
	Quantity types.Int64       `tfsdk:"quantity" json:"quantity,computed"`
	ShipDate timetypes.RFC3339 `tfsdk:"ship_date" json:"shipDate,computed" format:"date-time"`
	Status   types.String      `tfsdk:"status" json:"status,computed"`
}
