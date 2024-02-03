package token

import (
	"errors"
	"strings"

	"github.com/avp365/hl-sn/internal/pkg/configs"
	"github.com/golang-jwt/jwt"
)

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(configs.GetTokenSecretKey()), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}
