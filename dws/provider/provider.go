package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/deweb-services/terraform-provider-dws/dws/resource/gpu"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/deweb-services/terraform-provider-dws/dws/resource/deployment"
	"github.com/deweb-services/terraform-provider-dws/dws/resource/gpu"
	s3terraform "github.com/deweb-services/terraform-provider-dws/dws/resource/s3"
	"github.com/deweb-services/terraform-provider-dws/dws/resource/vpc"
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
		Description: "Interact with DWS provider",
		Attributes: map[string]schema.Attribute{
			AccessKey: schema.StringAttribute{
				Description: "Access Key for DWS",
				Optional:    true,
				Sensitive:   true,
			},
			SecretAccessKey: schema.StringAttribute{
				Description: "Secret Access Key for DWS",
				Optional:    true,
				Sensitive:   true,
			},
			SharedCredentialsFile: schema.StringAttribute{
				Description: "Path to credentials file DWS",

				Optional: true,
			},
			Profile: schema.StringAttribute{
				Description: "DWS profile name",
				Optional:    true,
			},
			S3Endpoint: schema.StringAttribute{
				Description: "DWS s3 endpoint address",
				Optional:    true,
			},
			S3Region: schema.StringAttribute{
				Description: "DWS s3 region",
				Optional:    true,
			},
		},
	}
}

// Configure prepares a DWS API client for data sources and resources.
func (p *dwsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring DWS client")

	accessKey := os.Getenv(EnvKeyAccessKey)
	secretAccessKey := os.Getenv(EnvKeySecretAccessKey)
	sharedCredentialsFile := os.Getenv(EnvKeySharedCredentialsFile)
	profile := os.Getenv(EnvKeyProfile)
	s3Endpoint := os.Getenv(EnvKeyS3Endpoint)
	s3Region := os.Getenv(EnvKeyS3Region)

	values := []string{
		accessKey,
		secretAccessKey,
		sharedCredentialsFile,
		profile,
		s3Endpoint,
		s3Region,
	}

	// Retrieve provider data from configuration
	var config dwsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors configuring DWS client", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	if config.AccessKey.ValueString() != "" {
		values[0] = config.AccessKey.ValueString()
	}

	if config.SecretAccessKey.ValueString() != "" {
		values[1] = config.SecretAccessKey.ValueString()
	}

	if config.SharedCredentialsFile.ValueString() != "" {
		values[2] = config.SharedCredentialsFile.ValueString()
	}

	if config.Profile.ValueString() != "" {
		values[3] = config.Profile.ValueString()
	}

	if config.S3Endpoint.ValueString() != "" {
		values[4] = config.S3Endpoint.ValueString()
	}

	if config.S3Region.ValueString() != "" {
		values[5] = config.S3Region.ValueString()
	}

	type Attribute struct {
		EnvName  string
		Param    *types.String
		Required bool
	}

	attributes := map[string]Attribute{
		AccessKey: {
			EnvName:  EnvKeyAccessKey,
			Param:    &config.AccessKey,
			Required: false,
		},
		SecretAccessKey: {
			EnvName:  EnvKeySecretAccessKey,
			Param:    &config.SecretAccessKey,
			Required: false,
		},
		SharedCredentialsFile: {
			EnvName:  EnvKeySharedCredentialsFile,
			Param:    &config.SharedCredentialsFile,
			Required: false,
		},
		Profile: {
			EnvName:  EnvKeyProfile,
			Param:    &config.Profile,
			Required: false,
		},
		S3Endpoint: {
			EnvName:  EnvKeyS3Endpoint,
			Param:    &config.S3Endpoint,
			Required: false,
		},
		S3Region: {
			EnvName:  EnvKeyS3Region,
			Param:    &config.S3Region,
			Required: false,
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
		if val == "" && v.Required {
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

	tflog.Debug(ctx, "Creating DWS client")
	var cfg client.DWSProviderConfiguration
	cfg.FromSlice(values)
	tflog.Debug(ctx, fmt.Sprintf("%+v", values))
	// Create a new dws client using the configuration values
	cli := client.NewClient(ctx, cfg, client.ClientOptWithURL(client.APIURL), client.ClientOptWithS3())
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
		deployment.NewDeploymentResource,
		vpc.NewVPCResource,
		gpu.NewGPUResource,
		s3terraform.NewBucketResource,
	}
}
