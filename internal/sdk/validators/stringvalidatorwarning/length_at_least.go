package stringvalidatorwarning

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = lengthAtLeastValidator{}

type lengthAtLeastValidator struct {
	min int
}

func (v lengthAtLeastValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v lengthAtLeastValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value length must be at least %d characters.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min,
	)
}

func (v lengthAtLeastValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()
	length := utf8.RuneCountInString(value)

	if length < v.min {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Length",
			fmt.Sprintf("Attribute %s has length %d, expected at least %d characters.", request.Path, length, v.min),
		))
	}
}

// LengthAtLeast checks that the String length is at least the given minimum.
// If not, it emits a warning diagnostic.
func LengthAtLeast(min int) validator.String {
	return lengthAtLeastValidator{
		min: min,
	}
}
