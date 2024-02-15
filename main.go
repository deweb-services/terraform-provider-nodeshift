package main

import (
	"context"
	"log"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.NewNodeshiftProvider,
		providerserver.ServeOpts{
			Address: "registry.terraform.io/nodeshift/nodeshift",
		}); err != nil {
		log.Printf("server error: %v", err)
	}
}
