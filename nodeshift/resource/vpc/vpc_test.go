package vpc

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewVPCResource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want resource.Resource
	}{
		{
			name: "new vpc resource",
			want: &vpcResource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, NewVPCResource())
		})
	}
}

func Test_vpcResource_Configure(t *testing.T) {
	t.Parallel()

	type args struct {
		req resource.ConfigureRequest
		in2 *resource.ConfigureResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource configure",
			args: args{
				req: resource.ConfigureRequest{
					ProviderData: &client.NodeshiftClient{},
				},
				in2: nil,
			},
		},
		{
			name: "vpc resource configure error",
			args: args{
				req: resource.ConfigureRequest{},
				in2: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vpcResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Configure(t.Context(), tt.args.req, tt.args.in2)
		})
	}
}

func Test_vpcResource_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.CreateRequest
		resp *resource.CreateResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource create",
			args: args{
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.CreateResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
		{
			name: "vpc resource create error",
			args: args{
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.CreateResponse{},
			},
		},
		{
			name: "vpc resource create error convert",
			args: args{
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, IPRangeKeys),
							NameKeys:        tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.CreateResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateVPC(gomock.Any(), gomock.Any()).Return(
				&client.GetVPCResponse{
					UUID:        "vpc-123e4567-e89b-12d3-a456-426614174000",
					Name:        "test-vpc",
					Description: "This is a test VPC",
					IPRange:     "10.0.0.0/24",
					Status:      "running",
					Resources: []client.Resources{
						{
							IP:     "10.0.0.2",
							Status: "running",
						},
						{
							IP:     "10.0.0.3",
							Status: "running",
						},
						{
							IP:     "10.0.0.4",
							Status: "running",
						},
					},
				}, nil,
			).AnyTimes()

			r := &vpcResource{
				client: c,
			}
			r.Create(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.DeleteRequest
		resp *resource.DeleteResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource delete",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, IPRangeKeys),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "vpc resource delete error",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().DeleteVPC(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &vpcResource{
				client: c,
			}
			r.Delete(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_ImportState(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.ImportStateRequest
		resp *resource.ImportStateResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource import state",
			args: args{
				req: resource.ImportStateRequest{
					ID: UUID,
				},
				resp: &resource.ImportStateResponse{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, IPRangeKeys),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vpcResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.ImportState(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Metadata(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.MetadataRequest
		resp *resource.MetadataResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource metadata",
			args: args{
				req: resource.MetadataRequest{
					ProviderTypeName: "type name",
				},
				resp: &resource.MetadataResponse{
					TypeName: "type name",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vpcResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Metadata(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Read(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.ReadRequest
		resp *resource.ReadResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource read",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.ReadResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
		{
			name: "vpc resource read error",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.ReadResponse{},
			},
		},
		{
			name: "vpc resource read convert error",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, IPRangeKeys),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.ReadResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().GetVPC(gomock.Any(), gomock.Any()).Return(
				&client.GetVPCResponse{
					UUID:        "vpc-123e4567-e89b-12d3-a456-426614174000",
					Name:        "test-vpc",
					Description: "This is a test VPC",
					IPRange:     "10.0.0.0/24",
					Status:      "running",
					Resources: []client.Resources{
						{
							IP:     "10.0.0.2",
							Status: "running",
						},
						{
							IP:     "10.0.0.3",
							Status: "running",
						},
						{
							IP:     "10.0.0.4",
							Status: "running",
						},
					},
				}, nil,
			).AnyTimes()

			r := &vpcResource{
				client: c,
			}
			r.Read(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Schema(t *testing.T) {
	t.Parallel()

	type args struct {
		request  resource.SchemaRequest
		response *resource.SchemaResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource schema",
			args: args{
				request: resource.SchemaRequest{},
				response: &resource.SchemaResponse{
					Schema: schema.Schema{
						Description: "Manages a VPC",
						Attributes: map[string]schema.Attribute{
							UUID: schema.StringAttribute{
								Description: "String UUID of the VPC, computed",
								Computed:    true,
							},
							IPRangeKeys: schema.StringAttribute{
								Description: "IP range of the VPC",
								Required:    true,
							},
							NameKeys: schema.StringAttribute{
								Description: "Name of the VPC",
								Required:    true,
							},
							DescriptionKeys: schema.StringAttribute{
								Description: "Description of the VPC",
								Optional:    true,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vpcResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Schema(t.Context(), tt.args.request, tt.args.response)
		})
	}
}

func Test_vpcResource_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		req  resource.UpdateRequest
		resp *resource.UpdateResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vpc resource update",
			args: args{
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.UpdateResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
		{
			name: "vpc resource update error",
			args: args{
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.UpdateResponse{},
			},
		},
		{
			name: "vpc resource update",
			args: args{
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:            tftypes.NewValue(tftypes.String, UUID),
							IPRangeKeys:     tftypes.NewValue(tftypes.String, IPRangeKeys),
							NameKeys:        tftypes.NewValue(tftypes.String, NameKeys),
							DescriptionKeys: tftypes.NewValue(tftypes.String, DescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the VPC, computed",
									Computed:    true,
								},
								IPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								NameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								DescriptionKeys: schema.StringAttribute{
									Description: "Description of the VPC",
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.UpdateResponse{
					State: tfsdk.State{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().UpdateVPC(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				&client.GetVPCResponse{
					UUID:        "vpc-123e4567-e89b-12d3-a456-426614174000",
					Name:        "test-vpc",
					Description: "This is a test VPC",
					IPRange:     "10.0.0.0/24",
					Status:      "running",
					Resources: []client.Resources{
						{
							IP:     "10.0.0.2",
							Status: "running",
						},
						{
							IP:     "10.0.0.3",
							Status: "running",
						},
						{
							IP:     "10.0.0.4",
							Status: "running",
						},
					},
				}, nil,
			).AnyTimes()
			c.EXPECT().GetVPC(gomock.Any(), gomock.Any()).Return(
				&client.GetVPCResponse{
					UUID:        "vpc-123e4567-e89b-12d3-a456-426614174000",
					Name:        "test-vpc",
					Description: "This is a test VPC",
					IPRange:     "10.0.0.0/24",
					Status:      "running",
					Resources: []client.Resources{
						{
							IP:     "10.0.0.2",
							Status: "running",
						},
						{
							IP:     "10.0.0.3",
							Status: "running",
						},
						{
							IP:     "10.0.0.4",
							Status: "running",
						},
					},
				}, nil,
			).AnyTimes()

			r := &vpcResource{
				client: c,
			}
			r.Update(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}
