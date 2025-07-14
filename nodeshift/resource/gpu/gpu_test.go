package gpu

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewGPUResource(t *testing.T) {
	tests := []struct {
		name string
		want resource.Resource
	}{
		{
			name: "new gpu resource",
			want: &gpuResource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGPUResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGPUResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gpuResource_Configure(t *testing.T) {
	type fields struct {
		client *client.NodeshiftClient
	}
	type args struct {
		in0 context.Context
		req resource.ConfigureRequest
		in2 *resource.ConfigureResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "vm resource configure",
			fields: fields{
				client: &client.NodeshiftClient{},
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{
					ProviderData: &client.NodeshiftClient{},
				},
				in2: &resource.ConfigureResponse{},
			},
		},
		{
			name: "vm resource configure error",
			fields: fields{
				client: &client.NodeshiftClient{},
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{},
				in2: &resource.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Configure(tt.args.in0, tt.args.req, tt.args.in2)
		})
	}
}

func Test_gpuResource_Create(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		ctx  context.Context
		req  resource.CreateRequest
		resp *resource.CreateResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource create",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			name: "gpu resource create error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.CreateResponse{},
			},
		},
		{
			name: "gpu resource create error convert",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Delete(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		ctx  context.Context
		req  resource.DeleteRequest
		resp *resource.DeleteResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource delete",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			name: "gpu resource delete error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "gpu resource delete convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:    tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
									Optional:    true,
								},
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Delete(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_ImportState(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		ctx  context.Context
		req  resource.ImportStateRequest
		resp *resource.ImportStateResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource import state",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ImportStateRequest{},
				resp: &resource.ImportStateResponse{
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
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.ImportState(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Metadata(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		in0  context.Context
		req  resource.MetadataRequest
		resp *resource.MetadataResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource metadata",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				in0:  context.TODO(),
				req:  resource.MetadataRequest{},
				resp: &resource.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Metadata(tt.args.in0, tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Read(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		ctx  context.Context
		req  resource.ReadRequest
		resp *resource.ReadResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource read",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			name: "gpu resource read error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
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
			name: "gpu resource read error convert",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Read(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Schema(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		c        context.Context
		request  resource.SchemaRequest
		response *resource.SchemaResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource schema",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				c:        context.TODO(),
				request:  resource.SchemaRequest{},
				response: &resource.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Schema(tt.args.c, tt.args.request, tt.args.response)
		})
	}
}

func Test_gpuResource_Update(t *testing.T) {
	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
		ctx  context.Context
		req  resource.UpdateRequest
		resp *resource.UpdateResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gpu resource update",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			name: "gpu resource update error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.UpdateResponse{},
			},
		},
		{
			name: "gpu resource update error convert",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:        tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:  tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:    tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:   tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount: tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:   tftypes.NewValue(tftypes.String, KeyRegion),
						}),
						Schema: schema.Schema{
							Description: "Manages a GPU",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyGPUName: schema.StringAttribute{
									Description: DescriptionGPUName,
									Required:    true,
								},
								KeyImage: schema.StringAttribute{
									Description: DescriptionImage,
									Required:    true,
								},
								KeySSHKey: schema.StringAttribute{
									Description: DescriptionSSHKey,
									Required:    true,
								},
								KeyGPUCount: schema.Int64Attribute{
									Description: DescriptionGPUCount,
									Optional:    true,
								},
								KeyRegion: schema.StringAttribute{
									Description: DescriptionRegion,
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
			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Update(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
