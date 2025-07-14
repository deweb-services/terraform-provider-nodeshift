resource "nodeshift_gpu" "example" {
  gpu_name = "RTX 3090"
  image = "ubuntu:latest"
  region = "Europe"
  ssh_key = "ssh-rsa ..."
  gpu_count = 1
  disk_size_gb = 30
  min_cuda_version = "12.6"
}
