resource "aws_lambda_function" "lambda" {

  function_name = var.function_name
  handler       = var.entrypoint
  runtime       = "provided.al2"

  architectures = ["arm64"]

  filename         = data.archive_file.zip.output_path
  source_code_hash = data.archive_file.zip.output_base64sha256

  publish = false
  role    = aws_iam_role.iam_for_lambda.arn

  environment {
    variables = {
      "APP_NAME"            = var.APP_NAME
      "COGNITO_CLIENT_ID"   = var.COGNITO_CLIENT_ID
      "COGNITO_JWK_URL"     = var.COGNITO_JWK_URL
      "DYNAMODB_TABLE_NAME" = var.DYNAMODB_TABLE_NAME
    }
  }
}

resource "aws_lambda_function_url" "lambda_url" {
  function_name      = aws_lambda_function.lambda.function_name
  authorization_type = "NONE"
}
