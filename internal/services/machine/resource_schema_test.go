// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine.MachineModel)(nil)
	schema := machine.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
