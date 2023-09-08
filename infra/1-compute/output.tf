output "image-uri" {
  value = data.terraform_remote_state.core.outputs.repository_url
}
