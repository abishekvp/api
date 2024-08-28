package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider              = &securdenProvider{}
	_ provider.ProviderWithFunctions = &securdenProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &securdenProvider{
			version: version,
		}
	}
}

// securdenProviderModel maps provider schema data to a Go type.
type securdenProviderModel struct {
	AuthToken types.String `tfsdk:"authtoken"`
}

// securdenProvider is the provider implementation.
type securdenProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *securdenProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hashicups"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *securdenProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"authtoken": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *securdenProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Securden client")

	// Retrieve provider data from configuration
	var config securdenProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.AuthToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("authtoken"),
			"Unknown Securden API AuthToken",
			"The provider cannot create the Securden API client as there is an unknown configuration value for the Securden API authtoken. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SECURDEN_AUTHTOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	authtoken := os.Getenv("SECURDEN_AUTHTOKEN")

	if !config.AuthToken.IsNull() {
		authtoken = config.AuthToken.ValueString()
	}

	ctx = tflog.SetField(ctx, "securden_authtoken", authtoken)
}

// DataSources defines the data sources implemented in the provider.
func (p *securdenProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCoffeesDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *securdenProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrderResource,
	}
}

func (p *securdenProvider) Functions(_ context.Context) []func() function.Function {
	return nil
	// return []func() function.Function{
	// 	get_access_credential,
	// }
}
