// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/dedalus-terraform/internal/services/store_order"
	"github.com/stainless-sdks/dedalus-terraform/internal/test_helpers"
)

func TestStoreOrderModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*store_order.StoreOrderModel)(nil)
	schema := store_order.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
