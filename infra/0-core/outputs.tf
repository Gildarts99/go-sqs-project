output "sqs_url" {
  value = aws_sqs_queue.go_sqs_queue.id
}

output "repository_url" {
  value = aws_ecr_repository.go-sqs-queue.repository_url
}

output "role_arn" {
  value = aws_iam_role.github-actions-role-go-sqs.arn
}
