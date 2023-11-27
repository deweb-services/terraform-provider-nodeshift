package s3

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestNewBucketResource(t *testing.T) {
	tests := []struct {
		name string
		want resource.Resource
	}{
		{
			name: "new bucket resource",
			want: &bucketResource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBucketResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBucketResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bucketResource_Configure(t *testing.T) {
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
			name: "bucket resource configure",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{},
				in2: &resource.ConfigureResponse{},
			},
		},
		{
			name: "bucket resource configure",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{
					ProviderData: &client.DWSClient{},
				},
				in2: &resource.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Configure(tt.args.in0, tt.args.req, tt.args.in2)
		})
	}
}

func Test_bucketResource_Create(t *testing.T) {
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
			name: "bucket resource create",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, KeyBucketName),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			name: "bucket resource create error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
				resp: &resource.CreateResponse{
					State: tfsdk.State{},
				},
			},
		},
		{
			name: "bucket resource create convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
								},
							},
						},
					},
				},
				resp: &resource.CreateResponse{
					State: tfsdk.State{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Delete(t *testing.T) {
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
			name: "bucket resource delete",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, KeyBucketName),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
								},
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "bucket resource delete error",
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
			name: "bucket resource delete convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Delete(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_ImportState(t *testing.T) {
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
			name: "bucket resource import state",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ImportStateRequest{
					ID: "id",
				},
				resp: &resource.ImportStateResponse{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, KeyBucketName),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.ImportState(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Metadata(t *testing.T) {
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
			name: "bucket resource metadata",
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
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Metadata(tt.args.in0, tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Read(t *testing.T) {
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
			name: "bucket resource read",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, KeyBucketName),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			name: "bucket resource read error",
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
			name: "bucket resource read convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Read(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Schema(t *testing.T) {
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
			name: "bucket resource schema",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				c:       nil,
				request: resource.SchemaRequest{},
				response: &resource.SchemaResponse{
					Schema: schema.Schema{
						Description: "Manages a s3 Bucket",
						Attributes: map[string]schema.Attribute{
							KeyBucketName: schema.StringAttribute{
								Description: DescriptionBucketName,
								Required:    true,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Schema(tt.args.c, tt.args.request, tt.args.response)
		})
	}
}

func Test_bucketResource_Update(t *testing.T) {
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
			name: "bucket resource update",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.String, KeyBucketName),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			name: "bucket resource update error",
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
			name: "bucket resource update convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyBucketName: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
						}),
						Schema: schema.Schema{
							Description: "Manages a s3 Bucket",
							Attributes: map[string]schema.Attribute{
								KeyBucketName: schema.StringAttribute{
									Description: DescriptionBucketName,
									Required:    true,
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
			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Update(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
