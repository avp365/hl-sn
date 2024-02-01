package login

import (
	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/repositories"
)

func LoginHandler(form *entities.LoginForm) (string, error) {

	user, err := repositories.UsrRep.UserPasswordGetById(form.ID)

	return "ok", nil

}
