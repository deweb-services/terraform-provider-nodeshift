package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.NewNodeshiftProvider,
		providerserver.ServeOpts{
			Address: "registry.terraform.io/nodeshift/nodeshift",
		}); err != nil {
		log.Printf("server error: %v", err)
	}
}
