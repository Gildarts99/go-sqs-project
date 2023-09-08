# setup github oidc
resource "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = [
    "6938fd4d98bab03faadb97b34396831e3780aea1",
    "1c58a3a8518e8759bf075b76b750d4f2df264fcd"
  ]
}

# setup the role and permissions that our github action will use
resource "aws_iam_policy" "go-sqs-lambda-ecr-policy" {
  name        = "go-sqs-lambda-ecr-policy"
  description = "perrmisions for go sqs github action"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action = [
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetRepositoryPolicy",
          "ecr:DescribeRepositories",
          "ecr:DescribeImages",
          "ecr:InitiateLayerUpload",
          "ecr:UploadLayerPart",
          "ecr:CompleteLayerUpload",
          "ecr:PutImage"
        ]
        Resource = "arn:aws:ecr:us-east-1:785210909375:repository/go-sqs-queue"
      },
      {
        Effect   = "Allow"
        Action = [
          "ecr:GetAuthorizationToken",
          "sts:GetServiceBearerToken"
        ]
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action = [
          "lambda:UpdateFunctionCode"
        ]
        Resource = "arn:aws:lambda:us-east-1:785210909375:function:go-sqs-lambda"
      }
    ]
  })
}

resource "aws_iam_role" "github-actions-role-go-sqs" {
  name               = "github-actions-role-go-sqs"
  path               = "/"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action = "sts:AssumeRoleWithWebIdentity"
        Principal = {
          "Federated" = "arn:aws:iam::785210909375:oidc-provider/token.actions.githubusercontent.com"
        },
        Condition = {
          StringEquals = {
            "token.actions.githubusercontent.com:sub" = "repo:Gildarts99/go-sqs-project:ref:refs/heads/main",
            "token.actions.githubusercontent.com:aud" = "sts.amazonaws.com"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "go-sqs-policy-attachment" {
  role       = aws_iam_role.github-actions-role-go-sqs.name
  policy_arn = aws_iam_policy.go-sqs-lambda-ecr-policy.arn
}
