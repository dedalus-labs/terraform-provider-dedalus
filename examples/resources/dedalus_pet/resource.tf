resource "dedalus_pet" "example_pet" {
  name = "doggie"
  photo_urls = ["string"]
  id = 10
  category = {
    id = 1
    name = "Dogs"
  }
  status = "available"
  tags = [{
    id = 0
    name = "name"
  }]
}
