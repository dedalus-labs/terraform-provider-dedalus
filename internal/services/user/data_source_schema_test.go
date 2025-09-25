// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/dedalus-terraform/internal/services/user"
	"github.com/stainless-sdks/dedalus-terraform/internal/test_helpers"
)

func TestUserDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user.UserDataSourceModel)(nil)
	schema := user.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
