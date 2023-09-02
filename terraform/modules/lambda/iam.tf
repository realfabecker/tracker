resource "aws_iam_role" "iam_for_lambda" {
  name = "${var.function_name}-role-34kqucx9"
  path = "/service-role/"

  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Principal" : {
          "Service" : "lambda.amazonaws.com"
        },
        "Effect" : "Allow",
        "Sid" : ""
      }
    ]
  })
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = jsonencode(
    {
      Statement = [
        {
          Action = [
            "logs:PutLogEvents",
            "logs:CreateLogStream",
            "logs:CreateLogGroup",
          ]
          Effect   = "Allow"
          Resource = "arn:aws:logs:*:*:*"
          Sid      = "CloudWatch"
        },
        {
          Action = [
            "dynamodb:GetItem",
            "dynamodb:Query",
            "dynamodb:DeleteItem",
            "dynamodb:PutItem",
          ]
          Effect   = "Allow"
          Resource = var.dynamodb_table_arn
          Sid      = "DynamoDB"
        },
        {
          Action = [
            "dynamodb:Query",
          ]
          Effect   = "Allow"
          Resource = "${var.dynamodb_table_arn}/index/GSI_1"
          Sid      = "DynamoDBIndexGSI1"
        },
      ]
      Version = "2012-10-17"
    }
  )
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}