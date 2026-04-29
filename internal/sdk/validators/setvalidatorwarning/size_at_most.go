package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Set = sizeAtMostValidator{}

type sizeAtMostValidator struct {
	max int64
}

func (v sizeAtMostValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v sizeAtMostValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Set size must be at most %d elements.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.max,
	)
}

func (v sizeAtMostValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	size := int64(len(request.ConfigValue.Elements()))

	if size > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Set Size",
			fmt.Sprintf("Attribute %s has %d elements, expected at most %d.", request.Path, size, v.max),
		))
	}
}

// SizeAtMost checks that the Set has at most max elements.
// If not, it emits a warning diagnostic.
func SizeAtMost(max int64) validator.Set {
	return sizeAtMostValidator{
		max: max,
	}
}
