resource "dedalus_machine" "example_machine" {
  autosleep = "autosleep"
  image_id = "image_id"
  memory_mib = 1
  storage_gib = 1
  vcpu = 1
}
