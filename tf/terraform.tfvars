project_name = "go-mongo"
hosted_zone = "example.com"
instance_type       = "t2.medium"
minsize             = 1
maxsize             = 4
tier = "WebServer"
solution_stack_name= "64bit Amazon Linux 2023 v4.1.5 running Go 1"
certificate_arn = ""

elastic_beanstalk_env = {
  "MONGO_URI" = ""
  "DB_NAME" = "test-go"
  "PORT" = 5000
}

codebuild_env = {
  "KEY" = "VALUE"
}

repository_id = "nedssoft/go-api-mongo"
branch_name = "main"
repository_url = "https://github.com/nedssoft/go-api-mongo.git"
