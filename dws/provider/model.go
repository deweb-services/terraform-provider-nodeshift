package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// dwsProviderModel maps provider schema data to a Go type.
type dwsProviderModel struct {
	AccountName  types.String `tfsdk:"account_name"`
	AccountKey   types.String `tfsdk:"account_key"`
	AccessRegion types.String `tfsdk:"access_region"`
	ApiKey       types.String `tfsdk:"api_key"`
	SessionToken types.String `tfsdk:"session_token"`
}
