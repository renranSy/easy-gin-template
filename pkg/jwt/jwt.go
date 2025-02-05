package jwt

import (
	"easy-gin-template/internal/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"regexp"
	"time"
)

type JWT struct {
	key []byte
}

type MyCustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

var Jwt *JWT

func init() {
	Jwt = NewJwt()
}

func NewJwt() *JWT {
	return &JWT{key: []byte(config.CFG.JWT.Key)}
}

func (j *JWT) GenToken(userId string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenStr string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenStr = re.ReplaceAllString(tokenStr, "")
	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
