resource "dws_gpu" "example" {
  gpu_name = "RTX 3090"
  image = "ubuntu:latest"
  region = "Europe"
  ssh_key = "ssh-rsa ..."
  gpu_count = 1
}
