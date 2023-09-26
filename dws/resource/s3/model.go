package s3

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BucketResourceModel struct {
	BucketName types.String `tfsdk:"bucket_name"`
}
