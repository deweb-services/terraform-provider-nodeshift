package load_balancer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestNewLBResource(t *testing.T) {
	tests := []struct {
		name string
		want resource.Resource
	}{
		{
			name: "new lb resource",
			want: &lbResource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLBResource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLBResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lbResource_Configure(t *testing.T) {
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Configure(tt.args.in0, tt.args.req, tt.args.in2)
		})
	}
}

func Test_lbResource_Create(t *testing.T) {
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
			name: "lb resource create",
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
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			name: "lb resource create error",
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
			name: "lb resource create error convert",
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
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Create(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Delete(t *testing.T) {
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
			name: "lb resource delete",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
								},
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "lb resource delete error",
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
			name: "lb resource delete convert error",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.DeleteRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Delete(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_ImportState(t *testing.T) {
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
			name: "lb resource import state",
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.ImportState(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Metadata(t *testing.T) {
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
			name: "lb resource metadata",
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Metadata(tt.args.in0, tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Read(t *testing.T) {
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
			name: "lb resource read",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			name: "lb resource read error",
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
			name: "lb resource read error convert",
			fields: fields{
				client: client.NewMockedClient(),
			},
			args: args{
				ctx: context.TODO(),
				req: resource.ReadRequest{
					State: tfsdk.State{
						Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Read(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Schema(t *testing.T) {
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
			name: "lb resource schema",
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Schema(tt.args.c, tt.args.request, tt.args.response)
		})
	}
}

func Test_lbResource_Update(t *testing.T) {
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
			name: "lb resource update",
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
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			name: "lb resource update error",
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
			name: "lb resource update error convert",
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
							KeyName: tftypes.NewValue(tftypes.String, "loadbalancer-1"),
							KeyReplicas: tftypes.NewValue(tftypes.Map{ElementType: tftypes.Number}, map[string]tftypes.Value{
								"replica1": tftypes.NewValue(tftypes.Number, 2),
								"replica2": tftypes.NewValue(tftypes.Number, 3),
							}),
							KeyCPUUUIDs: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
								tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
							}),
							KeyForwardingRules: tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									},
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.Object{
										AttributeTypes: map[string]tftypes.Type{
											"in": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
											"out": tftypes.Object{
												AttributeTypes: map[string]tftypes.Type{
													"protocol": tftypes.String,
													"port":     tftypes.Number,
												},
											},
										},
									}, map[string]tftypes.Value{
										"in": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTP"),
											"port":     tftypes.NewValue(tftypes.Number, 80),
										}),
										"out": tftypes.NewValue(tftypes.Object{
											AttributeTypes: map[string]tftypes.Type{
												"protocol": tftypes.String,
												"port":     tftypes.Number,
											},
										}, map[string]tftypes.Value{
											"protocol": tftypes.NewValue(tftypes.String, "HTTPS"),
											"port":     tftypes.NewValue(tftypes.Number, 443),
										}),
									}),
								},
							),
							KeyVPCUUID: tftypes.NewValue(tftypes.String, "a3d8e2f0-7a5f-11ec-90d6-0242ac120005"),

							UUID:      tftypes.NewValue(tftypes.String, "some-uuid"),
							KeyStatus: tftypes.NewValue(tftypes.String, "creating"),
							KeyTaskId: tftypes.NewValue(tftypes.String, "task-xyz"),
						}),
						Schema: schema.Schema{
							Description: "Manages a LB",
							Attributes: map[string]schema.Attribute{
								KeyName: schema.StringAttribute{
									Description: DescriptionName,
									Required:    true,
								},
								KeyReplicas: schema.MapAttribute{
									Description: DescriptionReplicas,
									Required:    true,
									ElementType: types.Int64Type,
								},
								KeyCPUUUIDs: schema.ListAttribute{
									Description: DescriptionCPUUUIDs,
									Required:    true,
									ElementType: types.StringType,
								},
								KeyForwardingRules: schema.ListAttribute{
									Description: DescriptionForwardingRules,
									Required:    true,
									ElementType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"in": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
											"out": types.ObjectType{
												AttrTypes: map[string]attr.Type{
													"protocol": types.StringType,
													"port":     types.Int64Type,
												},
											},
										},
									},
								},
								KeyVPCUUID: schema.StringAttribute{
									Description: DescriptionVPCUUID,
									Required:    true,
								},

								UUID: schema.StringAttribute{
									Description: DescriptionUUID,
									Computed:    true,
								},
								KeyStatus: schema.StringAttribute{
									Description: DescriptionStatus,
									Computed:    true,
								},
								KeyTaskId: schema.StringAttribute{
									Description: DescriptionTaskId,
									Computed:    true,
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
			r := &lbResource{
				client: tt.fields.client,
			}
			r.Update(tt.args.ctx, tt.args.req, tt.args.resp)
		})
	}
}
