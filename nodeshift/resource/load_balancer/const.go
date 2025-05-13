package load_balancer

const UUID = "uuid"

const (
	KeyName            = "name"
	KeyReplicas        = "replicas"
	KeyCPUUUIDs        = "cpu_uuids"
	KeyForwardingRules = "forwarding_rules"
	KeyVPCUUID         = "vpc_uuid"
	KeyStatus          = "status"
	KeyTaskId          = "task_id"
)

const (
	DescriptionName            = "Name of the load balancer"
	DescriptionReplicas        = "Number of replicas of the load balancer"
	DescriptionCPUUUIDs        = "CPU UUIDs of the load balancer"
	DescriptionForwardingRules = "Forwarding Rules of the load balancer in/out"
	DescriptionVPCUUID         = "VPC UUID of the load balancer"

	DescriptionUUID   = "String UUID of the LB, computed"
	DescriptionStatus = "Status of the load balancer"
	DescriptionTaskId = "Task ID of the load balancer"
)
