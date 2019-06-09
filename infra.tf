provider "aws" {
  access_key                  = "mock_access_key"
  region                      = "us-east-1"
  s3_force_path_style         = true
  secret_key                  = "mock_secret_key"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    sqs = "http://localhost:4576"
  }
}

resource "aws_sqs_queue" "playground" {
  name = "playground"

  message_retention_seconds  = 3600
  receive_wait_time_seconds  = 20
  visibility_timeout_seconds = 42
}
