package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeshiftClient_CreateLB(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		lb *CreateLBRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetLBResponse
	}{
		{
			name: "create lb",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				lb: &CreateLBRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, _ := c.CreateLB(context.Background(), tt.args.lb)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNodeshiftClient_DeleteLB(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
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
			name: "delete lb",
			fields: fields{
				url: exampleURLString,
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
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			err := c.DeleteLB(context.Background(), tt.args.id)
			assert.Errorf(t, err, "failed to delete LB: external API returned an error code: request failed, status code: 404")
		})
	}
}

func TestNodeshiftClient_GetLB(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetGPUResponse
	}{
		{
			name: "get lb",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				id: "id",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.GetLB(context.Background(), tt.args.id)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestNodeshiftClient_UpdateLB(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		id string
		lb *CreateLBRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetLBResponse
	}{
		{
			name: "update lb",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				id: "id",
				lb: &CreateLBRequest{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.UpdateLB(context.Background(), tt.args.id, tt.args.lb)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}
