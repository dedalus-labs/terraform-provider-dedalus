package test_helpers_test

import (
	"context"
	"testing"

	dstimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	rstimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	ds "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rs "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/test_helpers"
)

type resourceTimeoutModel struct {
	Timeouts rstimeouts.Value `tfsdk:"timeouts"`
}

type datasourceTimeoutModel struct {
	Timeouts dstimeouts.Value `tfsdk:"timeouts"`
}

type timeoutFields struct {
	Read types.String `tfsdk:"read"`
}

type typedTimeoutModel struct {
	Timeouts customfield.NestedObject[timeoutFields] `tfsdk:"timeouts"`
}

func TestTimeoutObjectSchemaParity(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("resource", func(t *testing.T) {
		schema := rs.Schema{Attributes: map[string]rs.Attribute{"timeouts": rstimeouts.AttributesAll(ctx)}}
		errs := test_helpers.ValidateResourceModelSchemaIntegrity(resourceTimeoutModel{}, schema)
		errs.Report(t)
	})
	t.Run("datasource", func(t *testing.T) {
		schema := ds.Schema{Attributes: map[string]ds.Attribute{"timeouts": dstimeouts.Attributes(ctx)}}
		errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(datasourceTimeoutModel{}, schema)
		errs.Report(t)
	})
	t.Run("object_attribute", func(t *testing.T) {
		fields := map[string]attr.Type{"read": types.StringType}
		schema := rs.Schema{Attributes: map[string]rs.Attribute{"timeouts": rs.ObjectAttribute{
			Optional: true, AttributeTypes: fields,
			CustomType: rstimeouts.Type{ObjectType: types.ObjectType{AttrTypes: fields}},
		}}}
		errs := test_helpers.ValidateResourceModelSchemaIntegrity(resourceTimeoutModel{}, schema)
		errs.Report(t)
	})
	t.Run("wrong_custom_type", func(t *testing.T) {
		schema := rs.Schema{Attributes: map[string]rs.Attribute{"timeouts": rs.SingleNestedAttribute{
			Optional:   true,
			Attributes: map[string]rs.Attribute{"read": rs.StringAttribute{Optional: true}},
			CustomType: customfield.NewNestedObjectType[timeoutFields](ctx),
		}}}
		if errs := test_helpers.ValidateResourceModelSchemaIntegrity(resourceTimeoutModel{}, schema); len(errs) == 0 {
			t.Fatal("accepted a timeout value with a different custom type")
		}
	})
	t.Run("missing_typed_field", func(t *testing.T) {
		schema := rs.Schema{Attributes: map[string]rs.Attribute{"timeouts": rs.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]rs.Attribute{
				"read":  rs.StringAttribute{Optional: true},
				"write": rs.StringAttribute{Optional: true},
			},
			CustomType: customfield.NewNestedObjectType[timeoutFields](ctx),
		}}}
		if errs := test_helpers.ValidateResourceModelSchemaIntegrity(typedTimeoutModel{}, schema); len(errs) == 0 {
			t.Fatal("accepted a typed object missing a schema field")
		}
	})
}
