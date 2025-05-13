resource "nodeshift_load_balancer" "example" {
  name = "loadbalancer-1"
  replicas = {
    replica1 = 1
    replica2 = 2
  }
  cpu_uuids = [
    "fb7e26fd-7a0d-4ff8-9dc6-c7649b3edfb4",
    "46ccbae8-eeed-49f6-a5a2-a4ecaeb787ce"
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
  vpc_uuid = "0cd9b534-caef-4997-ac60-f74c458c7abe"
}