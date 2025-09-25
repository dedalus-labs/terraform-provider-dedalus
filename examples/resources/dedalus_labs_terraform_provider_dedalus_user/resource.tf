resource "dedalus-labs-terraform-provider-dedalus_user" "example_user" {
  id = 10
  email = "john@email.com"
  first_name = "John"
  last_name = "James"
  password = "12345"
  phone = "12345"
  existing_username = "theUser"
  user_status = 1
}
