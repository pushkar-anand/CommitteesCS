package security

import (
	"committees/config"
	"committees/helpers"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	// ErrorSigningKey is thrown when the signing method is invalid
	ErrorSigningKey = errors.New("unexpected signing method")
)

const (
	// KeySubject stores the subject of the JWT token
	KeySubject = "sub"
	// KeyIssuer stores the issuer details of the JWT token
	KeyIssuer = "key"
	// KeyID stores the unique JWT token id
	KeyID = "jti"
	// KeyNotBefore stores the timestamp of the token should not be used before
	KeyNotBefore = "nbf"
	// KeyAudience stores the audience for whome the token is generated
	KeyAudience = "aud"
	// KeyIssuedAt stores the timestamp at which the token was issued
	KeyIssuedAt = "iat"
	// KeyExpiry stores the timestamp at which the token will expire
	KeyExpiry = "exp"
	// KeyTokenType stores the type of the token
	KeyTokenType = "type"
)

type tokenType struct {
	TypeName          string
	ExpirationMinutes int
}

/*var AccessToken = &tokenType{
	TypeName:          "access_token",
	ExpirationMinutes: 15,
}

var RefreshToken = &tokenType{
	TypeName:          "refresh_token",
	ExpirationMinutes: 365 * 24 * 60,
}*/

var secret []byte

func init() {
	secret = []byte(config.GetAppConfig().JWTSecret)
}

var audience = []string{"api"}

func buildClaim(tokenType *tokenType, m map[string]interface{}) jwt.MapClaims {
	var mapClaim = defaultClaim()

	if m != nil {
		for key, value := range m {
			mapClaim[key] = value
		}
	}

	mapClaim[KeyTokenType] = tokenType.TypeName
	mapClaim[KeyExpiry] = helpers.UNIXTimestampFromNow(tokenType.ExpirationMinutes)

	return mapClaim
}

func jwtKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, ErrorSigningKey
	}
	return secret, nil
}

func defaultClaim() jwt.MapClaims {
	return jwt.MapClaims{
		KeySubject:   "user",
		KeyIssuer:    "fold",
		KeyID:        uuid.New().String(),
		KeyNotBefore: time.Now().Unix(),
		KeyAudience:  audience,
		KeyIssuedAt:  time.Now().Unix(),
	}
}

// NewToken creates a new JWT signed token
func NewToken(tokenType *tokenType, m map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, buildClaim(tokenType, m))
	return token.SignedString(secret)
}

// ReadToken reads and validates a signed token
func ReadToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, jwtKey)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetClaim returns the JWT claim data
func GetClaim(token *jwt.Token) jwt.MapClaims {
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claim
	}
	return nil
}

// IsTokenExpiredError check if the error retuned by
// ReadToken function is thrown because the token expired
func IsTokenExpiredError(err error) bool {
	if validationErrs, ok := err.(jwt.ValidationError); ok {
		return (jwt.ValidationErrorExpired & validationErrs.Errors) != 0
	}

	return false
}

// ReadAndGetClaim reads and validates a signed token
// returns the claim data on success
func ReadAndGetClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := ReadToken(tokenStr)
	if err != nil {
		return nil, err
	}
	return GetClaim(token), nil
}
