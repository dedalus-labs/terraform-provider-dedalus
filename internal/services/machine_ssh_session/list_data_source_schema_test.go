// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_ssh_session_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_ssh_session"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachineSSHSessionsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_ssh_session.MachineSSHSessionsDataSourceModel)(nil)
	schema := machine_ssh_session.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
