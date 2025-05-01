# AWS Cognito Authentication in Go

This project demonstrates how to integrate **AWS Cognito User Pools** for secure authentication in a Go-based web application. It showcases how to protect APIs using `access_token` and `refresh_token`, and how to manage authentication flows such as sign-up, sign-in, token refresh, and authorization.

## 🚀 Features

- User authentication via AWS Cognito
- Secure middleware to protect routes
- Token validation and refreshing
- Dockerized setup for easy deployment
- Utility packages for AWS interactions and database connectivity

## 🗂 Project Structure

- `Dockerfile`: Docker image for the Go app
- `docker-compose.yml`: Multi-container setup
- `go.mod`: Go module dependencies
- `go.sum`: Go module dependency checksums
- `README.md`: Project documentation (this file)
- `pkg/middleware/auth.go`: Middleware for Cognito token verification
- `pkg/utils/db.go`: Database connection utilities
- `pkg/utils/aws.go`: AWS Cognito helper functions
- `pkg/models/userModel.go`: User model definition

## 🛠 Setup Instructions

### Prerequisites

- [Go (1.18+)](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- AWS account with a Cognito User Pool configured



```bash
git clone https://github.com/abdullahelwalid/aws-cognito-authentication-go.git
cd aws-cognito-authentication-go
export AWS_REGION=your_aws_region       # e.g., us-east-1
export COGNITO_USER_POOL_ID=your_user_pool_id
export COGNITO_APP_CLIENT_ID=your_app_client_id
export DB_CONNECTION_STRING=your_db_connection_string
docker-compose up --build
```
