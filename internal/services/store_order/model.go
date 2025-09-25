// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order

import (
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StoreOrderModel struct {
	ID       types.Int64       `tfsdk:"id" json:"id,required"`
	OrderID  types.Int64       `tfsdk:"order_id" path:"orderId,optional"`
	Complete types.Bool        `tfsdk:"complete" json:"complete,optional"`
	PetID    types.Int64       `tfsdk:"pet_id" json:"petId,optional"`
	Quantity types.Int64       `tfsdk:"quantity" json:"quantity,optional"`
	ShipDate timetypes.RFC3339 `tfsdk:"ship_date" json:"shipDate,optional" format:"date-time"`
	Status   types.String      `tfsdk:"status" json:"status,optional"`
}

func (m StoreOrderModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StoreOrderModel) MarshalJSONForUpdate(state StoreOrderModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
