package stringvalidatorwarning

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = lengthBetweenValidator{}

type lengthBetweenValidator struct {
	min int
	max int
}

func (v lengthBetweenValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v lengthBetweenValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value length must be between %d and %d characters.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min, v.max,
	)
}

func (v lengthBetweenValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()
	length := utf8.RuneCountInString(value)

	if length < v.min || length > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Length",
			fmt.Sprintf("Attribute %s has length %d, expected between %d and %d characters.", request.Path, length, v.min, v.max),
		))
	}
}

// LengthBetween checks that the String length is within the given inclusive bounds.
// If not, it emits a warning diagnostic.
func LengthBetween(min, max int) validator.String {
	return lengthBetweenValidator{
		min: min,
		max: max,
	}
}
