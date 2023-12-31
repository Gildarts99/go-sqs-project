variable "lambda-function-name" {
  type = string
  default = "go-sqs-lambda"
}

resource "aws_ssm_parameter" "encryption_key" {
  name  = "/go-sqs-project/ENCRYPTION_KEY"
  type  = "SecureString"
  # must be a 16, 24, or 32 bit string
  value = "supersecretkeyyy"
}

resource "aws_cloudwatch_log_group" "lambda-function" {
  name              = "/aws/lambda/${var.lambda-function-name}"
  retention_in_days = 14
}

resource "aws_secretsmanager_secret" "example" {
  name = "example"
}

resource "aws_lambda_function" "go-sqs-lambda-function" {
  function_name = var.lambda-function-name
  architectures = ["x86_64"]
  package_type = "Image"
  image_uri      = "${data.terraform_remote_state.core.outputs.repository_url}:latest"
  role = aws_iam_role.go-sqs-lambda-role.arn

  environment {
    variables = {
      S3_BUCKET = data.terraform_remote_state.core.outputs.s3_bucket_name,
      ENCRYPTION_KEY_PATH = aws_ssm_parameter.encryption_key.name
    }
  }
}

resource "aws_lambda_event_source_mapping" "sqs-lambda" {
  event_source_arn = data.terraform_remote_state.core.outputs.sqs_arn
  function_name    = aws_lambda_function.go-sqs-lambda-function.arn
}

