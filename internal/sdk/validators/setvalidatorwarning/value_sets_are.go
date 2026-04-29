package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = valueSetsAreValidator{}

type valueSetsAreValidator struct {
	validators []validator.Set
}

func (v valueSetsAreValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v valueSetsAreValidator) MarkdownDescription(_ context.Context) string {
	return "Each set element must satisfy the set validators.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct."
}

func (v valueSetsAreValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, el := range request.ConfigValue.Elements() {
		setVal, ok := el.(types.Set)
		if !ok {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path.AtSetValue(el),
				"Invalid Set Element",
				fmt.Sprintf("Set element at %s is not a set.", request.Path.AtSetValue(el)),
			))
			continue
		}

		elementPath := request.Path.AtSetValue(el)
		setReq := validator.SetRequest{
			Path:        elementPath,
			ConfigValue: setVal,
		}
		var setResp validator.SetResponse

		for _, val := range v.validators {
			val.ValidateSet(ctx, setReq, &setResp)
		}

		response.Diagnostics.Append(setResp.Diagnostics...)
	}
}

// ValueSetsAre checks that each element of the Set (of sets) satisfies all of the given set validators.
// If not, it emits warning diagnostics.
func ValueSetsAre(validators ...validator.Set) validator.Set {
	return valueSetsAreValidator{
		validators: validators,
	}
}
