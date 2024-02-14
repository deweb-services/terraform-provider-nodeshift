terraform {
  required_providers {
    nodeshift = {
      source = "hashicorp.com/edu/nodeshift"
    }
  }
}

provider "nodeshift" {
  access_key = "ACCESS_KEY"
  secret_access_key = "SECRET_ACCESS_KEY"
}