package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDWSClient_CreateVPC(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		vpc *VPCConfig
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VPCConfig
	}{
		{
			name: "create_vpc",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				vpc: &VPCConfig{
					ID:          "",
					Name:        "",
					Description: "",
					IPRange:     "",
				},
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
			got, err := c.CreateVPC(tt.args.ctx, tt.args.vpc)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestDWSClient_DeleteVPC(t *testing.T) {
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
			name: "delete_vpc",
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
			err := c.DeleteVPC(tt.args.ctx, tt.args.id)
			assert.NotNil(t, err)
		})
	}
}

func TestDWSClient_GetVPC(t *testing.T) {
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
		want   *VPCConfig
	}{
		{
			name: "get_vpc",
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
			want: &VPCConfig{},
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
			got, err := c.GetVPC(tt.args.ctx, tt.args.id)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestDWSClient_UpdateVPC(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx context.Context
		id  string
		vpc *VPCConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VPCConfig
	}{
		{
			name: "update_vpc",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx: context.TODO(),
				id:  "id",
				vpc: &VPCConfig{
					ID:          "id",
					Name:        "name",
					Description: "description",
					IPRange:     "127.0.0.1/24",
				},
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
			got, err := c.UpdateVPC(tt.args.ctx, tt.args.id, tt.args.vpc)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}
