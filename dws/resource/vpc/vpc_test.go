package vpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestNewVPCResource(t *testing.T) {
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
			if got := NewVPCResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVPCResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vpcResource_Configure(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource configure",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{
					ProviderData: &client.DWSClient{},
				},
				in2: nil,
			},
		},
		{
			name: "vpc resource configure error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{},
				in2: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Configure(tt.args.in0, tt.args.req, tt.args.in2)
		})
	}
}

func Test_vpcResource_Create(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource create",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, VPCIPRangeKeys),
							VPCNameKeys:        tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Delete(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource delete",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, VPCIPRangeKeys),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Delete(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_ImportState(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name:   "vpc resource import state",
			fields: fields{},
			args: args{
				ctx: context.TODO(),
				req: resource.ImportStateRequest{
					ID: ID,
				},
				resp: &resource.ImportStateResponse{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, VPCIPRangeKeys),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.ImportState(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Metadata(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name:   "vpc resource metadata",
			fields: fields{},
			args: args{
				in0: context.TODO(),
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Metadata(tt.args.in0, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Read(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource read",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			name: "vpc resource read convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, VPCIPRangeKeys),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Read(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vpcResource_Schema(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource schema",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				c:       context.TODO(),
				request: resource.SchemaRequest{},
				response: &resource.SchemaResponse{
					Schema: schema.Schema{
						Description: "Manages a VPC",
						Attributes: map[string]schema.Attribute{
							ID: schema.StringAttribute{
								Description: "String ID of the VPC, computed",
								Computed:    true,
							},
							VPCIPRangeKeys: schema.StringAttribute{
								Description: "IP range of the VPC",
								Required:    true,
							},
							VPCNameKeys: schema.StringAttribute{
								Description: "Name of the VPC",
								Required:    true,
							},
							VPCDescriptionKeys: schema.StringAttribute{
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Schema(tt.args.c, tt.args.request, tt.args.response)
		})
	}
}

func Test_vpcResource_Update(t *testing.T) {
	type fields struct {
		client client.IDWSClient
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
			name: "vpc resource update",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, "127.0.0.1/24"),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                 tftypes.NewValue(tftypes.String, ID),
							VPCIPRangeKeys:     tftypes.NewValue(tftypes.String, VPCIPRangeKeys),
							VPCNameKeys:        tftypes.NewValue(tftypes.String, VPCNameKeys),
							VPCDescriptionKeys: tftypes.NewValue(tftypes.String, VPCDescriptionKeys),
						}),
						Schema: schema.Schema{
							Description: "Manages a VPC",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the VPC, computed",
									Computed:    true,
								},
								VPCIPRangeKeys: schema.StringAttribute{
									Description: "IP range of the VPC",
									Required:    true,
								},
								VPCNameKeys: schema.StringAttribute{
									Description: "Name of the VPC",
									Required:    true,
								},
								VPCDescriptionKeys: schema.StringAttribute{
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
			r := &vpcResource{
				client: tt.fields.client,
			}
			r.Update(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
