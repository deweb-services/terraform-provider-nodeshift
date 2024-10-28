terraform {
  backend "s3" {
    bucket     = "mybucket"
    region     = "us-west-1"
    key        = "mys3/state.tf"
    access_key = ""
    secret_key = ""
    endpoint   = "https://eu.storage.nodeshift.co"
    skip_credentials_validation = true
  }

  required_providers {
    nodeshift = {
      source = "hashicorp.com/edu/nodeshift"
    }
  }
}
