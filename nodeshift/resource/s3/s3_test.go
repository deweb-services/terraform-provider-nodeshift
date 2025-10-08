package s3

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"go.uber.org/mock/gomock"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewBucketResource(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			if got := NewBucketResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBucketResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bucketResource_Configure(t *testing.T) {
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
			name: "bucket resource configure",
			args: args{
				req: resource.ConfigureRequest{},
				in2: &resource.ConfigureResponse{},
			},
		},
		{
			name: "bucket resource configure",
			args: args{
				req: resource.ConfigureRequest{
					ProviderData: &client.NodeshiftClient{},
				},
				in2: &resource.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Configure(context.Background(), tt.args.req, tt.args.in2)
		})
	}
}

func Test_bucketResource_Create(t *testing.T) {
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
			name: "bucket resource create",
			args: args{
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
			args: args{
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
			args: args{
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
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Create(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Delete(t *testing.T) {
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
			name: "bucket resource delete",
			args: args{
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
		{
			name: "bucket resource delete convert error",
			args: args{
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
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().DeleteBucket(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Delete(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_ImportState(t *testing.T) {
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
			name: "bucket resource import state",
			args: args{
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
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.ImportState(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Metadata(t *testing.T) {
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
			name: "bucket resource metadata",
			args: args{
				req:  resource.MetadataRequest{},
				resp: &resource.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Metadata(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Read(t *testing.T) {
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
			name: "bucket resource read",
			args: args{
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
			name: "bucket resource read convert error",
			args: args{
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
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().GetBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Read(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_bucketResource_Schema(t *testing.T) {
	t.Parallel()

	type fields struct {
		client client.INodeshiftClient
	}
	type args struct {
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
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			},
			args: args{
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
			t.Parallel()

			r := &bucketResource{
				client: tt.fields.client,
			}
			r.Schema(context.Background(), tt.args.request, tt.args.response)
		})
	}
}

func Test_bucketResource_Update(t *testing.T) {
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
			name: "bucket resource update",
			args: args{
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
			name: "bucket resource update convert error",
			args: args{
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
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().GetBucket(gomock.Any(), gomock.Any()).Return(
				&client.S3BucketConfig{Key: "bucket"},
				nil,
			).AnyTimes()
			c.EXPECT().UpdateBucket(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &bucketResource{
				client: c,
			}
			r.Update(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}
