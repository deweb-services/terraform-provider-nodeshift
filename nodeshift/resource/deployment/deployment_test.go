package deployment

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewDeploymentResource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want resource.Resource
	}{
		{
			name: "new deployment resource",
			want: &vmResource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, NewDeploymentResource())
		})
	}
}

func Test_vmResource_Configure(t *testing.T) {
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
			name: "vm resource configure",
			args: args{
				req: resource.ConfigureRequest{
					ProviderData: &client.NodeshiftClient{},
				},
				in2: &resource.ConfigureResponse{},
			},
		},
		{
			name: "vm resource configure error",
			args: args{
				req: resource.ConfigureRequest{},
				in2: &resource.ConfigureResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vmResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Configure(t.Context(), tt.args.req, tt.args.in2)
		})
	}
}

func Test_vmResource_Create(t *testing.T) {
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
			name: "vm resource create",
			args: args{
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
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
			name: "vm resource create schema error",
			args: args{
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw: tftypes.Value{},
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
								},
							},
						},
					},
					Plan: tfsdk.Plan{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
					},
				},
				resp: &resource.CreateResponse{},
			},
		},
		{
			name: "vm resource create convert error",
			args: args{
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
								},
							},
						},
					},
				},
				resp: &resource.CreateResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			c := client.NewMockINodeshiftClient(ctrl)

			c.EXPECT().CreateDeployment(gomock.Any(), gomock.Any()).Return(
				&client.GetDeploymentResponse{
					UUID:         "123e4567-e89b-12d3-a456-426614174000",
					Status:       "running",
					IP:           "192.168.1.10",
					Cru:          4,
					Mru:          8192,
					Sru:          200,
					Hru:          100,
					HddType:      1,
					Provider:     42,
					Hostname:     "test-node.local",
					SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
					SSHKeyName:   "test-key",
					Image:        7,
					ImageVersion: "v1.2.3",
					ChosenPlanID: 101,
					Price:        "12.34",
					CreatedAt:    time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				}, nil,
			).AnyTimes()

			r := &vmResource{
				client: c,
			}
			r.Create(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Delete(t *testing.T) {
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
			name: "vm resource delete",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
								},
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "vm resource delete error",
			args: args{
				req: resource.DeleteRequest{
					State: tfsdk.State{
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

			c.EXPECT().DeleteDeployment(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &vmResource{
				client: c,
			}
			r.Delete(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_ImportState(t *testing.T) {
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
			name: "vm resource import state",
			args: args{
				req: resource.ImportStateRequest{
					ID: "test",
				},
				resp: &resource.ImportStateResponse{
					State: tfsdk.State{
						Schema: schema.Schema{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vmResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.ImportState(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Metadata(t *testing.T) {
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
			name: "vm resource metadata",
			args: args{
				req: resource.MetadataRequest{
					ProviderTypeName: "test",
				},
				resp: &resource.MetadataResponse{
					TypeName: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vmResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Metadata(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Read(t *testing.T) {
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
			name: "vm resource read",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
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
			name: "vm resource read error",
			args: args{
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw:    tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{}),
						Schema: schema.Schema{},
					},
				},
				resp: &resource.ReadResponse{
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

			c.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(
				&client.GetDeploymentResponse{
					UUID:         "123e4567-e89b-12d3-a456-426614174000",
					Status:       "running",
					IP:           "192.168.1.10",
					Cru:          4,
					Mru:          8192,
					Sru:          200,
					Hru:          100,
					HddType:      1,
					Provider:     42,
					Hostname:     "test-node.local",
					SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
					SSHKeyName:   "test-key",
					Image:        7,
					ImageVersion: "v1.2.3",
					ChosenPlanID: 101,
					Price:        "12.34",
					CreatedAt:    time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				}, nil,
			).AnyTimes()

			r := &vmResource{
				client: c,
			}
			r.Read(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Schema(t *testing.T) {
	t.Parallel()

	type args struct {
		in1  resource.SchemaRequest
		resp *resource.SchemaResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vm resource schema",
			args: args{
				in1:  resource.SchemaRequest{},
				resp: &resource.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &vmResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Schema(t.Context(), tt.args.in1, tt.args.resp)
		})
	}
}

func Test_vmResource_Update(t *testing.T) {
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
			name: "vm resource update",
			args: args{
				req: resource.UpdateRequest{
					Config: tfsdk.Config{},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
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
			name: "vm resource update error",
			args: args{
				req: resource.UpdateRequest{
					Config: tfsdk.Config{},
					Plan: tfsdk.Plan{
						Schema: schema.Schema{},
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
			name: "vm resource update convert error",
			args: args{
				req: resource.UpdateRequest{
					Config: tfsdk.Config{},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							UUID:                           tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								UUID: schema.StringAttribute{
									Description: "String UUID of the deployment, computed",
									Computed:    true,
								},
								DeploymentKeysImage: schema.StringAttribute{
									Required:    true,
									Description: ImageDescription,
								},
								DeploymentKeysRegion: schema.StringAttribute{
									Required:    true,
									Description: RegionDescription,
								},
								DeploymentKeysCPU: schema.Int64Attribute{
									Required:    true,
									Description: CPUDescription,
								},
								DeploymentKeysRAM: schema.Int64Attribute{
									Required:    true,
									Description: RAMDescription,
								},
								DeploymentKeysDiskSize: schema.Int64Attribute{
									Required:    true,
									Description: DiskSizeDescription,
								},
								DeploymentKeysDiskType: schema.StringAttribute{
									Required:    true,
									Description: DiskTypeDescription,
								},
								DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv4Description,
								},
								DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignPublicIPv6Description,
								},
								DeploymentKeysSSHKey: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyDescription,
									Sensitive:   true,
								},
								DeploymentKeysSSHKeyName: schema.StringAttribute{
									Required:    true,
									Description: SSHKeyNameDescription,
								},
								DeploymentKeysHostName: schema.StringAttribute{
									Required:    true,
									Description: HostNameDescription,
								},
								DeploymentKeysNetworkUUID: schema.StringAttribute{
									Optional:    true,
									Description: NetworkUUIDDescription,
								},
								DeploymentKeysPublicIPv4: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv4Description,
								},
								DeploymentKeysPublicIPv6: schema.StringAttribute{
									Computed:    true,
									Description: PublicIPv6Description,
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

			c.EXPECT().UpdateDeployment(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				&client.GetDeploymentResponse{
					UUID:         "123e4567-e89b-12d3-a456-426614174000",
					Status:       "running",
					IP:           "192.168.1.10",
					Cru:          4,
					Mru:          8192,
					Sru:          200,
					Hru:          100,
					HddType:      1,
					Provider:     42,
					Hostname:     "test-node.local",
					SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
					SSHKeyName:   "test-key",
					Image:        7,
					ImageVersion: "v1.2.3",
					ChosenPlanID: 101,
					Price:        "12.34",
					CreatedAt:    time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				}, nil,
			).AnyTimes()
			c.EXPECT().GetDeployment(gomock.Any(), gomock.Any()).Return(
				&client.GetDeploymentResponse{
					UUID:         "123e4567-e89b-12d3-a456-426614174000",
					Status:       "running",
					IP:           "192.168.1.10",
					Cru:          4,
					Mru:          8192,
					Sru:          200,
					Hru:          100,
					HddType:      1,
					Provider:     42,
					Hostname:     "test-node.local",
					SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
					SSHKeyName:   "test-key",
					Image:        7,
					ImageVersion: "v1.2.3",
					ChosenPlanID: 101,
					Price:        "12.34",
					CreatedAt:    time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				}, nil,
			).AnyTimes()

			r := &vmResource{
				client: c,
			}
			r.Update(t.Context(), tt.args.req, tt.args.resp)
		})
	}
}
