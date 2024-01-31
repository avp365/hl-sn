package repositories

import (
	"github.com/avp365/hl-sn/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

var UsrRep UserRepository

type UserRepository struct {
	DBPostr *pgxpool.Pool
}

func InitUserRepository(db *pgxpool.Pool) error {
	UsrRep = UserRepository{db}

	return nil
}
func (r *UserRepository) CreateUser(user entities.User) (int, error) {
	return 10, nil
}
