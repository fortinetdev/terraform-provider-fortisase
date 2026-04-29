package stringvalidatorwarning

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// RegexMatchesCI is a convenience wrapper around RegexMatches that enforces
// case-insensitive matching by prefixing the pattern with the (?i) flag.
func RegexMatchesCI(pattern string) validator.String {
	return RegexMatches("(?i)" + pattern)
}
