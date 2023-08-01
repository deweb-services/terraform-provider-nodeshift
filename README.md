# Terraform Provider for DWS

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

To make this plugin work locally, after the installation rename the binary file to
`terraform-provider-dws` and make a file in your `$HOME` directory named `.terraformrc`
to override the terraform repository with contents like 
```
provider_installation {

    dev_overrides {
        "hashicorp.com/edu/dws" = "%PATH TO $GOBIN%"
    }

    direct {}
}
````
