package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeshiftClient_CreateLB(t *testing.T) {
	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		lb  *LoadBalancerConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LoadBalancerConfigResponse
	}{
		{
			name: "create lb",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				lb:  &LoadBalancerConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, _ := c.CreateLB(tt.args.ctx, tt.args.lb)
			assert.Equalf(t, tt.want, got, "CreateLB(%v, %v)", tt.args.ctx, tt.args.lb)
		})
	}
}

func TestNodeshiftClient_DeleteLB(t *testing.T) {
	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete lb",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				id:  "id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			err := c.DeleteLB(tt.args.ctx, tt.args.id)
			assert.Errorf(t, err, "failed to delete LB: external API returned an error code: request failed, status code: 404")
		})
	}
}

func TestNodeshiftClient_GetLB(t *testing.T) {
	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RentedGpuInfoResponse
	}{
		{
			name: "get lb",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				id:  "id",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, err := c.GetLB(tt.args.ctx, tt.args.id)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestNodeshiftClient_UpdateLB(t *testing.T) {
	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		id  string
		lb  *LoadBalancerConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LoadBalancerConfigResponse
	}{
		{
			name: "update lb",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				id:  "id",
				lb:  &LoadBalancerConfig{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, err := c.UpdateLB(tt.args.ctx, tt.args.id, tt.args.lb)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}
