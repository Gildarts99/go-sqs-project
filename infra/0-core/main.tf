# setup an sqs queue
resource "aws_sqs_queue" "go_sqs_queue" {
  name                      = "go-sqs-queue"
  message_retention_seconds = 600
}

# we want an ecr repository for our future lambda function to pull images from
resource "aws_ecr_repository" "go-sqs-queue" {
  name                 = "go-sqs-queue"
}

# create an s3 bucket for our lambda function to push too
resource "aws_s3_bucket" "example" {
  bucket = "go-sqs-project-bucket"

  tags = {
    Name        = "go-sqs-project-bucket"
  }
}
