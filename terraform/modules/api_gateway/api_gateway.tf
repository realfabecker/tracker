resource "aws_api_gateway_rest_api" "rest_api" {
  name = "sintese"
}

resource "aws_api_gateway_deployment" "deployment" {
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
  description = "v1"
  depends_on  = [
    aws_api_gateway_integration.wallet_lambda
  ]
}

resource "aws_api_gateway_stage" "api_stage" {
  deployment_id = aws_api_gateway_deployment.deployment.id
  rest_api_id   = aws_api_gateway_rest_api.rest_api.id
  stage_name    = "v1"
}