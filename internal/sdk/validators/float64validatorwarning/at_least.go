package float64validatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Float64 = atLeastValidator{}

type atLeastValidator struct {
	min float64
}

func (v atLeastValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v atLeastValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value must be at least %g.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min,
	)
}

func (v atLeastValidator) ValidateFloat64(ctx context.Context, request validator.Float64Request, response *validator.Float64Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueFloat64()

	if value < v.min {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Value",
			fmt.Sprintf("Attribute %s has value %g, expected at least %g.", request.Path, value, v.min),
		))
	}
}

// AtLeast checks that the Float64 value is at least the given minimum.
// If not, it emits a warning diagnostic.
func AtLeast(min float64) validator.Float64 {
	return atLeastValidator{
		min: min,
	}
}
