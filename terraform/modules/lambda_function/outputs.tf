output "alias_function_name" {
  value = data.aws_lambda_alias.current.function_name
}

output "alias_name" {
  value = data.aws_lambda_alias.current.name
}

output "alias_arn" {
  value = data.aws_lambda_alias.current.arn
}
