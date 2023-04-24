package main

import (
	"context"
	"log"

	"github.com/deweb-services/dws-terraform-provider/dws/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.NewDWSProvider,
		providerserver.ServeOpts{
			Address: "hashicorp.com/edu/dws",
		}); err != nil {
		log.Printf("server error: %v", err)
	}
}
