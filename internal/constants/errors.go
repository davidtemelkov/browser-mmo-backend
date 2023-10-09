package constants

import "errors"

var (
	ErrAWSAccessKey   = errors.New("AWS access key not found in env file")
	ErrAWSSecretKey   = errors.New("AWS secret key not found in env file")
	ErrDynamoDBClient = errors.New("Couldn't create dynamoDB client")
)
