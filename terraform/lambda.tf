locals {
  lambda_config = {
    main = {
      role_arn = aws_iam_role.lambda.arn
    }
  }
}

module "lambda_function" {
  source = "./modules/lambda_function"

  function_name = local.prefix
  function_dir  = "${path.module}/../lambda"
  iam_role_arn  = aws_iam_role.lambda.arn
}
