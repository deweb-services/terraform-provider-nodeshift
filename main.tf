terraform {
  required_providers {
    dws = {
      source = "registry.terraform.io/dws/dws"
    }
  }
}

provider "dws" {
  access_key = "jwsaz4oednvfu6adykk7o4jqbqea"
  secret_access_key = "j2unsxboqvvovg5n3qklghsut3qpsgxpsbvw5rwmt7jo5iuuwyl3u"
}

resource "dws_vpc" "example" {
  name = "example"
  description = "just an example vpc"
  ip_range = "10.0.0.0/16" 
}

resource "dws_deployment" "hello_world" {
  image = "Ubuntu-v22.04"
  region = "USA"
  host_name = "test"
  cpu = 2
  ram = 2
  disk_size = 20
  disk_type = "hdd"
  assign_public_ipv4 = true
  assign_public_ipv6 = false
  assign_ygg_ip = false
  ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC/TXgmXJhqwRvJC7g69JTvwlCheV6SlOHvK2Rx/WLAmL1HKgLWHm2DSpqUFMuxWeE3gLbUJ43YPzKFwf2qiIwjVCIfm+DHBp0md+sg/5xX9ZiB1hfhROy3CMOilZN07Q3wuB4bwnq/0dAhDcuP3NOPiZWyRTtx2Ix1icrZmEt9gJGacfV/3bG8eFJ+SMEUapLiRn711rwIyFWHpfYkS4ir4YHsFz7AfLMQJ+2rLgxoM3Q5GjVBU8Rz97DyQ/DOD/P4cB/CkfpEQG792YNNROF5NylK57eMsZDNa2bT8/VG6GnpO8GARIw9qluByqg5HRrz4qmONrOdriNN5D9orNX8RLlLJRvEAc6bb7rRuPzksi1H1kH4vkFcuB27iJ+iye+0E56XrjL3CZcNur3Go1q1DdhVm+SBvGbZk7RVD9rCM9e2gBX+Vdi0aQkuywOEQpv8HldlXml32twmUaz5uANJD1q9coiy0Z5BJ/2QCeE6d88KLS59u77NxN4Udd1XEtU= aleksandruglovskiy@MacBook-Pro-Aleksandr.local"
  vpc_id = dws_vpc.example.id
}