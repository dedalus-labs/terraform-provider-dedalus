# Provider generation

`openapi.json` pins the Machines API at source revision
`54f38198424c36422da53af5d9148f38c3df6bde`, with SHA-256
`0e63748da0a11fbaf28e09ea90e2896113cbb32fe02398029f7512e95f2d7ee1`.
`config.yml` selects the Terraform resource surface and pins Go SDK v0.8.0.
The provider supplies JSON request bodies directly, including password inputs
introduced after that SDK release.

Verified with Dedalus SDK Generator revision `52e673c` and Terraform plugin
SHA-256 `67ba7aaef89309f295306f8420de7ee2af2b72bd8702c9501aa11b02bda98864`.

Generate a candidate with a local Dedalus SDK Generator installation:

```sh
./scripts/generate /path/to/sdkgen/dist/index.cjs
```

The command writes into a fresh `.tmp/generated.*` directory, runs the generated
build, lint and tests, then applies `custom.patch` and runs the tests again.
It does not overwrite the repository or publish anything. Compare the generated
candidate with this repository before integrating it. Keep release workflows,
contribution instructions, lifecycle tests, and this generation directory.

The custom patch reconstructs an imported machine's autosleep duration from the
actual seconds returned by the API and encodes the port update contract: required
visibility and an optional non-null password. OpenAPI transforms mark write-only
port passwords sensitive. The private port example omits credentials and expiry.
The snapshot itself remains unchanged. Terraform still stores sensitive values
in state.

Before release, run the real Terraform CLI contract tests:

```sh
TF_ACC=1 ./scripts/test
```

To prove an upgrade, additionally set `TF_ACC_PREVIOUS_PROVIDER` to a built
v0.3.0 provider binary. Live staging verification is separate: see
`scripts/test-live` and `CONTRIBUTING.md`.
