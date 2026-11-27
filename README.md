# Dedalus Terraform Provider

The [Dedalus Terraform provider](https://registry.terraform.io/providers/dedalus-labs/dedalus/latest/docs) provides convenient access to
the [Dedalus REST API](https://docs.dedaluslabs.ai) from Terraform.

It is generated with Dedalus SDK Generator from the pinned API definition in [generation](generation/README.md).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    dedalus = {
      source  = "dedalus-labs/dedalus"
      version = "~> 0.4.0"
    }
  }
}

# Initialize the provider
provider "dedalus" {
  # Dedalus API key sent as Authorization Bearer.
  api_key = "My API Key" # or set DEDALUS_API_KEY env variable
  # Dedalus API key sent as x-api-key header.
  x_api_key = "My X API Key" # or set DEDALUS_X_API_KEY env variable
  # Organization ID header for all DCS requests.
  dedalus_org_id = "My Dedalus Org ID" # or set DEDALUS_ORG_ID env variable
}

# Configure a resource
resource "dedalus_machine" "example_machine" {
  autosleep = "300s"
  memory_mib = 2048
  storage_gib = 10
  vcpu = 1
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/dedalus-labs/dedalus/latest/docs).

## Upgrading to 0.4.0

Machine ports are now available as a resource and as singular/list data sources.
Port visibility changes update the existing binding. API credentials and port
sharing passwords are marked sensitive in Terraform output; Terraform state still
contains those values and must be protected.

CPU and memory are creation-only properties in this provider: changing an
explicitly configured `vcpu` or `memory_mib` now plans machine replacement.
Review that plan before applying because replacement deletes the old machine and
its storage. External resizing can produce the same replacement plan when the
configured shape differs from the API. If autoresizing owns these values, use
`lifecycle.ignore_changes` for `vcpu` and `memory_mib`. For an in-place manual
resize, use the Machines resize API.

Import reads `autosleep_seconds` from the API and records `never` for zero or
`<seconds>s` otherwise. Use that canonical spelling for a no-change plan after
import (`60s`, for example, rather than `1m`). Normal refresh preserves an
existing configured spelling.

Key-required ports use an owner-chosen `password` of exactly eight ASCII letters
or digits. The API never returns it, so normal refresh preserves the configured
value and import leaves it null. Reapply the password after importing a protected
port. Password updates include visibility in the same request.

## Upgrading to 0.3.0

The provider follows the public Machines contract: machines, SSH sessions, and
executions. Preview and terminal resources and data sources are removed. Migrate
or destroy any managed `dedalus_machine_preview` and `dedalus_machine_terminal`
resources with the previous provider before upgrading. Remove their data-source
references from configuration as well.

Machine creation defaults to 1 vCPU, 4096 MiB of memory, 10 GiB of storage, and
`300s` autosleep. The `dedalus_machines` data source exposes `items[*].phase`
instead of `items[*].status.phase`.

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property       | Environment variable | Required | Default value |
| -------------- | -------------------- | -------- | ------------- |
| x_api_key      | `DEDALUS_X_API_KEY`  | false    | —             |
| dedalus_org_id | `DEDALUS_ORG_ID`     | false    | —             |
| api_key        | `DEDALUS_API_KEY`    | false    | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/dedalus-labs/terraform-provider-dedalus/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
