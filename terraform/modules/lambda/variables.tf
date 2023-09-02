variable "APP_NAME" {
  type = string
}

variable "DYNAMODB_TABLE_NAME" {
  type = string
}

variable "COGNITO_CLIENT_ID" {
  type = string
}

variable "COGNITO_JWK_URL" {
  type = string
}

variable "function_name" {
  type = string
}

variable "entrypoint" {
  type = string
}

variable "zip_file_name" {
  type = string
}

variable "zip_source_file" {
  type = string
}

variable "zip_output_path" {
  type = string
}

variable "dynamodb_table_arn" {
  type = string
}
