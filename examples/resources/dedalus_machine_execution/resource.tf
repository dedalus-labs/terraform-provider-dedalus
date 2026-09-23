resource "dedalus_machine_execution" "example_machine_execution" {
  machine_id = "dm-ecc2efdd-ddfa-31a9-c6f1-b833d337aa7c"
  command = ["string"]
  cwd = "cwd"
  env = {
    foo = "string"
  }
  stdin = "stdin"
  timeout_ms = 0
}
