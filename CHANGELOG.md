# Changelog

## 0.1.1 (2026-04-30)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/dedalus-labs/terraform-provider-dedalus/compare/v0.1.0...v0.1.1)

### Chores

* add local tmpfile directory ([f3799af](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/f3799afb773c443b13d1df132808825017bb959b))
* **internal:** codegen related update ([74e669b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/74e669b30ecfadc8fecbd5494f7c567535741d48))

## 0.1.0 (2026-04-22)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/dedalus-labs/terraform-provider-dedalus/compare/v0.0.2...v0.1.0)

### Features

* add per-resource api permissions to schema description ([3efa1a0](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/3efa1a0f82b922b114fd2377bd451ec44f07abb4))


### Bug Fixes

* **ci:** in custom setup-go, pass through go-version and cache-dependency-path ([1a0affd](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1a0affd236016a62b661646b2359e404572ee458))
* fall back to main branch if linking fails in CI ([de7a902](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/de7a9021701819d83d0a34dc3fac954825e19737))
* fix for failing to drop invalid module replace in link script ([3d5986b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/3d5986bec0e6033eaa6b37d93e22253d17c78934))
* fix quoting typo ([24d23fa](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/24d23fa9ca7624f9251a775f564185aa3143f524))
* **tests:** update hc-install to fix PGP key mismatch ([3158857](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/31588571ea57000c60f55f84f95774d14dea68d1))


### Chores

* **ci:** remove release-doctor workflow ([95af827](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/95af827ad4e9561ae908cf722ff5af167372ec04))
* **internal:** more robust bootstrap script ([c860d46](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c860d46c2a17643cc38df5c290f200d98feb5944))
* pin go releaser version ([22897ca](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/22897ca94132b9b8d7e6290b1e0332b526469af5))
* **tests:** bump steady to v0.22.1 ([097c14c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/097c14cc83a8dbf1bb8013db7dfb8a1f6c553f54))

## 0.0.2 (2026-04-02)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/dedalus-labs/terraform-provider-dedalus/compare/v0.0.1...v0.0.2)

### Bug Fixes

* **internal:** emit boolean flag in datasource timeout schema opts ([75cd237](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/75cd237e328eeae653d164193555f1686972a8bf))
* **schema:** remove invalid Read field from datasource timeouts.Opts ([#4](https://github.com/dedalus-labs/terraform-provider-dedalus/issues/4)) ([c37ef1c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c37ef1c2ec184c8a29685a92dbfbe1330b278153))


### Chores

* **api:** refresh codegen ([2124763](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/21247631bb16586135aada4e2303e4fefd336c46))
* **api:** rename workspaces to machines ([04b0561](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/04b056180f23bce4555cc494259658da447e9115))
* **internal:** update multipart form array serialization ([382ef51](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/382ef5112773cc4695d7b7f1ff28e2bb8b7eae66))
* remove custom code ([f7798cc](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/f7798ccf8fb58aede4c59726942c9f2deff3b03c))
* **tests:** bump steady to v0.20.1 ([6e7094f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/6e7094f3faddd3afe41da73f3abb1c7eadb2cdf8))
* **tests:** bump steady to v0.20.2 ([08dfc04](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/08dfc04c4dcef9070e29201369954b0b274de5cf))

## 0.0.1 (2026-03-25)

Full Changelog: [v0.0.1...v0.0.1](https://github.com/dedalus-labs/terraform-provider-dedalus/compare/v0.0.1...v0.0.1)

### Features

* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([60ec21c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/60ec21c6aedb73cd2e0f7125e530c106c73d38ae))
* **api:** config update for dedalus-ai/dev ([2b0632f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/2b0632f2d13fef6598b637403e7b61b1c7dd61b3))
* **api:** Merge pull request [#2](https://github.com/dedalus-labs/terraform-provider-dedalus/issues/2) from stainless-sdks/dedalus-ai/dev ([088e9ca](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/088e9cadb3880d5afadb1a9456fbb36e901f667f))
* **api:** stable beta ([83b0da4](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/83b0da43176e0882a2b25aeb65221abcda227719))


### Bug Fixes

* **api:** add workspace resource/data sources, remove pet/store/user, add auth parameters ([00f5c69](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/00f5c6914bc9a922ff7ce16a16a1b2138a65f68c))
* **api:** update flags ([37a76c1](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/37a76c17cabc02a76d55b2ef6030d47490f5c253))
* bugfix for setting JSON keys with special characters ([143a365](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/143a3657beb7340bc7620e665ab3baa657640f4b))
* **client:** correctly encode map patches ([da1dd39](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/da1dd39cb18fb231fce395b4941e968a8a4aeae2))
* **client:** correctly patch `null` -&gt; zero value ([9b81d8f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/9b81d8ffed0057e8d9e2067e0051966f9d80d905))
* configurability of some attributes ([69b9802](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/69b98029d3b52d6f6b3e58b10e2f20e81c649228))
* correctly detect more ID attributes for data sources ([d62ff02](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d62ff02973f7eaf8a299bd9f95bc18e2e342cf6f))
* ensure dynamic values always yield valid container inner values ([aee1bba](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/aee1bba37af3b3c8440e61b5f5595e0300e1cd9d))
* improve linking behavior when developing on a branch not in the Go SDK ([b2394d3](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/b2394d3b751898ead997183422683150c72eb104))
* improved workflow for developing on branches ([dce7dc7](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/dce7dc740075e1f821d4848c98501ac24f90d33e))
* list style data sources should always have id value populated ([1ed0f29](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1ed0f29db126fa7391c97008385fd9e5fe9a914e))
* no longer require an API key when building on production repos ([8e2b1d4](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/8e2b1d4a6872dfd6223df26688e82b674d249745))
* patch style requests should never send empty json body for objects ([1c0c684](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1c0c6840003e4b6e0e293403f91d88d747bf7b6e))
* prefer named identifier field over id alias, missing ImportStates in certain resources ([5dd2304](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/5dd2304a394b8982b71a66ecf8c8090c7bdb0bca))
* properly convert properties into parameter types ([e569376](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/e569376fae2ac338ad09f50401809bc8659b55e3))
* spurious update plans for float attributes after import ([eff2221](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/eff2221735a682d3ee2e89efb9f5dffaff7aa3fa))
* **tests:** handle opaque timeouts.Value in schema parity checker ([5318c2d](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/5318c2d3fc64b0fb27dedc5e606ba8cbf6247a9b))
* **tests:** skip timeout parity tests until Stainless fixes codegen ([8df9e7b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/8df9e7b4f762f18aec2210ae3257172ab6ed19ae))


### Chores

* **api:** resolving merge conflicts ([088e9ca](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/088e9cadb3880d5afadb1a9456fbb36e901f667f))
* **api:** resolving merge conflicts ([ef93913](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/ef939139e22ca57095729b0281b8dc6fa24d0244))
* **api:** update homebrew tap and code samples ([60c0ba9](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/60c0ba939c909c4af465518d05ca8a3fef8de843))
* bump dependency version ([677fe2b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/677fe2bddc6b5d8501fc9db4b9ee3c76c34ad3a9))
* configure new SDK language ([6a9c6f1](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/6a9c6f188f3aff8cf2a25ef36046e24f87ce1334))
* configure new SDK language ([dc754bc](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/dc754bc7e86d5db0a962f9d165a295fa031bd6e8))
* **docs:** add missing descriptions ([d56fce5](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d56fce5e3a56914eb6da54fe2dc33f37d730f155))
* **docs:** update terraform-plugin-docs to v0.24.0 ([74e392d](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/74e392dcf19c0d937e5f9085a4462ce040fbeab3))
* ensure tests build as part of lint step ([448cdd6](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/448cdd6848860fdfe95268fb7286b953ce4f48af))
* **internal:** address linter warnings ([10eb6e0](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/10eb6e0e49c658eb650de58cb4d88012a03b7ca1))
* **internal:** codegen related update ([c6b5e64](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c6b5e64a58bf5a6541e3cef7b5a02e1d47b3b143))
* **internal:** codegen related update ([693cc77](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/693cc7784015d00d27ddbbeafd7af71205a5db15))
* **internal:** codegen related update ([e0cfd1f](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/e0cfd1f1c68fca12a5ab334a6a361cc2099a2760))
* **internal:** codegen related update ([c9d3932](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c9d393238b839fca9335ffea654ec826e1f78357))
* **internal:** codegen related update ([bda57ea](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/bda57ea57740d1a62c55731676872ab3148d1c25))
* **internal:** codegen related update ([49c4f74](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/49c4f7450439fc69c30bb4c7a9cf4bd117cfc26b))
* **internal:** codegen related update ([866b034](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/866b034c8ecb3489b2016beb51b052a359d669e6))
* **internal:** codegen related update ([0be2fba](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/0be2fbac7292e85b66e32116b97b890e7f0b84f1))
* **internal:** codegen related update ([7186cad](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/7186cade5df50fa900019162cc59e545d87be761))
* **internal:** refactor the apijson encoder ([798c1a1](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/798c1a182ff896bd4e63b0a213012cb63d2ef802))
* **internal:** remove mock server code ([1888fd8](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1888fd8d7714d7ca9080aa80083deba94c0f8def))
* **internal:** tweak CI branches ([1cdc5ce](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/1cdc5ce7be8bb51059b3a895713053dd49b7ef48))
* **internal:** update `actions/checkout` version ([e3bd3ad](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/e3bd3ad0720953695b9312c5572228a9932a39e2))
* **internal:** update `interface{}` to `any` ([43f6696](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/43f66968ff649cfb995fa61e481baf1396456f01))
* **internal:** update gitignore ([d1d2710](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d1d2710ab1672b97d052763155a3216bd09b9a84))
* **tests:** bump steady to v0.19.4 ([6049152](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/6049152ee85d5ef14fa8bb04a52db10868fbc870))
* **tests:** bump steady to v0.19.5 ([85a3ec6](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/85a3ec620035c2facffd5fa6a4d9d1e2394dcf1e))
* **tests:** bump steady to v0.19.6 ([d285cd5](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/d285cd55727c4045e4e2ac6752a8d13023873448))
* **tests:** bump steady to v0.19.7 ([c24825c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/c24825cd5b579e1caec53004b070205a1b45de26))
* update Go SDK version ([9a0e07c](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/9a0e07c6a7565e622ae4f148d961a1f7f52f9cb9))
* update SDK settings ([5f56f80](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/5f56f80d3b7e58df1e05b57f21faaddc4cf6082d))
* update SDK settings ([2704de6](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/2704de639853ec64c6235c8ade87b3f3d0a50966))
* update SDK settings ([8a58ec9](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/8a58ec90a15a51d6cfdb065dd91f9ebf535e030f))


### Refactors

* **tests:** switch from prism to steady ([9e3134b](https://github.com/dedalus-labs/terraform-provider-dedalus/commit/9e3134b6f5e8ba85fa3930a8c8e31d2cfd7ce5fb))
