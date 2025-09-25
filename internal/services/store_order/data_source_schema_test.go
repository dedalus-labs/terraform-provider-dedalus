// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_order_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/store_order"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestStoreOrderDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*store_order.StoreOrderDataSourceModel)(nil)
	schema := store_order.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
