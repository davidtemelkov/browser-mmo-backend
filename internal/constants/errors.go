package constants

// General errors
const (
	RecordNotFoundError           = "record not found"
	EditConflictError             = "edit conflict"
	DBConnectionError             = "db connection error"
	InvalidJSONFormatError        = "invalid JSON format"
	InvalidBase64ImagePrefixError = "invalid base64 image format prefix"
	InternalServerError           = "internal server error"
	BadRequestError               = "bad request"
)

// Env errors
const (
	LoadingEnvFileError     = "error loading .env file"
	AWSAccessKeyError       = "aws access key not found in env file"
	AWSSecretKeyError       = "aws secret key not found in env file"
	JWTPrivateKeyError      = "jwt private key not found in env file"
	FirebaseURLKeyError     = "firebase URL key not found in env file"
	FirebaseBucketNameError = "firebase bucket name not found in env file"
)

// Firebase errors
const (
	FirebaseClientError  = "failed to initialize Firebase Storage client"
	FileFolderEmptyError = "fileFolder is empty"
	FileNameEmptyError   = "fileName is empty"
)

// DynamoDB errors
var (
	DynamoDBClientError = "couldn't create dynamoDB client"
)

// Authentication errors
const (
	MissingAuthorizationHeaderError       = "missing authorization header"
	InvalidAuthorizationHeaderFormatError = "invalid authorization header format"
	InvalidTokenError                     = "invalid token"
	InvalidTokenClaimsError               = "invalid token claims"
)

// User errors
const (
	RequiredFieldError        = "field is required"
	EmailFormatError          = "email must be in the correct email format"
	PasswordMinLengthError    = "password must be at least 8 symbols"
	PasswordMaxLengthError    = "password must be less than 72 symbols"
	UserNameMinLengthError    = "name must be at least 4 symbols"
	UserNameMaxLengthError    = "name must be less than 50 symbols"
	UserNameNoWhitespaceError = "must contain two names seperated by whitespace"
	UserIsNotAuthorizedError  = "user is not authorized"
	DuplicateEmailError       = "duplicate email"
	UserNotFoundError         = "user not found"
	FailedLoginError          = "invalid email or password"
)
