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

resource "aws_lambda_function" "go-sqs-lambda-function" {
  image_uri      = "gildarts99/go-sqs-lambda:dd4f7f2deb46cea0814dc09a33fef2be1d62ce14"
  function_name = "go-sqs-lambda"
  package_type = "Image"
  role = aws_iam_role.go-sqs-lambda-role.arn
}
