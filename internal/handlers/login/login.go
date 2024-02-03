package login

import (
	"errors"
	"time"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/pkg/configs"
	"github.com/avp365/hl-sn/internal/pkg/password"
	"github.com/avp365/hl-sn/internal/repositories"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(form *entities.LoginForm) (string, error) {

	user, err := repositories.UsrRep.UserPasswordGetById(form.ID)

	if err != nil {
		return "", err
	}

	if password.CheckPasswordHash(form.Password, user.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  form.ID,
			"nbf": time.Date(2024, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
		})

		return token.SignedString([]byte(configs.GetTokenSecretKey()))

	}

	return "", errors.New("password incorrect")

}
