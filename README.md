# Terraform Provider for Nodeshift

## Build provider

Run the following command to build the provider

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory.

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

## Local debug setup

```shell
make docker-build # example output:  "path for local bin is <path>/terraform-provider-nodeshift/bin/"
nano $HOME/.terraformrc # with output path above and with config below
terraform init
NODESHIFT_TERRAFORM_API_URL=<desired http/https path> terraform apply
```

To make this plugin work locally, after the installation rename the binary file to
`terraform-provider-nodeshift` and make a file in your `$HOME` directory named `.terraformrc`
to override the terraform repository with contents like

.terraformrc file contents
```
provider_installation {

    dev_overrides {
        "deweb-services/nodeshift" = "%PATH TO $GOBIN%"
    }

    direct {}
}
````
