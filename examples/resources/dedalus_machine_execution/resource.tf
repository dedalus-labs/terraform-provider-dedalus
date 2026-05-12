resource "dedalus_machine_execution" "example_machine_execution" {
  machine_id = "dm-3"
  command = ["string"]
  cwd = "cwd"
  env = {
    foo = "string"
  }
  stdin = "stdin"
  timeout_ms = 0
}
