package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Set = sizeAtLeastValidator{}

type sizeAtLeastValidator struct {
	min int64
}

func (v sizeAtLeastValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v sizeAtLeastValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"Set size must be at least %d elements.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min,
	)
}

func (v sizeAtLeastValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	size := int64(len(request.ConfigValue.Elements()))

	if size < v.min {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid Set Size",
			fmt.Sprintf("Attribute %s has %d elements, expected at least %d.", request.Path, size, v.min),
		))
	}
}

// SizeAtLeast checks that the Set has at least min elements.
// If not, it emits a warning diagnostic.
func SizeAtLeast(min int64) validator.Set {
	return sizeAtLeastValidator{
		min: min,
	}
}
