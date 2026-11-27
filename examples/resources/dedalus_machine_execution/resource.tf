resource "dedalus_machine_execution" "example_machine_execution" {
  machine_id = "017f22e2-79b0-7cc3-98c4-dc0c0c07398f"
  command = ["string"]
  cwd = "cwd"
  env = {
    foo = "string"
  }
  stdin = "stdin"
  timeout_ms = 0
}
