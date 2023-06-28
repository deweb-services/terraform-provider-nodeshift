package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// dwsProviderModel maps provider schema data to a Go type.
type dwsProviderModel struct {
	AccessKey       types.String `tfsdk:"access_key"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`

	SharedCredentialsFile types.String `tfsdk:"shared_credentials_file"`
	Profile               types.String `tfsdk:"profile"`
}
