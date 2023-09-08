data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "go-sqs-lambda-role" {
  name               = "go-sqs-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_ecr_repository" "public" {
  name = "go-sqs-lambda"
}

output "image" {
  value = data.aws_ecr_repository.public.repository_url
}
# resource "aws_lambda_function" "go-sqs-lambda-function" {
#   image_uri      = "gildarts99/go-sqs-lambda:latest"
#   function_name = "go-sqs-lambda"
#   package_type = "Image"
#   role = aws_iam_role.go-sqs-lambda-role.arn
# }

