// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*StoreOrderResource)(nil)

func (r *StoreOrderResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
