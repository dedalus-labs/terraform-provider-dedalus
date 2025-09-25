// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pet_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/dedalus-terraform/internal/services/pet"
	"github.com/stainless-sdks/dedalus-terraform/internal/test_helpers"
)

func TestPetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pet.PetModel)(nil)
	schema := pet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
