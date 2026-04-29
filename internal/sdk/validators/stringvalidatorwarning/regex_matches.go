package stringvalidatorwarning

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = regexMatchesValidator{}

type regexMatchesValidator struct {
	re      *regexp.Regexp
	pattern string
}

func (v regexMatchesValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v regexMatchesValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Value should match regular expression: %q.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.pattern,
	)
}

func (v regexMatchesValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if !v.re.MatchString(value) {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Attribute Value Pattern",
			fmt.Sprintf("Attribute %s has value %q which does not match required pattern %q.", request.Path, value, v.pattern),
		))
	}
}

// RegexMatches checks that the String value matches the given regular expression
// pattern (case-sensitive). If it does not, it emits a warning diagnostic.
// The pattern must be a valid Go regular expression.
func RegexMatches(pattern string) validator.String {
	re := regexp.MustCompile(pattern)

	return regexMatchesValidator{
		re:      re,
		pattern: pattern,
	}
}
