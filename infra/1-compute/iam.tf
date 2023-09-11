resource "aws_iam_policy" "go-sqs-lambda" {
  name        = "go-sqs-lambda-policy"
  description = "perrmisions for lambda function to post metrics to cloudwatch"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
        ]
        Resource = "arn:aws:logs:*:*:*"
      },
      {
        Effect   = "Allow"
        Action = [
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes",
          "sqs:ReceiveMessage"
        ]
        Resource = data.terraform_remote_state.core.outputs.sqs_arn
      },
      {
        Effect   = "Allow"
        Action = [
          "s3:PutObject",
        ]
        Resource = [
          "${data.terraform_remote_state.core.outputs.s3_bucket_arn}",
          "${data.terraform_remote_state.core.outputs.s3_bucket_arn}/*"
        ]
      }
    ]
  })
}

resource "aws_iam_role" "go-sqs-lambda-role" {
  name               = "go-sqs-lambda-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action = "sts:AssumeRole"
        Principal = {
          "Service" = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "go-lambda-policy-attachment" {
  role       = aws_iam_role.go-sqs-lambda-role.name
  policy_arn = aws_iam_policy.go-sqs-lambda.arn
}
