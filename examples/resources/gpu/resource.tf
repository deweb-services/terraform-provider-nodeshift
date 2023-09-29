resource "dws_gpu" "example" {
  gpu_name = "RTX_A4000"
  image = "ubuntu:latest"
  region = "Europe"
  ssh-key = "ssh-rsa ..."
  gpu_count = 1
}
