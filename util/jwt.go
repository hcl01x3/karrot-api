package util

import (
	"fmt"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var bearerRegexp = regexp.MustCompile(`^\s*Bearer\s+`)

type JWTClaims struct {
	jwt.StandardClaims

	UserId int64 `json:"userId"`
}

func EncodeJWT(id int64, issuer, secret string, expiry time.Duration) (string, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(expiry)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			Issuer:    issuer,
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}).SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("encode jwt: signed secret: %w", err)
	}
	return token, nil
}

func DecodeJWT(token, secret string) (*JWTClaims, error) {
	claims := JWTClaims{}
	parser := jwt.Parser{
		SkipClaimsValidation: true,
	}
	_, err := parser.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("decode jwt: parse claims: %w", err)
	}
	return &claims, nil
}
