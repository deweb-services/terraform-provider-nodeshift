resource "dws_deployment" "example" {
  image = "Ubuntu-v22.04"
  region = "USA"
  cpu = 4
  // RAM in GB
  ram = 2
  // Disk in GB
  disk_size = 20
  disk_type = "hdd"
  assign_public_ipv4 = true
  assign_public_ipv6 = true
  assign_ygg_ip = true
  ssh-key = "ssh-rsa ..."
}
