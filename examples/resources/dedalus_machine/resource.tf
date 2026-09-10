# @custom start
resource "dedalus_machine" "example_machine" {
  autosleep   = "300s"
  memory_mib  = 4096
  storage_gib = 10
  vcpu        = 1
}
# @custom end
