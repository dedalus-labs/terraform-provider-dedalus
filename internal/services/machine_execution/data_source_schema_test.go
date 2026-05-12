// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_execution_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_execution"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineExecutionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_execution.MachineExecutionDataSourceModel)(nil)
	schema := machine_execution.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
