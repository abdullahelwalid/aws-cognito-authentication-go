package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)


type awsInitializer interface {
	initfunc() (any any)
}

type CognitoAuth struct {
	Cfg aws.Config
	UserPoolID      string
	AppClientID     string
	AppClientSecret string
}

func InitAWSConfig() (*CognitoAuth, error){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	awsAuth := &CognitoAuth{
		Cfg: cfg,
		UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
		AppClientID:     os.Getenv("COGNITO_APP_CLIENT_ID"),
		AppClientSecret: os.Getenv("COGNITO_APP_CLIENT_SECRET"),
	}
	return awsAuth, nil
}

func computeSecretHash(clientSecret string, username string, clientId string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientId))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (c *CognitoAuth) ValidateToken(token string) (*cognitoidentityprovider.GetUserOutput, error) {
	client := cognitoidentityprovider.NewFromConfig(c.Cfg)
	getUserInputFields := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(token),
	}
	resp, err := client.GetUser(context.TODO(), getUserInputFields)
	fmt.Println(*resp.Username)	
	return resp, err
}

func (c *CognitoAuth) Login(email string, password string) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	client := cognitoidentityprovider.NewFromConfig(c.Cfg)
	authParams := map[string]string{"USERNAME": email, "PASSWORD": password, "SECRET_HASH": computeSecretHash(c.AppClientSecret, email, c.AppClientID)}
	signInInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: &c.AppClientID,
		AuthParameters: authParams,
	}
	resp, err := client.InitiateAuth(context.TODO(), signInInput)
	fmt.Println("response: ", err)
	return resp, err
}

func (c *CognitoAuth) ConfirmSignUp(email string, code string) error {
	client := cognitoidentityprovider.NewFromConfig(c.Cfg)
	confirmSignUpInput := &cognitoidentityprovider.ConfirmSignUpInput{
		Username: aws.String(email),
		ClientId: &c.AppClientID,
		ConfirmationCode: aws.String(code),
		SecretHash: aws.String(computeSecretHash(c.AppClientSecret, email, c.AppClientID)),
	}
	resp, err := client.ConfirmSignUp(context.TODO(), confirmSignUpInput)
	fmt.Println(resp)
	return err
}

func (c *CognitoAuth) Signup(email string, password string) error {
	client := cognitoidentityprovider.NewFromConfig(c.Cfg)
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: &c.AppClientID,
		Username: aws.String(email),
		Password: aws.String(password),
		SecretHash: aws.String(computeSecretHash(c.AppClientSecret, email, c.AppClientID)),
	}
	resp, err := client.SignUp(context.TODO(), signUpInput)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	return nil
}
