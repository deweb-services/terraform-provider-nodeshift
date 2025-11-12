package gpu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewGPUResource(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			assert.Equal(t, tt.want, NewGPUResource())
		})
	}
}

func Test_gpuResource_Configure(t *testing.T) {
	t.Parallel()

	type fields struct {
		client *client.NodeshiftClient
	}
	type args struct {
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
				req: resource.ConfigureRequest{},
				in2: &resource.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &gpuResource{
				client: tt.fields.client,
			}
			r.Configure(t.Context(), tt.args.req, tt.args.in2)
		})
	}
}

func Test_gpuResource_Create(t *testing.T) {
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
			name: "gpu resource create",
			args: args{
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                  tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:            tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:              tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:             tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:           tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:             tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:         tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion:     tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
							KeyMachineTypeVersion: tftypes.NewValue(tftypes.String, "vm"),
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
									Computed:    true,
								},
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
									Optional:    true,
								},
								KeyMachineTypeVersion: schema.StringAttribute{
									Description: DescriptionMachineTypeVersion,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOf(availableMachineTypes...),
									},
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
			args: args{
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
			args: args{
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                  tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:            tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:              tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:             tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:           tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:             tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:         tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion:     tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
							KeyMachineTypeVersion: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
									Optional:    true,
								},
								KeyMachineTypeVersion: schema.StringAttribute{
									Description: DescriptionMachineTypeVersion,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOf(availableMachineTypes...),
									},
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

			c := client.NewMockINodeshiftClient(gomock.NewController(t))
			c.EXPECT().CreateGPU(gomock.Any(), gomock.Any()).Return(&client.GetGPUResponse{
				UUID:    "gpu-123e4567-e89b-12d3-a456-426614174000",
				GpuName: "NVIDIA Tesla V100",
				NumGpus: 2,
				SSHHost: "192.168.1.50",
				SSHPort: 22,
				Status:  "running",
			}, nil).AnyTimes()

			r := &gpuResource{
				client: c,
			}
			r.Create(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Delete(t *testing.T) {
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
			name: "gpu resource delete",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:          tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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
			name: "gpu resource delete convert error",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:          tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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
			t.Parallel()

			c := client.NewMockINodeshiftClient(gomock.NewController(t))
			c.EXPECT().DeleteGPU(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &gpuResource{
				client: c,
			}
			r.Delete(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_ImportState(t *testing.T) {
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
			name: "gpu resource import state",
			args: args{
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
			t.Parallel()

			r := &gpuResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.ImportState(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Metadata(t *testing.T) {
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
			name: "gpu resource metadata",
			args: args{
				req:  resource.MetadataRequest{},
				resp: &resource.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &gpuResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Metadata(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Read(t *testing.T) {
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
			name: "gpu resource read",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:          tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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
			name: "gpu resource read error convert",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:          tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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

			c := client.NewMockINodeshiftClient(gomock.NewController(t))
			c.EXPECT().GetGPU(gomock.Any(), gomock.Any()).Return(&client.GetGPUResponse{
				UUID:    "gpu-123e4567-e89b-12d3-a456-426614174000",
				GpuName: "NVIDIA Tesla V100",
				NumGpus: 2,
				SSHHost: "192.168.1.50",
				SSHPort: 22,
				Status:  "running",
			}, nil).AnyTimes()

			r := &gpuResource{
				client: c,
			}
			r.Read(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_gpuResource_Schema(t *testing.T) {
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
			name: "gpu resource schema",
			args: args{
				request:  resource.SchemaRequest{},
				response: &resource.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &gpuResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Schema(t.Context(), tt.args.request, tt.args.response)
		})
	}
}

func Test_gpuResource_Update(t *testing.T) {
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
			name: "gpu resource update",
			args: args{
				req: resource.UpdateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.String, KeyGPUName),
							KeyImage:          tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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
			args: args{
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
			args: args{
				req: resource.UpdateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:              tftypes.NewValue(tftypes.String, UUID),
							KeyGPUName:        tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							KeyImage:          tftypes.NewValue(tftypes.String, KeyImage),
							KeySSHKey:         tftypes.NewValue(tftypes.String, KeySSHKey),
							KeyGPUCount:       tftypes.NewValue(tftypes.Number, 2),
							KeyRegion:         tftypes.NewValue(tftypes.String, KeyRegion),
							KeyDiskSizeGB:     tftypes.NewValue(tftypes.Number, 30),
							KeyMinCudaVersion: tftypes.NewValue(tftypes.String, KeyMinCudaVersion),
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
								KeyDiskSizeGB: schema.Int64Attribute{
									Description: DescriptionDiskSizeGB,
									Optional:    true,
								},
								KeyMinCudaVersion: schema.StringAttribute{
									Description: DescriptionMinCudaVersion,
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

			c := client.NewMockINodeshiftClient(gomock.NewController(t))
			c.EXPECT().GetGPU(gomock.Any(), gomock.Any()).Return(&client.GetGPUResponse{
				UUID:    "gpu-123e4567-e89b-12d3-a456-426614174000",
				GpuName: "NVIDIA Tesla V100",
				NumGpus: 2,
				SSHHost: "192.168.1.50",
				SSHPort: 22,
				Status:  "running",
			}, nil).AnyTimes()
			c.EXPECT().UpdateGPU(gomock.Any(), gomock.Any(), gomock.Any()).Return(&client.GetGPUResponse{
				UUID:    "gpu-123e4567-e89b-12d3-a456-426614174000",
				GpuName: "NVIDIA Tesla V100",
				NumGpus: 2,
				SSHHost: "192.168.1.50",
				SSHPort: 22,
				Status:  "running",
			}, nil).AnyTimes()

			r := &gpuResource{
				client: c,
			}
			r.Update(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}
