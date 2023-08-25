# setup an sqs queue
resource "aws_sqs_queue" "go_sqs_queue" {
  name                      = "go-sqs-queue"
  message_retention_seconds = 600

  tags = {
    project = "GO-SQS-PROJECT"
  }
}