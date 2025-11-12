package s3

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type BucketResourceModel struct {
	Key types.String `tfsdk:"bucket_name"`
}

func (m *BucketResourceModel) ToClientRequest() *client.S3BucketConfig {
	return &client.S3BucketConfig{
		Key: m.Key.ValueString(),
	}
}

func (m *BucketResourceModel) FromClientResponse(c *client.S3BucketConfig) {
	m.Key = types.StringValue(c.Key)
}
