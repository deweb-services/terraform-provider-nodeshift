package client

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestNodeshiftClient_CreateBucket(t *testing.T) {
	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		ctx    context.Context
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
				ctx:    context.TODO(),
				bucket: &S3BucketConfig{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			got, err := c.CreateBucket(tt.args.ctx, tt.args.bucket)
			assert.Nil(t, got)
			assert.NotNil(t, err)
		})
	}
}

func TestNodeshiftClient_DeleteBucket(t *testing.T) {
	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		ctx context.Context
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
				ctx: context.TODO(),
				key: "key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			err := c.DeleteBucket(tt.args.ctx, tt.args.key)
			assert.NotNil(t, err)
		})
	}
}

func TestNodeshiftClient_GetBucket(t *testing.T) {
	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		ctx context.Context
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
				ctx: context.TODO(),
				key: "key",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			got, err := c.GetBucket(tt.args.ctx, tt.args.key)
			assert.Nil(t, got)
			assert.NotNil(t, err)
		})
	}
}

func TestNodeshiftClient_UpdateBucket(t *testing.T) {
	type fields struct {
		Config   NodeshiftProviderConfiguration
		s3client *s3.Client
	}
	type args struct {
		ctx    context.Context
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
				ctx:    context.TODO(),
				bucket: &S3BucketConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NodeshiftClient{
				Config:   tt.fields.Config,
				s3client: tt.fields.s3client,
			}
			err := c.UpdateBucket(tt.args.ctx, tt.args.bucket)
			assert.NotNil(t, err)
		})
	}
}
