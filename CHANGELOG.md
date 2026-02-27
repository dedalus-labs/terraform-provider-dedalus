# Changelog

## 0.1.0 (2026-02-27)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/dedalus-labs/terraform-provider-dedalus/compare/v0.0.1...v0.1.0)

### Features

* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([60ec21c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/60ec21c6aedb73cd2e0f7125e530c106c73d38ae))
* **api:** config update for dedalus-ai/dev ([2b0632f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/2b0632f2d13fef6598b637403e7b61b1c7dd61b3))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([143a365](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/143a3657beb7340bc7620e665ab3baa657640f4b))
* **client:** correctly encode map patches ([da1dd39](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/da1dd39cb18fb231fce395b4941e968a8a4aeae2))
* **client:** correctly patch `null` -&gt; zero value ([9b81d8f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/9b81d8ffed0057e8d9e2067e0051966f9d80d905))
* configurability of some attributes ([69b9802](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/69b98029d3b52d6f6b3e58b10e2f20e81c649228))
* correctly detect more ID attributes for data sources ([d62ff02](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d62ff02973f7eaf8a299bd9f95bc18e2e342cf6f))
* ensure dynamic values always yield valid container inner values ([aee1bba](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/aee1bba37af3b3c8440e61b5f5595e0300e1cd9d))
* list style data sources should always have id value populated ([1ed0f29](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1ed0f29db126fa7391c97008385fd9e5fe9a914e))
* prefer named identifier field over id alias, missing ImportStates in certain resources ([5dd2304](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/5dd2304a394b8982b71a66ecf8c8090c7bdb0bca))
* properly convert properties into parameter types ([e569376](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/e569376fae2ac338ad09f50401809bc8659b55e3))
* spurious update plans for float attributes after import ([eff2221](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/eff2221735a682d3ee2e89efb9f5dffaff7aa3fa))


### Chores

* bump dependency version ([677fe2b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/677fe2bddc6b5d8501fc9db4b9ee3c76c34ad3a9))
* configure new SDK language ([6a9c6f1](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/6a9c6f188f3aff8cf2a25ef36046e24f87ce1334))
* configure new SDK language ([dc754bc](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/dc754bc7e86d5db0a962f9d165a295fa031bd6e8))
* **docs:** add missing descriptions ([d56fce5](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d56fce5e3a56914eb6da54fe2dc33f37d730f155))
* ensure tests build as part of lint step ([448cdd6](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/448cdd6848860fdfe95268fb7286b953ce4f48af))
* **internal:** address linter warnings ([10eb6e0](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/10eb6e0e49c658eb650de58cb4d88012a03b7ca1))
* **internal:** codegen related update ([c9d3932](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c9d393238b839fca9335ffea654ec826e1f78357))
* **internal:** codegen related update ([bda57ea](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/bda57ea57740d1a62c55731676872ab3148d1c25))
* **internal:** codegen related update ([49c4f74](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/49c4f7450439fc69c30bb4c7a9cf4bd117cfc26b))
* **internal:** codegen related update ([866b034](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/866b034c8ecb3489b2016beb51b052a359d669e6))
* **internal:** codegen related update ([0be2fba](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/0be2fbac7292e85b66e32116b97b890e7f0b84f1))
* **internal:** codegen related update ([7186cad](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/7186cade5df50fa900019162cc59e545d87be761))
* **internal:** refactor the apijson encoder ([798c1a1](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/798c1a182ff896bd4e63b0a213012cb63d2ef802))
* **internal:** remove mock server code ([1888fd8](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1888fd8d7714d7ca9080aa80083deba94c0f8def))
* **internal:** update `actions/checkout` version ([e3bd3ad](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/e3bd3ad0720953695b9312c5572228a9932a39e2))
* **internal:** update `interface{}` to `any` ([43f6696](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/43f66968ff649cfb995fa61e481baf1396456f01))
* update Go SDK version ([9a0e07c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/9a0e07c6a7565e622ae4f148d961a1f7f52f9cb9))
* update SDK settings ([5f56f80](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/5f56f80d3b7e58df1e05b57f21faaddc4cf6082d))
* update SDK settings ([2704de6](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/2704de639853ec64c6235c8ade87b3f3d0a50966))
* update SDK settings ([8a58ec9](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/8a58ec90a15a51d6cfdb065dd91f9ebf535e030f))
