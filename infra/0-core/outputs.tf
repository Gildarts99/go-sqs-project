output "sqs_url" {
  value = aws_sqs_queue.go_sqs_queue.url
}

output "sqs_arn" {
  value = aws_sqs_queue.go_sqs_queue.arn
}

output "s3_bucket_name" {
  value = aws_s3_bucket.go-sqs-bucket.bucket
}

output "s3_bucket_arn" {
  value = aws_s3_bucket.go-sqs-bucket.arn
}

output "repository_url" {
  value = aws_ecr_repository.go-sqs-queue.repository_url
}

output "role_arn" {
  value = aws_iam_role.github-actions-role-go-sqs.arn
}
