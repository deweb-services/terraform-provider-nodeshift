package gpu

const UUID = "uuid"

const (
	KeyGPUName  = "gpu_name"
	KeyImage    = "image"
	KeySSHKey   = "ssh_key"
	KeyGPUCount = "gpu_count"
	KeyRegion   = "region"
)

const (
	DescriptionUUID     = "String UUID of the GPU, computed"
	DescriptionGPUName  = "Name of the GPU to be used in deployment"
	DescriptionImage    = "Image used to run your docker container name:version"
	DescriptionSSHKey   = "SSH key to add to the target GPU to make it possible to connect to your deployment"
	DescriptionGPUCount = "Number of GPU instances for your deployment"
	DescriptionRegion   = `Region where you want to deploy your GPU. Available options: 
"Northern America", "Central America", "South America", "Europe", "Asia", "Africa", "Oceania", "Caribbean"`
)
