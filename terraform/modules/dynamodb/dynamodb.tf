resource "aws_dynamodb_table" "sintese" {
  name = "sintese"
  
  billing_mode = "PAY_PER_REQUEST"

  hash_key  = "PK"
  range_key = "SK"

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }

  attribute {
    name = "GSI1_PK"
    type = "S"
  }

  attribute {
    name = "GSI1_SK"
    type = "S"
  }

  global_secondary_index {
    name            = "GSI_1"
    projection_type = "ALL"

    hash_key  = "GSI1_PK"
    range_key = "GSI1_SK"
  }
}