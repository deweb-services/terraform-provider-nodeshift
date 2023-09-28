package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c *DWSClient) CreateBucket(ctx context.Context, bucket *S3BucketConfig) (*S3BucketConfig, error) {
	_, err := c.s3client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket.Key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create bucket: %w", err)
	}
	tflog.Info(ctx, fmt.Sprintf("created bucket: %s", bucket.Key))
	return &S3BucketConfig{Key: bucket.Key}, nil
}

func (c *DWSClient) GetBucket(ctx context.Context, key string) (*S3BucketConfig, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get bucket by key: %s", key))
	_, err := c.s3client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(key + " "),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}
	tflog.Info(ctx, fmt.Sprintf("head bucket: %s", key))
	return &S3BucketConfig{Key: key}, nil
}

func (c *DWSClient) UpdateBucket(ctx context.Context, bucket *S3BucketConfig) error {
	return errors.New("not implemented")
}

func (c *DWSClient) DeleteBucket(ctx context.Context, key string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete bucket by key: %s", key))
	_, err := c.s3client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete bucket: %w", err)
	}
	tflog.Info(ctx, fmt.Sprintf("delete bucket: %s", key))
	return nil
}
