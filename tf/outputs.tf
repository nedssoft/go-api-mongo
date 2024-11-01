output "domain_name" {
  value = "https://${aws_route53_record.beanstalkappenv.name}/api"
  
}
