package deployment

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

func TestNewDeploymentResource(t *testing.T) {
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
			if got := NewDeploymentResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeploymentResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vmResource_Configure(t *testing.T) {
	type fields struct {
		client *client.DWSClient
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
				client: &client.DWSClient{},
			},
			args: args{
				in0: context.TODO(),
				req: resource.ConfigureRequest{
					ProviderData: &client.DWSClient{},
				},
				in2: &resource.ConfigureResponse{},
			},
		},
		{
			name: "vm resource configure error",
			fields: fields{
				client: &client.DWSClient{},
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Configure(tt.args.in0, tt.args.req, tt.args.in2)
		})
	}
}

func Test_vmResource_Create(t *testing.T) {
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
			name: "vm resource create",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
			fields: fields{
				client: &client.DWSClient{},
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Config: tfsdk.Config{
						Raw:    tftypes.Value{},
						Schema: schema.Schema{},
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.CreateRequest{
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Delete(t *testing.T) {
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
			name: "vm resource delete",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, ""),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
			fields: fields{
				client: &client.DWSClient{
					Config: client.DWSProviderConfiguration{},
				},
			},
			args: args{
				ctx: context.TODO(),
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Delete(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_ImportState(t *testing.T) {
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
			name: "vm resource import state",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.ImportState(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Metadata(t *testing.T) {
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
			name: "vm resource metadata",
			fields: fields{
				client: &client.DWSClient{},
			},
			args: args{
				in0: context.TODO(),
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Metadata(tt.args.in0, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Read(t *testing.T) {
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
			name: "vm resource read",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
				resp: &resource.ReadResponse{
					State: tfsdk.State{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Read(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_vmResource_Schema(t *testing.T) {
	type fields struct {
		client *client.DWSClient
	}
	type args struct {
		in0  context.Context
		in1  resource.SchemaRequest
		resp *resource.SchemaResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "vm resource schema",
			fields: fields{
				client: &client.DWSClient{},
			},
			args: args{
				in0:  context.TODO(),
				in1:  resource.SchemaRequest{},
				resp: &resource.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Schema(tt.args.in0, tt.args.in1, tt.args.resp)
		})
	}
}

func Test_vmResource_Update(t *testing.T) {
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
			name: "vm resource update",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Config: tfsdk.Config{},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.String, DeploymentKeysImage),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
			fields: fields{
				client: &client.DWSClient{},
			},
			args: args{
				ctx: context.TODO(),
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
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.UpdateRequest{
					Config: tfsdk.Config{},
					Plan: tfsdk.Plan{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							ID:                             tftypes.NewValue(tftypes.String, "id"),
							DeploymentKeysImage:            tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
							DeploymentKeysRegion:           tftypes.NewValue(tftypes.String, DeploymentKeysRegion),
							DeploymentKeysCPU:              tftypes.NewValue(tftypes.Number, 1),
							DeploymentKeysRAM:              tftypes.NewValue(tftypes.Number, 2),
							DeploymentKeysDiskSize:         tftypes.NewValue(tftypes.Number, 3),
							DeploymentKeysDiskType:         tftypes.NewValue(tftypes.String, DeploymentKeysDiskType),
							DeploymentKeysAssignPublicIPv4: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignPublicIPv6: tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysAssignYggIP:      tftypes.NewValue(tftypes.Bool, false),
							DeploymentKeysSSHKey:           tftypes.NewValue(tftypes.String, DeploymentKeysSSHKey),
							DeploymentKeysSSHKeyName:       tftypes.NewValue(tftypes.String, DeploymentKeysSSHKeyName),
							DeploymentKeysHostName:         tftypes.NewValue(tftypes.String, DeploymentKeysHostName),
							DeploymentKeysNetworkUUID:      tftypes.NewValue(tftypes.String, DeploymentKeysNetworkUUID),
							DeploymentKeysPublicIPv4:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv4),
							DeploymentKeysPublicIPv6:       tftypes.NewValue(tftypes.String, DeploymentKeysPublicIPv6),
							DeploymentKeysYggIP:            tftypes.NewValue(tftypes.String, DeploymentKeysYggIP),
						}),
						Schema: schema.Schema{
							Description: "Manages a deployment",
							Attributes: map[string]schema.Attribute{
								ID: schema.StringAttribute{
									Description: "String ID of the deployment, computed",
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
								DeploymentKeysAssignYggIP: schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: AssignYggIPDescription,
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
								DeploymentKeysYggIP: schema.StringAttribute{
									Computed:    true,
									Description: YggIPDescription,
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
			r := &vmResource{
				client: tt.fields.client,
			}
			r.Update(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
