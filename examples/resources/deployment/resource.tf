resource "dws_deployment" "example" {
  image = "Ubuntu-v22.04"
  region = "United States"
  cpu = 1
  // RAM in MB
  ram = 1024
  // Disk in MB
  disk_size = 61440
  disk_type = "hdd"
  assign_public_ipv4 = true
  assign_public_ipv6 = false
  assign_ygg_ip = true
  ssh_key = "ssh-ed25519"
  ssh_key_name = "very-unique-name"
  host_name = "bestname"
}
