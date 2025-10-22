package loadbalancer

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestNewLBResource(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			assert.Equal(t, tt.want, NewLBResource())
		})
	}
}

func Test_lbResource_Configure(t *testing.T) {
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

			r := &lbResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Configure(context.Background(), tt.args.req, tt.args.in2)
		})
	}
}

func Test_lbResource_Create(t *testing.T) {
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
			name: "lb resource create",
			args: args{
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
			name: "lb resource create error convert",
			args: args{
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

			c.EXPECT().CreateLB(gomock.Any(), gomock.Any()).Return(
				&client.GetLBResponse{
					UUID:           "lb-123e4567-e89b-12d3-a456-426614174000",
					Name:           "test-lb",
					Status:         "running",
					ReplicasAmount: 2,
					CPUAmount:      8,
					PriceInUSD:     "123.45",
					CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				},
				nil,
			).AnyTimes()

			r := &lbResource{client: c}
			r.Create(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Delete(t *testing.T) {
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
			name: "lb resource delete",
			args: args{
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
							},
						},
					},
				},
				resp: &resource.DeleteResponse{},
			},
		},
		{
			name: "lb resource delete error",
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
			name: "lb resource delete convert error",
			args: args{
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

			c.EXPECT().DeleteLB(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			r := &lbResource{client: c}
			r.Delete(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_ImportState(t *testing.T) {
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
			name: "lb resource import state",
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

			r := &lbResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.ImportState(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Metadata(t *testing.T) {
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
			name: "lb resource metadata",
			args: args{
				req:  resource.MetadataRequest{},
				resp: &resource.MetadataResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &lbResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Metadata(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Read(t *testing.T) {
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
			name: "lb resource read",
			args: args{
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
			name: "lb resource read error convert",
			args: args{
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

			c.EXPECT().GetLB(gomock.Any(), gomock.Any()).Return(
				&client.GetLBResponse{
					UUID:           "lb-123e4567-e89b-12d3-a456-426614174000",
					Name:           "test-lb",
					Status:         "running",
					ReplicasAmount: 2,
					CPUAmount:      8,
					PriceInUSD:     "123.45",
					CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				},
				nil,
			).AnyTimes()

			r := &lbResource{client: c}
			r.Read(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}

func Test_lbResource_Schema(t *testing.T) {
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
			name: "lb resource schema",
			args: args{
				request:  resource.SchemaRequest{},
				response: &resource.SchemaResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := &lbResource{
				client: client.NewMockINodeshiftClient(gomock.NewController(t)),
			}
			r.Schema(context.Background(), tt.args.request, tt.args.response)
		})
	}
}

func Test_lbResource_Update(t *testing.T) {
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
			name: "lb resource update",
			args: args{
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
			name: "lb resource update error convert",
			args: args{
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

			c.EXPECT().UpdateLB(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				&client.GetLBResponse{
					UUID:           "lb-123e4567-e89b-12d3-a456-426614174000",
					Name:           "test-lb",
					Status:         "running",
					ReplicasAmount: 2,
					CPUAmount:      8,
					PriceInUSD:     "123.45",
					CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				},
				nil,
			).AnyTimes()

			c.EXPECT().GetLB(gomock.Any(), gomock.Any()).Return(
				&client.GetLBResponse{
					UUID:           "lb-123e4567-e89b-12d3-a456-426614174000",
					Name:           "test-lb",
					Status:         "running",
					ReplicasAmount: 2,
					CPUAmount:      8,
					PriceInUSD:     "123.45",
					CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				},
				nil,
			).AnyTimes()

			r := &lbResource{client: c}
			r.Update(context.Background(), tt.args.req, tt.args.resp)
		})
	}
}
