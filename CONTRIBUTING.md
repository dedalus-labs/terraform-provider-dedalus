## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/build
```

This will install all the required dependencies and build the Provider binary into the root directory.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## Running the Provider locally

You can build the provider locally and have your `.tf` files refer to the local build instead of the one in the Hashicorp registry.

First, build the provider binary:

```sh
$ ./scripts/build
```

Then edit (or create) your `~/.terraformrc` to look something like this:

```hcl
  provider_installation {
    dev_overrides {
      "dedalus-labs/dedalus" = "/local/path/to/this/repo"
    }
    direct {}
  }
```

## Running tests

To execute the schema and unit tests, run:

```sh
$ ./scripts/test
```

The acceptance tests run the real Terraform CLI against a local HTTP contract
server. They create no cloud resources and use no saved credentials:

```sh
TF_ACC=1 ./scripts/test
```

The version-upgrade test additionally takes the path to a previously built
provider binary in `TF_ACC_PREVIOUS_PROVIDER`.

A separate live smoke test is available through `./scripts/test-live`; it requires
an explicit staging endpoint and API credential and creates disposable resources.
See the script's usage and cleanup behavior before running it.

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```

## Running Tests

To run schema tests,
