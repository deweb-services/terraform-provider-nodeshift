package main

import (
	"context"
	"log"

	"github.com/deweb-services/terraform-provider-dws/dws/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.NewDWSProvider,
		providerserver.ServeOpts{
			Address: "registry.terraform.io/dws/dws",
		}); err != nil {
		log.Printf("server error: %v", err)
	}
}
