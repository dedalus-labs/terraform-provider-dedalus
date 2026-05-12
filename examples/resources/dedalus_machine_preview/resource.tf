resource "dedalus_machine_preview" "example_machine_preview" {
  machine_id = "dm-3"
  port = 0
  protocol = "http"
  visibility = "public"
}
