// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_port"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachinePortsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_port.MachinePortsDataSourceModel)(nil)
	schema := machine_port.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
