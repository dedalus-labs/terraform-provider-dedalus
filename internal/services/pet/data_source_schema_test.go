// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pet_test

import (
	"context"
	"testing"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/services/pet"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

func TestPetDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pet.PetDataSourceModel)(nil)
	schema := pet.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
