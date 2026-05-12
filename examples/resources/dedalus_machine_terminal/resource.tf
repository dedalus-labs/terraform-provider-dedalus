resource "dedalus_machine_terminal" "example_machine_terminal" {
  machine_id = "dm-3"
  height = 0
  width = 0
  cwd = "cwd"
  env = {
    foo = "string"
  }
  shell = "shell"
}
