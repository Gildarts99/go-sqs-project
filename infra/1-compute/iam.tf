resource "aws_iam_policy" "go-sqs-lambda-cloudwatch" {
  name        = "go-sqs-lambda-cloudwatch-policy"
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
      }
    ]
  })
}

resource "aws_iam_role" "go-sqs-lambda-cloudwatch-role" {
  name               = "go-sqs-lambda-cloudwatch-role"
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

resource "aws_iam_role_policy_attachment" "go-lambda-cloudwatch-policy-attachment" {
  role       = aws_iam_role.go-sqs-lambda-cloudwatch-role.name
  policy_arn = aws_iam_policy.go-sqs-lambda-cloudwatch.arn
}
