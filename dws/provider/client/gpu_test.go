package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDWSClient_CreateGPU(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		gpu *GPUConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GPUConfigResponse
	}{
		{
			name: "create gpu",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				gpu: &GPUConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, _ := c.CreateGPU(tt.args.ctx, tt.args.gpu)
			assert.Equalf(t, tt.want, got, "CreateGPU(%v, %v)", tt.args.ctx, tt.args.gpu)
		})
	}
}

func TestDWSClient_DeleteGPU(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
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
			name: "delete gpu",
			fields: fields{
				Config: DWSProviderConfiguration{},
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
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			err := c.DeleteGPU(tt.args.ctx, tt.args.id)
			assert.Errorf(t, err, "failed to delete GPU: external API returned an error code: request failed, status code: 404")
		})
	}
}

func TestDWSClient_GetGPU(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
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
			name: "get gpu",
			fields: fields{
				Config: DWSProviderConfiguration{},
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
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, err := c.GetGPU(tt.args.ctx, tt.args.id)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestDWSClient_UpdateGPU(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		id  string
		gpu *GPUConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GPUConfig
	}{
		{
			name: "update gpu",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				id:  "id",
				gpu: &GPUConfig{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, err := c.UpdateGPU(tt.args.ctx, tt.args.id, tt.args.gpu)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}
