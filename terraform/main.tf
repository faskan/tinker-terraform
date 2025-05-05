terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "random_id" "suffix" {
  byte_length = 4
}

resource "aws_s3_bucket" "example" {
  bucket = "tofu-s3-bucket-${random_id.suffix.hex}"
  tags = {
    Environment = var.env
  }
}

output "bucket_name" {
  value = aws_s3_bucket.example.bucket
}

variable "env" {
  default = "test"
}
