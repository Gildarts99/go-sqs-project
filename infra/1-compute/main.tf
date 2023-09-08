variable "lambda-function-name" {
  type = string
  default = "go-sqs-lambda"
}

resource "aws_cloudwatch_log_group" "example" {
  name              = "/aws/lambda/${var.lambda-function-name}"
  retention_in_days = 14
}

resource "aws_lambda_function" "go-sqs-lambda-function" {
  image_uri      = "${data.terraform_remote_state.core.outputs.repository_url}:latest"
  function_name = var.lambda-function-name
  package_type = "Image"
  role = aws_iam_role.go-sqs-lambda-cloudwatch-role.arn
}
