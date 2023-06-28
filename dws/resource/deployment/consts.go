package deployment

const (
	ID = "id"
)

// key names for vm resource
const (
	DeploymentKeysImage            = "image"
	DeploymentKeysRegion           = "region"
	DeploymentKeysCPU              = "cpu"
	DeploymentKeysRAM              = "ram"
	DeploymentKeysDiskSize         = "disk_size"
	DeploymentKeysDiskType         = "disk_type"
	DeploymentKeysAssignPublicIPv4 = "assign_public_ipv4"
	DeploymentKeysAssignPublicIPv6 = "assign_public_ipv6"
	DeploymentKeysAssignYggIP      = "assign_ygg_ip"
	DeploymentKeysSSHKey           = "ssh_key"
	DeploymentKeysHostName         = "host_name"
	DeploymentKeysVPCID            = "vpc_id"

	DeploymentKeysPublicIPv6 = "public_ipv6"
	DeploymentKeysPublicIPv4 = "public_ipv4"
	DeploymentKeysYggIP      = "ygg_ip"
)

const (
	ImageDescription = `OS Image used to install on the target Vitrual Machine Deployment. 
Available options: Ubuntu-v22.04`
	RegionDescription = `Region where you want to deploy your Deployment.
Available options: USA`
	CPUDescription              = `Number of CPU cores for your Deployment`
	RAMDescription              = `Amount of RAM for your Deployment in GB`
	DiskSizeDescription         = `Amount of disk size for your Deployment in GB`
	DiskTypeDescription         = `Disk type for your Deployment. Available options: hdd, ssd`
	AssignPublicIPv4Description = `If true assigns a public ipv4 address for your Deployment`
	AssignPublicIPv6Description = `If true assigns a public ipv6 address for your Deployment`
	AssignYggIPDescription      = `If true assigns a yggdrasil address for your Deployment`
	SSHKeyDescription           = `SSH key to add to the target VM to make it possible to connect to your VM`
	HostNameDescription         = `Host name for your Deployment`
	VPCIDDescription            = `ID of the vpc to deploy your VM into`
	PublicIPv4Description       = `Public IPv4 of your VM`
	PublicIPv6Description       = `Public IPv6 of your VM`
	YggIPDescription            = `Yggdrasil IP of your VM`
)
