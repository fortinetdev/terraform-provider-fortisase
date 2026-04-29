package listvalidatorwarning

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.List = valueMapsAreValidator{}

type valueMapsAreValidator struct {
	validators []validator.Map
}

func (v valueMapsAreValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v valueMapsAreValidator) MarkdownDescription(_ context.Context) string {
	return "Each list element must satisfy the map validators.\nThis check is based on current supported OS version. But it may have changes if you are using a new version of OS that the provider not supported yet. Please make sure it is correct."
}

func (v valueMapsAreValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, el := range request.ConfigValue.Elements() {
		mapVal, ok := el.(types.Map)
		if !ok {
			response.Diagnostics.Append(diag.NewAttributeWarningDiagnostic(
				request.Path.AtListValue(el),
				"Invalid List Element",
				fmt.Sprintf("List element at %s is not a map.", request.Path.AtListValue(el)),
			))
			continue
		}

		elementPath := request.Path.AtListValue(el)
		mapReq := validator.MapRequest{
			Path:        elementPath,
			ConfigValue: mapVal,
		}
		var mapResp validator.MapResponse

		for _, val := range v.validators {
			val.ValidateMap(ctx, mapReq, &mapResp)
		}

		response.Diagnostics.Append(mapResp.Diagnostics...)
	}
}

// ValueMapsAre checks that each element of the List (of maps) satisfies all of the given map validators.
// If not, it emits warning diagnostics.
func ValueMapsAre(validators ...validator.Map) validator.List {
	return valueMapsAreValidator{
		validators: validators,
	}
}
