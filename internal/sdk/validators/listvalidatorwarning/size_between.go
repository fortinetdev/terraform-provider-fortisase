package listvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = sizeBetweenValidator{}

type sizeBetweenValidator struct {
	min int64
	max int64
}

func (v sizeBetweenValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v sizeBetweenValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf(
		"List size must be between %d and %d elements.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct.",
		v.min, v.max,
	)
}

func (v sizeBetweenValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	size := int64(len(request.ConfigValue.Elements()))

	if size < v.min || size > v.max {
		response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
			request.Path,
			"Invalid List Size",
			fmt.Sprintf("Attribute %s has %d elements, expected between %d and %d.", request.Path, size, v.min, v.max),
		))
	}
}

// SizeBetween checks that the List has between min and max elements (inclusive).
// If not, it emits a warning diagnostic.
func SizeBetween(min, max int64) validator.List {
	return sizeBetweenValidator{
		min: min,
		max: max,
	}
}
