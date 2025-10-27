package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeshiftClient_CreateGPU(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		gpu *CreateGPURequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetGPUResponse
	}{
		{
			name: "create gpu",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				gpu: &CreateGPURequest{},
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
			got, _ := c.CreateGPU(t.Context(), tt.args.gpu)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNodeshiftClient_DeleteGPU(t *testing.T) {
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
			name: "delete gpu",
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
			err := c.DeleteGPU(t.Context(), tt.args.id)
			assert.Errorf(t, err, "failed to delete GPU: external API returned an error code: request failed, status code: 404")
		})
	}
}

func TestNodeshiftClient_GetGPU(t *testing.T) {
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
			name: "get gpu",
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
			got, err := c.GetGPU(t.Context(), tt.args.id)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestNodeshiftClient_UpdateGPU(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		id  string
		gpu *CreateGPURequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateGPURequest
	}{
		{
			name: "update gpu",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				id:  "id",
				gpu: &CreateGPURequest{},
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
			got, err := c.UpdateGPU(t.Context(), tt.args.id, tt.args.gpu)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}
