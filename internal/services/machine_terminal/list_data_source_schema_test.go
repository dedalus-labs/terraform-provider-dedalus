// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_terminal"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineTerminalsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_terminal.MachineTerminalsDataSourceModel)(nil)
	schema := machine_terminal.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
