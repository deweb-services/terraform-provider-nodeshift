package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/resource/deployment"
	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/resource/gpu"
	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/resource/loadbalancer"
	s3terraform "github.com/deweb-services/terraform-provider-nodeshift/nodeshift/resource/s3"
	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/resource/vpc"
)

func TestNewNodeshiftProvider(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want provider.Provider
	}{
		{
			name: "new nodeshift provider",
			want: &nodeshiftProvider{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, NewNodeshiftProvider())
		})
	}
}

func Test_nodeshiftProvider_Configure(t *testing.T) {
	t.Parallel()

	type args struct {
		req  provider.ConfigureRequest
		resp *provider.ConfigureResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nodeshift provider configure",
			args: args{
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
							Description: "Interact with Nodeshift provider",
							Attributes: map[string]schema.Attribute{
								AccessKey: schema.StringAttribute{
									Description: "Access Key for Nodeshift",
									Optional:    true,
									Sensitive:   true,
								},
								SecretAccessKey: schema.StringAttribute{
									Description: "Secret Access Key for Nodeshift",
									Optional:    true,
									Sensitive:   true,
								},
								SharedCredentialsFile: schema.StringAttribute{
									Description: "Path to credentials file Nodeshift",

									Optional: true,
								},
								Profile: schema.StringAttribute{
									Description: "Nodeshift profile name",
									Optional:    true,
								},
								S3Endpoint: schema.StringAttribute{
									Description: "Nodeshift s3 endpoint address",
									Optional:    true,
								},
								S3Region: schema.StringAttribute{
									Description: "Nodeshift s3 region",
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
			name: "nodeshift provider configure error",
			args: args{
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
			name: "nodeshift provider configure unknown param",
			args: args{
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
							Description: "Interact with Nodeshift provider",
							Attributes: map[string]schema.Attribute{
								AccessKey: schema.StringAttribute{
									Description: "Access Key for Nodeshift",
									Optional:    true,
									Sensitive:   true,
								},
								SecretAccessKey: schema.StringAttribute{
									Description: "Secret Access Key for Nodeshift",
									Optional:    true,
									Sensitive:   true,
								},
								SharedCredentialsFile: schema.StringAttribute{
									Description: "Path to credentials file Nodeshift",

									Optional: true,
								},
								Profile: schema.StringAttribute{
									Description: "Nodeshift profile name",
									Optional:    true,
								},
								S3Endpoint: schema.StringAttribute{
									Description: "Nodeshift s3 endpoint address",
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
			t.Parallel()

			p := &nodeshiftProvider{}

			p.Configure(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_nodeshiftProvider_DataSources(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want []func() datasource.DataSource
	}{
		{
			name: "nodeshift provider data sources",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &nodeshiftProvider{}
			assert.Equal(t, tt.want, p.DataSources(context.Background()))
		})
	}
}

func Test_nodeshiftProvider_Metadata(t *testing.T) {
	t.Parallel()

	type args struct {
		in1  provider.MetadataRequest
		resp *provider.MetadataResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nodeshift provider metadata",
			args: args{
				in1:  provider.MetadataRequest{},
				resp: &provider.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &nodeshiftProvider{}
			p.Metadata(context.Background(), tt.args.in1, tt.args.resp)
		})
	}
}

func Test_nodeshiftProvider_Resources(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want []func() resource.Resource
	}{
		{
			name: "nodeshift provider resources",
			want: []func() resource.Resource{
				deployment.NewDeploymentResource,
				vpc.NewVPCResource,
				gpu.NewGPUResource,
				s3terraform.NewBucketResource,
				loadbalancer.NewLBResource,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &nodeshiftProvider{}
			assert.Equal(t, len(tt.want), len(p.Resources(context.Background())))
		})
	}
}

func Test_nodeshiftProvider_Schema(t *testing.T) {
	t.Parallel()

	type args struct {
		in1  provider.SchemaRequest
		resp *provider.SchemaResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nodeshift provider schema",
			args: args{
				in1:  provider.SchemaRequest{},
				resp: &provider.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &nodeshiftProvider{}
			p.Schema(context.Background(), tt.args.in1, tt.args.resp)
		})
	}
}
