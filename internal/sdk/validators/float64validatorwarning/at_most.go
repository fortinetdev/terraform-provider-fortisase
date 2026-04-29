package float64validatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Float64 = atMostValidator{}

type atMostValidator struct {
	max float64
}

func (v atMostValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v atMostValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value must be at most %g.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.max,
	)
}

func (v atMostValidator) ValidateFloat64(ctx context.Context, request validator.Float64Request, response *validator.Float64Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueFloat64()

	if value > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Value",
			fmt.Sprintf("Attribute %s has value %g, expected at most %g.", request.Path, value, v.max),
		))
	}
}

// AtMost checks that the Float64 value is at most the given maximum.
// If not, it emits a warning diagnostic.
func AtMost(max float64) validator.Float64 {
	return atMostValidator{
		max: max,
	}
}
