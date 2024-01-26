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

// OR

provider "dws" {
  shared_credentials_file = "~/.dws/credentials"
  profile = "main-profile"
}

// OR in case you want to set params with environment variables

provider "dws" {}

resource "dws_vpc" "example" {
  name = "example"
  description = "just an example vpc"
}

resource "dws_deployment" "hello_world" {
  image = "Ubuntu-v22.04"
  region = "USA"
  cpu = 4
  // RAM in MB
  ram = 2
  // Disk in MB
  disk_size = 20
  disk_type = "hdd"
  assign_public_ipv4 = true
  assign_public_ipv6 = true
  assign_ygg_ip = true
  ssh-key = "ssh-rsa ..."
  network_id = dws_network.example.id
}
