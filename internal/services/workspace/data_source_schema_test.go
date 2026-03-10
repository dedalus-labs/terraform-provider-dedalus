// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/workspace"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestWorkspaceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workspace.WorkspaceDataSourceModel)(nil)
	schema := workspace.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
