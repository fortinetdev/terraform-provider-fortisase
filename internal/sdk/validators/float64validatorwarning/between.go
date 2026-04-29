package float64validatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Float64 = betweenValidator{}

type betweenValidator struct {
	min float64
	max float64
}

func (v betweenValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v betweenValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value must be between %g and %g.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min, v.max,
	)
}

func (v betweenValidator) ValidateFloat64(ctx context.Context, request validator.Float64Request, response *validator.Float64Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueFloat64()

	if value < v.min || value > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Value",
			fmt.Sprintf("Attribute %s has value %g, expected between %g and %g.", request.Path, value, v.min, v.max),
		))
	}
}

// Between checks that the Float64 value is within the given inclusive bounds.
// If not, it emits a warning diagnostic.
func Between(min, max float64) validator.Float64 {
	return betweenValidator{
		min: min,
		max: max,
	}
}
