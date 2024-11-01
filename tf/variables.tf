
variable "solution_stack_name" {
  type = string
  description = "The solution stack name to use for the Elastic Beanstalk environment. Docs: https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platform-history-go.html"
}
variable "tier" {
  type = string
  description = "The tier to use for the Elastic Beanstalk environment e.g. WebServer"
}
 

variable "instance_type" {
  default =  "t2.micro"
  description = "The instance type to use for the Elastic Beanstalk environment e.g. t2.micro"
}

variable "minsize" {
  default = 1
  description = "The minimum number of instances to use for the Elastic Beanstalk environment e.g. 1"
}

variable "maxsize" {
  default = 2
  description = "The maximum number of instances to use for the Elastic Beanstalk environment e.g. 2"
}

variable "certificate_arn" {
  type = string
  description = "The ARN of the certificate to use for the ELB e.g. arn:aws:acm:region:account-id:certificate/certificate-id"
}

variable "elb_policy_name" {
  default = "ELBSecurityPolicy-2016-08"
  description = "The name of the ELB policy to use e.g. ELBSecurityPolicy-2016-08"
  
}

variable "hosted_zone" {
  type = string
  description = "The hosted zone for the project e.g. example.com"
}

variable "project_name" {
  type = string
  description = "The name of the project e.g. project-name"
}

variable "elastic_beanstalk_env" {
  type = map(string)
  description = "The environment variables for the Elastic Beanstalk environment e.g. { \"key\" = \"value\" }"
}

variable "codebuild_env" {
  type = map(string)
  description = "The environment variables to use for the build e.g. { \"key\" = \"value\" }"
}

variable "build_image" {
  type = string
  default = "aws/codebuild/standard:7.0"
}

variable "repository_id" {
  type = string
  description = "The ID of the repository to use for the build e.g nedssoft/repository-name. nedssoft is the Github user name"
}

variable "branch_name" {
  type = string
  description = "The branch name to use for the build e.g. main"
}

variable "repository_url" {
  type = string
  description = "The URL of the repository to use for the build e.g. https://github.com/nedssoft/repository-name.git"
}