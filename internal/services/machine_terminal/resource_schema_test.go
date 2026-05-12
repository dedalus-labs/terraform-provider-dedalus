// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_terminal_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_terminal"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineTerminalModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_terminal.MachineTerminalModel)(nil)
	schema := machine_terminal.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
