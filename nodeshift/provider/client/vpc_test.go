package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateVPC(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		url    string
	}
	type args struct {
		vpc *CreateVPCRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateVPCRequest
	}{
		{
			name: "create_vpc",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				url:    exampleURLString,
			},
			args: args{
				vpc: &CreateVPCRequest{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.CreateVPC(context.Background(), tt.args.vpc)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

func Test_DeleteVPC(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		url    string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete_vpc",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				url:    exampleURLString,
			},
			args: args{
				id: "id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			err := c.DeleteVPC(context.Background(), tt.args.id)
			require.Error(t, err)
		})
	}
}

func Test_GetVPC(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		url    string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateVPCRequest
	}{
		{
			name: "get_vpc",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				url:    exampleURLString,
			},
			args: args{
				id: "id",
			},
			want: &CreateVPCRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.GetVPC(context.Background(), tt.args.id)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

func Test_UpdateVPC(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config NodeshiftProviderConfiguration
		client *http.Client
		url    string
	}
	type args struct {
		id  string
		vpc *CreateVPCRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateVPCRequest
	}{
		{
			name: "update_vpc",
			fields: fields{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				url:    exampleURLString,
			},
			args: args{
				id: "id",
				vpc: &CreateVPCRequest{
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
			t.Parallel()

			c := &NodeshiftClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.UpdateVPC(context.Background(), tt.args.id, tt.args.vpc)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}
