resource "nodeshift_load_balancer" "example" {
  name = "loadbalancer-1"
  replicas = {
    replica1 = 1
  }
  cpu_uuids = [
    "cpu_uuid_1",
    "cpu_uuid_2"
  ]
  forwarding_rules = [
    {
      in = {
        protocol = "HTTP"
        port     = 80
      }
      out = {
        protocol = "HTTPS"
        port     = 443
      }
    }
  ]
  vpc_uuid = "vpc_uuid"
}