package setvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = valueListsAreValidator{}

type valueListsAreValidator struct {
	validators []validator.List
}

func (v valueListsAreValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v valueListsAreValidator) MarkdownDescription(_ context.Context) string {
	return "Each set element must satisfy the list validators.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct."
}

func (v valueListsAreValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, el := range request.ConfigValue.Elements() {
		listVal, ok := el.(types.List)
		if !ok {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path.AtSetValue(el),
				"Invalid Set Element",
				fmt.Sprintf("Set element at %s is not a list.", request.Path.AtSetValue(el)),
			))
			continue
		}

		elementPath := request.Path.AtSetValue(el)
		listReq := validator.ListRequest{
			Path:        elementPath,
			ConfigValue: listVal,
		}
		var listResp validator.ListResponse

		for _, val := range v.validators {
			val.ValidateList(ctx, listReq, &listResp)
		}

		response.Diagnostics.Append(listResp.Diagnostics...)
	}
}

// ValueListsAre checks that each element of the Set (of lists) satisfies all of the given list validators.
// If not, it emits warning diagnostics.
func ValueListsAre(validators ...validator.List) validator.Set {
	return valueListsAreValidator{
		validators: validators,
	}
}
