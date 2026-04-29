package stringvalidatorwarning

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = lengthAtMostValidator{}

type lengthAtMostValidator struct {
	max int
}

func (v lengthAtMostValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v lengthAtMostValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value length must be at most %d characters.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.max,
	)
}

func (v lengthAtMostValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()
	length := utf8.RuneCountInString(value)

	if length > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Length",
			fmt.Sprintf("Attribute %s has length %d, expected at most %d characters.", request.Path, length, v.max),
		))
	}
}

// LengthAtMost checks that the String length is at most the given maximum.
// If not, it emits a warning diagnostic.
func LengthAtMost(max int) validator.String {
	return lengthAtMostValidator{
		max: max,
	}
}
