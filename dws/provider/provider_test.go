package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deweb-services/terraform-provider-dws/dws/resource/deployment"
	"github.com/deweb-services/terraform-provider-dws/dws/resource/gpu"
	s3terraform "github.com/deweb-services/terraform-provider-dws/dws/resource/s3"
	"github.com/deweb-services/terraform-provider-dws/dws/resource/vpc"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func TestNewDWSProvider(t *testing.T) {
	tests := []struct {
		name string
		want provider.Provider
	}{
		{
			name: "new dws provider",
			want: &dwsProvider{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDWSProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDWSProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dwsProvider_Configure(t *testing.T) {
	type args struct {
		ctx  context.Context
		req  provider.ConfigureRequest
		resp *provider.ConfigureResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "dws provider configure",
			args: args{
				ctx: context.TODO(),
				req: provider.ConfigureRequest{
					TerraformVersion: "",
					Config: tfsdk.Config{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							"access_key":              tftypes.NewValue(tftypes.String, "access_key"),
							"secret_access_key":       tftypes.NewValue(tftypes.String, "secret_access_key"),
							"shared_credentials_file": tftypes.NewValue(tftypes.String, "shared_credentials_file"),
							"profile":                 tftypes.NewValue(tftypes.String, "profile"),
							"s3_region":               tftypes.NewValue(tftypes.String, "s3_region"),
							"s3_endpoint":             tftypes.NewValue(tftypes.String, "s3_endpoint"),
						}),
						Schema: schema.Schema{
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
						},
					},
				},
				resp: &provider.ConfigureResponse{},
			},
		},
		{
			name: "dws provider configure error",
			args: args{
				ctx: context.TODO(),
				req: provider.ConfigureRequest{
					TerraformVersion: "14.1",
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
				resp: &provider.ConfigureResponse{},
			},
		},
		{
			name: "dws provider configure unknown param",
			args: args{
				ctx: context.TODO(),
				req: provider.ConfigureRequest{
					Config: tfsdk.Config{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							"access_key":              tftypes.NewValue(tftypes.String, "access_key"),
							"secret_access_key":       tftypes.NewValue(tftypes.String, "secret_access_key"),
							"shared_credentials_file": tftypes.NewValue(tftypes.String, "shared_credentials_file"),
							"profile":                 tftypes.NewValue(tftypes.String, "profile"),
							"s3_endpoint":             tftypes.NewValue(tftypes.String, "s3_endpoint"),
							"s3_region":               tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
						}),
						Schema: schema.Schema{
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
									Required: true,
								},
							},
						},
					},
				},
				resp: &provider.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dwsProvider{}

			p.Configure(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_dwsProvider_DataSources(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name string
		args args
		want []func() datasource.DataSource
	}{
		{
			name: "dws provider data sources",
			args: args{
				in0: context.TODO(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dwsProvider{}
			if got := p.DataSources(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataSources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dwsProvider_Metadata(t *testing.T) {
	type args struct {
		in0  context.Context
		in1  provider.MetadataRequest
		resp *provider.MetadataResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "dws provider metadata",
			args: args{
				in0:  context.TODO(),
				in1:  provider.MetadataRequest{},
				resp: &provider.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dwsProvider{}
			p.Metadata(tt.args.in0, tt.args.in1, tt.args.resp)
		})
	}
}

func Test_dwsProvider_Resources(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name string
		args args
		want []func() resource.Resource
	}{
		{
			name: "dws provider resources",
			args: args{
				in0: context.TODO(),
			},
			want: []func() resource.Resource{
				deployment.NewDeploymentResource,
				vpc.NewVPCResource,
				gpu.NewGPUResource,
				s3terraform.NewBucketResource,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dwsProvider{}
			assert.Equal(t, len(tt.want), len(p.Resources(tt.args.in0)))
		})
	}
}

func Test_dwsProvider_Schema(t *testing.T) {
	type args struct {
		in0  context.Context
		in1  provider.SchemaRequest
		resp *provider.SchemaResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "dws provider schema",
			args: args{
				in0:  context.TODO(),
				in1:  provider.SchemaRequest{},
				resp: &provider.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dwsProvider{}
			p.Schema(tt.args.in0, tt.args.in1, tt.args.resp)
		})
	}
}
