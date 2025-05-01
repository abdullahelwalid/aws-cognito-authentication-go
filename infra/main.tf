terraform {
  required_version = ">= 1.5.7"
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


resource "aws_cognito_user_pool" "user_pool" {
  name = var.user_pool_name
  admin_create_user_config {
    allow_admin_create_user_only = false
  }
  deletion_protection = "ACTIVE"
  user_pool_add_ons {
    advanced_security_mode = "OFF"
  }
  username_attributes      = ["email", "phone_number"]
  auto_verified_attributes = ["email"]
  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
    recovery_mechanism {
      name     = "verified_phone_number"
      priority = 2
    }
  }

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }
}

resource "aws_cognito_user_pool_client" "user_pool_client" {
  depends_on                   = [aws_cognito_user_pool.user_pool]
  name                         = var.user_pool_client_name
  user_pool_id                 = aws_cognito_user_pool.user_pool.id
  generate_secret              = true
  access_token_validity        = 1
  refresh_token_validity       = 90
  explicit_auth_flows          = ["USER_PASSWORD_AUTH", "ADMIN_NO_SRP_AUTH", "USER_PASSWORD_AUTH"]
  supported_identity_providers = ["COGNITO"]
}
