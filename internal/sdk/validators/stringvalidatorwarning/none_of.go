package stringvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = noneOfValidator{}

type noneOfValidator struct {
	values []types.String
}

func (v noneOfValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v noneOfValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Value should be none of: %v.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.", v.values)
}

func (v noneOfValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue

	for _, otherValue := range v.values {
		if value.Equal(otherValue) {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path,
				"Invalid Attribute Value Match",
				fmt.Sprintf("Attribute %s got disallowed value: %s. %s", request.Path, value.ValueString(), v.Description(ctx)),
			))
			return
		}
	}
}

// NoneOf checks that the String held in the attribute is none of the given values.
// If it matches any, it emits a warning diagnostic.
func NoneOf(values ...string) validator.String {
	frameworkValues := make([]types.String, 0, len(values))

	for _, value := range values {
		frameworkValues = append(frameworkValues, types.StringValue(value))
	}

	return noneOfValidator{
		values: frameworkValues,
	}
}
