// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachinesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine.MachinesDataSourceModel)(nil)
	schema := machine.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
