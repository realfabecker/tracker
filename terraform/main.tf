module "database_dynamodb" {
  source = "./modules/dynamodb"
}

module "backend_lambda" {
  source = "./modules/lambda"

  function_name = "wallet"
  entrypoint    = "bootstrap"

  zip_source_file = "../backend/out/bootstrap"
  zip_output_path = "../backend/out/"

  zip_file_name = "lambda_backend_handler.zip"

  dynamodb_table_arn = module.database_dynamodb.dynamodb_table_arn

  COGNITO_JWK_URL     = var.COGNITO_JWK_URL
  COGNITO_CLIENT_ID   = var.COGNITO_CLIENT_ID
  APP_NAME            = var.APP_NAME
  DYNAMODB_TABLE_NAME = var.DYNAMODB_TABLE_NAME
}
