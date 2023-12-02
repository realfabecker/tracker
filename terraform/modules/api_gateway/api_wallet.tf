resource "aws_api_gateway_resource" "wallet" {
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
  parent_id   = aws_api_gateway_rest_api.rest_api.root_resource_id
  path_part   = "wallet"
}

resource "aws_api_gateway_resource" "wallet_proxy" {
  parent_id   = aws_api_gateway_resource.wallet.id
  path_part   = "{proxy+}"
  rest_api_id = aws_api_gateway_rest_api.rest_api.id
}

resource "aws_api_gateway_method" "wallet_proxy" {
  authorization = "NONE"
  http_method   = "ANY"
  resource_id   = aws_api_gateway_resource.wallet_proxy.id
  rest_api_id   = aws_api_gateway_rest_api.rest_api.id
  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_integration" "wallet_lambda" {
  http_method = aws_api_gateway_method.wallet_proxy.http_method
  resource_id = aws_api_gateway_method.wallet_proxy.resource_id
  rest_api_id = aws_api_gateway_rest_api.rest_api.id

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = var.wallet.invoke_arn
}

resource "aws_lambda_permission" "wallet_lambda" {
  statement_id  = "94c40a24-e8c7-529c-a50b-a4e9ed4a9aeb"
  action        = "lambda:InvokeFunction"
  function_name = var.wallet.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.rest_api.execution_arn}/*/*"
}