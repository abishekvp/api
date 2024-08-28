package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &GetAccessCredential{}

type GetAccessCredential struct{}

func get_access_credential() function.Function {
	return &GetAccessCredential{}
}

func (f *GetAccessCredential) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "get_access_credential"
}

func (f *GetAccessCredential) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Compute tax for coffee",
		Description: "Given a price and tax rate, return the total cost including tax.",

		// Parameters: []function.Parameter{
		// 	function.Float64Parameter{
		// 		Name:        "account",
		// 		Description: "Price of coffee item.",
		// 	},
		// 	function.Float64Parameter{
		// 		Name:        "password",
		// 		Description: "Tax rate. 0.085 == 8.5%",
		// 	},
		// },
		Return: function.StringReturn{},
	}
}

func (f *GetAccessCredential) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Set the result

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, "total"))
}
