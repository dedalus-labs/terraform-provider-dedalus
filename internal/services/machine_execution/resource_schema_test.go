// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_execution_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_execution"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineExecutionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_execution.MachineExecutionModel)(nil)
	schema := machine_execution.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
