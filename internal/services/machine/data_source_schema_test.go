// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("codegen bug: timeouts.Value is opaque to parity walker, pending upstream fix")
	t.Parallel()
	model := (*machine.MachineDataSourceModel)(nil)
	schema := machine.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
