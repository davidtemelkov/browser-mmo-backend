package constants

import "errors"

// General errors
var (
	RecordNotFoundError           = errors.New("record not found")
	EditConflictError             = errors.New("edit conflict")
	DBConnectionError             = errors.New("Db connection error")
	InvalidJSONFormatError        = errors.New("Invalid JSON format")
	InvalidBase64ImagePrefixError = errors.New("Invalid base64 image format prefix")
	InternalServerError           = errors.New("Internal server error")
	BadRequestError               = errors.New("Bad request")
)

// Env errors
var (
	LoadingEnvFileError     = errors.New("Error loading .env file")
	AWSAccessKeyError       = errors.New("AWS access key not found in env file")
	AWSSecretKeyError       = errors.New("AWS secret key not found in env file")
	JWTPrivateKeyError      = errors.New("JWT private key not found in env file")
	FirebaseURLKeyError     = errors.New("Firebase URL key not found in env file")
	FirebaseBucketNameError = errors.New("Firebase bucket name not found in env file")
)

// Firebase errors
var (
	FirebaseClientError  = errors.New("Failed to initialize Firebase Storage client")
	FileFolderEmptyError = errors.New("FileFolder is empty")
	FileNameEmptyError   = errors.New("FileName is empty")
)

// DynamoDB errors
var (
	DynamoDBClientError = errors.New("Couldn't create dynamoDB client")
)

// Authentication errors
var (
	MissingAuthorizationHeaderError       = errors.New("Missing authorization header")
	InvalidAuthorizationHeaderFormatError = errors.New("Invalid authorization header format")
	InvalidTokenError                     = errors.New("Invalid token")
	InvalidTokenClaimsError               = errors.New("Invalid token claims")
)

// User errors
var (
	RequiredFieldError        = errors.New("Field is required")
	EmailFormatError          = errors.New("Email must be in the correct email format")
	PasswordMinLengthError    = errors.New("Password must be at least 8 symbols")
	PasswordMaxLengthError    = errors.New("Password must be less than 72 symbols")
	UserNameMinLengthError    = errors.New("Name must be at least 4 symbols")
	UserNameMaxLengthError    = errors.New("Name must be less than 50 symbols")
	UserNameNoWhitespaceError = errors.New("Must contain two names seperated by whitespace")
	UserIsNotAuthorizedError  = errors.New("User is not authorized")
	DuplicateEmailError       = errors.New("Duplicate email")
	UserNotFoundError         = errors.New("User not found")
	FailedLoginError          = errors.New("Invalid email or password")
)
