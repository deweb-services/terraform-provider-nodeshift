package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// nodeshiftProviderModel maps provider schema data to a Go type.
type nodeshiftProviderModel struct {
	AccessKey       types.String `tfsdk:"access_key"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`

	SharedCredentialsFile types.String `tfsdk:"shared_credentials_file"`
	Profile               types.String `tfsdk:"profile"`

	S3Endpoint types.String `tfsdk:"s3_endpoint"`
	S3Region   types.String `tfsdk:"s3_region"`
}
