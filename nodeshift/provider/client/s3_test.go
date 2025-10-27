package client

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeshiftClient_CreateBucket(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		bucket *S3BucketConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *S3BucketConfig
	}{
		{
			name: "create_bucket",
			fields: fields{
				Config:   NodeshiftProviderConfiguration{},
				s3client: &s3.Client{},
			},
			args: args{
				bucket: &S3BucketConfig{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			got, err := c.CreateBucket(t.Context(), tt.args.bucket)
			require.Error(t, err)
			require.Nil(t, got)
		})
	}
}

func TestNodeshiftClient_DeleteBucket(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete_bucket",
			fields: fields{
				Config:   NodeshiftProviderConfiguration{},
				s3client: &s3.Client{},
			},
			args: args{
				key: "key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			err := c.DeleteBucket(t.Context(), tt.args.key)
			require.Error(t, err)
		})
	}
}

func TestNodeshiftClient_GetBucket(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *S3BucketConfig
	}{
		{
			name: "get_bucket",
			fields: fields{
				Config:   NodeshiftProviderConfiguration{},
				s3client: &s3.Client{},
			},
			args: args{
				key: "key",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			got, err := c.GetBucket(t.Context(), tt.args.key)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

func TestNodeshiftClient_UpdateBucket(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		bucket *S3BucketConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "update_bucket",
			fields: fields{
				Config:   NodeshiftProviderConfiguration{},
				s3client: &s3.Client{},
			},
			args: args{
				bucket: &S3BucketConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			err := c.UpdateBucket(t.Context(), tt.args.bucket)
			assert.Error(t, err)
		})
	}
}
