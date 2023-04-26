package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	res "github.com/deweb-services/terraform-provider-dws/dws/resource/vm"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &dwsProvider{}
)

// NewDWSProvider is a helper function to simplify provider server and testing implementation.
func NewDWSProvider() provider.Provider {
	return &dwsProvider{}
}

// dwsProvider is the provider implementation.
type dwsProvider struct{}

// Metadata returns the provider type name.
func (p *dwsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dws"
}

// Schema defines the provider-level schema for configuration data.
func (p *dwsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			AccountName: schema.StringAttribute{
				Required: true,
			},
			AccountKey: schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			AccessRegion: schema.StringAttribute{
				Optional: true,
			},
			ApiKey: schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			SessionToken: schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a DWS API client for data sources and resources.
func (p *dwsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring DWS client")

	// Retrieve provider data from configuration
	var config dwsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors configuring DWS client", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	type Attribute struct {
		EnvName string
		Param   *types.String
	}

	values := make([]string, 0)
	attributes := map[string]Attribute{
		AccountName: {
			EnvName: "DWS_ACCOUNT_NAME",
			Param:   &config.AccountName,
		},
		AccountKey: {
			EnvName: "DWS_ACCOUNT_KEY",
			Param:   &config.AccountKey,
		},
		AccessRegion: {
			EnvName: "DWS_ACCESS_REGION",
			Param:   &config.AccessRegion,
		},
		ApiKey: {
			EnvName: "DWS_API_KEY",
			Param:   &config.ApiKey,
		},
		SessionToken: {
			EnvName: "DWS_SESSION_TOKEN",
			Param:   &config.SessionToken,
		},
	}

	// If practitioner provided a configuration value for any of the attributes, it must be a known value.
	for attrName, v := range attributes {
		if v.Param.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root(attrName),
				fmt.Sprintf("Unknown DWS API %s", attrName),
				fmt.Sprintf("The provider cannot create the DWS API client as there is an unknown configuration "+
					"value for the dws API %s. Either target apply the source of the value first, set the value "+
					"statically in the configuration, or use the %s environment variable.", attrName, v.EnvName),
			)
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override with Terraform configuration value if set.
	for attrKey, v := range attributes {
		val := os.Getenv(v.EnvName)
		if !v.Param.IsNull() {
			val = v.Param.ValueString()
		}
		if val == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root(attrKey),
				fmt.Sprintf("Missing DWS API %s", attrKey),
				fmt.Sprintf("The provider cannot create the DWS API client as there is "+
					"a missing or empty value for the DWS API %s. Set the host value in the configuration "+
					"or use the %s environment variable. If either is already set, ensure the value is not empty.",
					attrKey, v.EnvName),
			)
		}
		values = append(values, val)
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "dws_region", config.AccessRegion)
	ctx = tflog.SetField(ctx, "dws_account_name", config.AccountName)

	tflog.Debug(ctx, "Creating DWS client")
	var cfg client.DWSProviderConfiguration
	cfg.FromSlice(values)
	// Create a new dws client using the configuration values
	cli := client.NewClient(cfg)
	// Make the dws client available during DataSource and Resource
	resp.DataSourceData = cli
	resp.ResourceData = cli
	tflog.Info(ctx, "Configured DWS client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *dwsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *dwsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		res.NewVMResource,
	}
}
