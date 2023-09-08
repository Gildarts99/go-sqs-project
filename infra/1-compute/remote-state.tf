data "terraform_remote_state" "core" {
  backend = "local"

  config = {
    path = "../0-core/terraform.tfstate"
  }
}