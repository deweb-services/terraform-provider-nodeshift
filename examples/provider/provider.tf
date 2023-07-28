terraform {
  required_providers {
    dws = {
      source = "hashicorp.com/edu/dws"
    }
  }
}

provider "dws" {
  access_key = "ACCESS_KEY"
  secret_access_key = "SECRET_ACCESS_KEY"
}