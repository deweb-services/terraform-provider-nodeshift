resource "dws_gpu" "example" {
  gpu_name = "RTX A4000"
  image = "ubuntu:latest"
  region = "Central America"
  ssh-key = "ssh-rsa ..."
  gpu_count = 2
}
