package deployment

const (
	UUID = "uuid"
)

// key names for vm resource.
const (
	DeploymentKeysImage            = "image"
	DeploymentKeysRegion           = "region"
	DeploymentKeysCPU              = "cpu"
	DeploymentKeysRAM              = "ram"
	DeploymentKeysDiskSize         = "disk_size"
	DeploymentKeysDiskType         = "disk_type"
	DeploymentKeysAssignPublicIPv4 = "assign_public_ipv4"
	DeploymentKeysAssignPublicIPv6 = "assign_public_ipv6"
	DeploymentKeysSSHKey           = "ssh_key"
	DeploymentKeysSSHKeyName       = "ssh_key_name"
	DeploymentKeysHostName         = "host_name"
	DeploymentKeysNetworkUUID      = "network_uuid"

	DeploymentKeysPublicIPv6 = "public_ipv6"
	DeploymentKeysPublicIPv4 = "public_ipv4"
)

const (
	ImageDescription = `OS Image used to install on the target Virtual Machine Deployment. 
Available options: Ubuntu-v22.04`
	RegionDescription = `Region where you want to deploy your Deployment.
Available options: USA`
	CPUDescription              = `Number of CPU cores for your Deployment`
	RAMDescription              = `Amount of RAM for your Deployment in MB`
	DiskSizeDescription         = `Amount of disk size for your Deployment in MB`
	DiskTypeDescription         = `Disk type for your Deployment. Available options: hdd, ssd`
	AssignPublicIPv4Description = `If true assigns a public ipv4 address for your Deployment`
	AssignPublicIPv6Description = `If true assigns a public ipv6 address for your Deployment`
	SSHKeyDescription           = `SSH key to add to the target VM to make it possible to connect to your VM`
	SSHKeyNameDescription       = `SSH key name for Deployment`
	HostNameDescription         = `Host name for your Deployment`
	NetworkUUIDDescription      = `UUID of the network to deploy your VM into`
	PublicIPv4Description       = `Public IPv4 of your VM`
	PublicIPv6Description       = `Public IPv6 of your VM`
)

// nolint: gochecknoglobals
var availableDiskTypes = []string{"hdd", "ssd"}
