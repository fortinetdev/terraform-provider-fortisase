package stringvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = oneOfValidator{}

type oneOfValidator struct {
	values []types.String
}

func (v oneOfValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v oneOfValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Value should be one of: %v.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.", v.values)
}

func (v oneOfValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue

	for _, otherValue := range v.values {
		if value.Equal(otherValue) {
			return
		}
	}

	response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
		request.Path,
		"Invalid Attribute Value Match",
		fmt.Sprintf("Attribute %s got value: %s. %s", request.Path, value.ValueString(), v.Description(ctx)),
	))
}

// OneOf checks that the String held in the attribute is one of the given values.
// If not, it emits a warning diagnostic.
func OneOf(values ...string) validator.String {
	frameworkValues := make([]types.String, 0, len(values))

	for _, value := range values {
		frameworkValues = append(frameworkValues, types.StringValue(value))
	}

	return oneOfValidator{
		values: frameworkValues,
	}
}
