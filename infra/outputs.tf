output "sqs_url" {
  value = aws_sqs_queue.go_sqs_queue.id
}
