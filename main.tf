terraform {
  required_version = ">= 1.1.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.29"
    }
    local = {
      source = "hashicorp/local"
    }
  }
}
provider "aws" {
    region = var.aws_region
}

resource "aws_vpc" "main" {
  cidr_block           = var.cidr_block
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = var.vpc_name
  }
}

variable "aws_region" {
  default = "eu-central-1"
}

variable "cidr_block" {
  type    = string
  default = "10.0.0.0/16"
}

variable "vpc_name" {
  type    = string
  default = "my-first-vpc"
}

output "vpc_id" {
  value = aws_vpc.main.id
}