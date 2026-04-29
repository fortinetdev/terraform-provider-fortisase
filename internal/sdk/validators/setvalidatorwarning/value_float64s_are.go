package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = valueFloat64sAreValidator{}

type valueFloat64sAreValidator struct {
	validators []validator.Float64
}

func (v valueFloat64sAreValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v valueFloat64sAreValidator) MarkdownDescription(_ context.Context) string {
	return "Each set element must satisfy the float64 validators.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct."
}

func (v valueFloat64sAreValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, el := range request.ConfigValue.Elements() {
		floatVal, ok := el.(types.Float64)
		if !ok {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path.AtSetValue(el),
				"Invalid Set Element",
				fmt.Sprintf("Set element at %s is not a float64.", request.Path.AtSetValue(el)),
			))
			continue
		}

		elementPath := request.Path.AtSetValue(el)
		floatReq := validator.Float64Request{
			Path:        elementPath,
			ConfigValue: floatVal,
		}
		var floatResp validator.Float64Response

		for _, val := range v.validators {
			val.ValidateFloat64(ctx, floatReq, &floatResp)
		}

		response.Diagnostics.Append(floatResp.Diagnostics...)
	}
}

// ValueFloat64sAre checks that each element of the Set (of float64s) satisfies all of the given float64 validators.
// If not, it emits warning diagnostics.
func ValueFloat64sAre(validators ...validator.Float64) validator.Set {
	return valueFloat64sAreValidator{
		validators: validators,
	}
}
