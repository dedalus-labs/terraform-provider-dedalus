// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_preview_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/machine_preview"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestMachinePreviewDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*machine_preview.MachinePreviewDataSourceModel)(nil)
	schema := machine_preview.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
