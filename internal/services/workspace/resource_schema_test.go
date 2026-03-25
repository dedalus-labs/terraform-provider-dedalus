// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/workspace"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestWorkspaceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workspace.WorkspaceModel)(nil)
	schema := workspace.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
