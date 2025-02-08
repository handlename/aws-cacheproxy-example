terraform {
  required_version = "~> 1.10.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0.0"
    }
  }

  backend "local" {}
}

provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      ManagedBy = "terraform"
      For       = "aws-cacheproxy-example"
    }
  }
}
