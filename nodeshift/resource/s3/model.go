package s3

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type BucketResourceModel struct {
	Key types.String `tfsdk:"bucket_name"`
}

func (m *BucketResourceModel) ToClientRequest() (*client.S3BucketConfig, error) {
	if m.Key.IsUnknown() || m.Key.IsNull() {
		return nil, fmt.Errorf("bucket name is required: %w", client.ErrPropertyEmpty)
	}
	bucket := client.S3BucketConfig{
		Key: m.Key.ValueString(),
	}

	return &bucket, nil
}

func (m *BucketResourceModel) FromClientResponse(c *client.S3BucketConfig) error {
	m.Key = types.StringValue(c.Key)

	return nil
}
