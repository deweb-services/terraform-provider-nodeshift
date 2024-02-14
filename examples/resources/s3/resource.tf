provider "nodeshift" {
  access_key = "access_key"
  secret_access_key = "secret_access_key"
  s3_region = "us-west-1"
  s3_endpoint = "https://s3.nodeshift.sh/"
}

resource "nodeshift_bucket" "example" {
  bucket_name = "absolutely_unique_name_19"
}

