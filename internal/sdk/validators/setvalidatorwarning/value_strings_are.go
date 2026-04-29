package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = valueStringsAreValidator{}

type valueStringsAreValidator struct {
	validators []validator.String
}

func (v valueStringsAreValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v valueStringsAreValidator) MarkdownDescription(_ context.Context) string {
	return "Each set element must satisfy the string validators.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct."
}

func (v valueStringsAreValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, el := range request.ConfigValue.Elements() {
		strVal, ok := el.(types.String)
		if !ok {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path.AtSetValue(el),
				"Invalid Set Element",
				fmt.Sprintf("Set element at %s is not a string.", request.Path.AtSetValue(el)),
			))
			continue
		}

		elementPath := request.Path.AtSetValue(el)
		strReq := validator.StringRequest{
			Path:        elementPath,
			ConfigValue: strVal,
		}
		var strResp validator.StringResponse

		for _, val := range v.validators {
			val.ValidateString(ctx, strReq, &strResp)
		}

		response.Diagnostics.Append(strResp.Diagnostics...)
	}
}

// ValueStringsAre checks that each element of the Set (of strings) satisfies all of the given string validators.
// If not, it emits warning diagnostics.
func ValueStringsAre(validators ...validator.String) validator.Set {
	return valueStringsAreValidator{
		validators: validators,
	}
}
