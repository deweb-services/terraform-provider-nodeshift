package load_balancer

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestLBResourceModel_FromClientRentedLBResponse(t *testing.T) {
	type fields struct {
		Name            types.String
		Replicas        types.Map
		CPUUUIDs        types.List
		ForwardingRules types.List
		VPCUUID         types.String
		UUID            types.String
		Status          types.String
		TaskID          types.String
	}
	type args struct {
		c *client.GetLBResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "lb resource model from client rented lb response",
			fields: fields{
				Name: types.StringValue("my-loadbalancer"),

				Replicas: types.MapValueMust(
					types.Int64Type,
					map[string]attr.Value{
						"replica-1": types.Int64Value(2),
						"replica-2": types.Int64Value(3),
					},
				),

				CPUUUIDs: types.ListValueMust(
					types.StringType,
					[]attr.Value{
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
					},
				),

				ForwardingRules: types.ListValueMust(
					types.ObjectType{
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
					[]attr.Value{
						types.ObjectValueMust(
							map[string]attr.Type{
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
							map[string]attr.Value{
								"in": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTP"),
										"port":     types.Int64Value(80),
									},
								),
								"out": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTPS"),
										"port":     types.Int64Value(443),
									},
								),
							},
						),
					},
				),
				VPCUUID: types.StringValue("vpc-1234-uuid"),
				UUID:    types.StringValue("lb-uuid-5678"),
				Status:  types.StringValue("running"),
				TaskID:  types.StringValue("task-abc-123"),
			},
			args: args{
				c: &client.GetLBResponse{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LBResourceModel{
				Name:            tt.fields.Name,
				Replicas:        tt.fields.Replicas,
				CPUUUIDs:        tt.fields.CPUUUIDs,
				ForwardingRules: tt.fields.ForwardingRules,
				VPCUUID:         tt.fields.VPCUUID,
				UUID:            tt.fields.UUID,
				Status:          tt.fields.Status,
				TaskID:          tt.fields.TaskID,
			}
			if err := m.FromClientRentedLBResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientRentedLBResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLBResourceModel_FromClientResponse(t *testing.T) {
	type fields struct {
		Name            types.String
		Replicas        types.Map
		CPUUUIDs        types.List
		ForwardingRules types.List
		VPCUUID         types.String
		UUID            types.String
		Status          types.String
		TaskID          types.String
	}
	type args struct {
		c *client.LoadBalancerConfigResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "lb resource model from client response",
			fields: fields{
				Name: types.StringValue("my-loadbalancer"),

				Replicas: types.MapValueMust(
					types.Int64Type,
					map[string]attr.Value{
						"replica-1": types.Int64Value(2),
						"replica-2": types.Int64Value(3),
					},
				),

				CPUUUIDs: types.ListValueMust(
					types.StringType,
					[]attr.Value{
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
					},
				),

				ForwardingRules: types.ListValueMust(
					types.ObjectType{
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
					[]attr.Value{
						types.ObjectValueMust(
							map[string]attr.Type{
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
							map[string]attr.Value{
								"in": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTP"),
										"port":     types.Int64Value(80),
									},
								),
								"out": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTPS"),
										"port":     types.Int64Value(443),
									},
								),
							},
						),
					},
				),
				VPCUUID: types.StringValue("vpc-1234-uuid"),
				UUID:    types.StringValue("lb-uuid-5678"),
				Status:  types.StringValue("running"),
				TaskID:  types.StringValue("task-abc-123"),
			},
			args: args{
				c: &client.LoadBalancerConfigResponse{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LBResourceModel{
				Name:            tt.fields.Name,
				Replicas:        tt.fields.Replicas,
				CPUUUIDs:        tt.fields.CPUUUIDs,
				ForwardingRules: tt.fields.ForwardingRules,
				VPCUUID:         tt.fields.VPCUUID,
				UUID:            tt.fields.UUID,
				Status:          tt.fields.Status,
				TaskID:          tt.fields.TaskID,
			}
			if err := m.FromClientResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLBResourceModel_ToClientRequest(t *testing.T) {
	type fields struct {
		Name            types.String
		Replicas        types.Map
		CPUUUIDs        types.List
		ForwardingRules types.List
		VPCUUID         types.String
		UUID            types.String
		Status          types.String
		TaskID          types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.LoadBalancerConfig
		wantErr bool
	}{
		{
			name: "lb resource model to client request",
			fields: fields{
				Name: types.StringValue("my-loadbalancer"),
				Replicas: types.MapValueMust(
					types.Int64Type,
					map[string]attr.Value{
						"replica-1": types.Int64Value(2),
						"replica-2": types.Int64Value(3),
					},
				),
				CPUUUIDs: types.ListValueMust(
					types.StringType,
					[]attr.Value{
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120003"),
						types.StringValue("a3d8e2f0-7a5f-11ec-90d6-0242ac120004"),
					},
				),
				ForwardingRules: types.ListValueMust(
					types.ObjectType{
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
					[]attr.Value{
						types.ObjectValueMust(
							map[string]attr.Type{
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
							map[string]attr.Value{
								"in": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTP"),
										"port":     types.Int64Value(80),
									},
								),
								"out": types.ObjectValueMust(
									map[string]attr.Type{
										"protocol": types.StringType,
										"port":     types.Int64Type,
									},
									map[string]attr.Value{
										"protocol": types.StringValue("HTTPS"),
										"port":     types.Int64Value(443),
									},
								),
							},
						),
					},
				),
				VPCUUID: types.StringValue("vpc-1234-uuid"),
				UUID:    types.StringValue("lb-uuid-5678"),
				Status:  types.StringValue("running"),
				TaskID:  types.StringValue("task-abc-123"),
			},
			want: &client.LoadBalancerConfig{
				Name: "my-loadbalancer",
				Replicas: map[string]int{
					"replica-1": 2,
					"replica-2": 3,
				},
				CPUUUIDs: []string{
					"a3d8e2f0-7a5f-11ec-90d6-0242ac120003",
					"a3d8e2f0-7a5f-11ec-90d6-0242ac120004",
				},
				ForwardingRules: []client.ForwardingRule{
					{
						In: client.RuleEndpoint{
							Protocol: "HTTP",
							Port:     80,
						},
						Out: client.RuleEndpoint{
							Protocol: "HTTPS",
							Port:     443,
						},
					},
				},
				VPCUUID: "vpc-1234-uuid",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LBResourceModel{
				Name:            tt.fields.Name,
				Replicas:        tt.fields.Replicas,
				CPUUUIDs:        tt.fields.CPUUUIDs,
				ForwardingRules: tt.fields.ForwardingRules,
				VPCUUID:         tt.fields.VPCUUID,
				UUID:            tt.fields.UUID,
				Status:          tt.fields.Status,
				TaskID:          tt.fields.TaskID,
			}
			got, err := m.ToClientRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClientRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToClientRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
