terraform {
  required_providers {
    dws = {
      source = "hashicorp.com/edu/dws"
    }
  }
}

provider "dws" {
  account_name = "dws_acc"
  account_key = "dws_key"
  access_region = "us"
  api_key = "api"
  session_token = "tok"
}

resource "dws_vm" "hello_world" {
  deployment = {
    name = "string"
    image = "string"
    network = "string"
    region = "string"
  }
  cpu = [
    {
      quantity = 1
      type = "string"
    }
  ]
  ram = {
    volume = 1024
  }
  disk = [
    {
      type = "string"
      volume = 12
    }
  ]
  protocols = {
    ip = {
      v4 = true
      v6 = false
    }
  }
}
