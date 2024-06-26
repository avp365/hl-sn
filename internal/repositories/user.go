package repositories

import (
	"context"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/pkg/password"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var UsrRep UserRepository
var tableName = "users"

type UserRepository struct {
	DBPostr   *pgxpool.Pool
	DBPostrS1 *pgxpool.Pool
}

func InitUserRepository(dbMaster *pgxpool.Pool, dbs1 *pgxpool.Pool) error {
	UsrRep = UserRepository{dbMaster, dbs1}

	return nil
}
func (r *UserRepository) CreateUser(user entities.User) (int, error) {
	password, _ := password.HashPassword(user.Password)

	query := `INSERT INTO ` + tableName + ` (first_name, second_name, birthdate, biography, city, password) VALUES (@FirstName, @SecondName, @Birthdate, @Biography, @City, @Password) returning (id)`

	args := pgx.NamedArgs{
		"FirstName":  user.FirstName,
		"SecondName": user.SecondName,
		"Birthdate":  user.Birthdate,
		"Biography":  user.Biography,
		"City":       user.City,
		"Password":   password,
	}
	var user_id int
	err := r.DBPostr.QueryRow(context.Background(), query, args).Scan(&user_id)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return 0, err
	}

	return user_id, nil
}
func (r *UserRepository) UserGetById(userid int) (entities.User, error) {

	query := `SELECT id, first_name, second_name, birthdate, biography, city FROM ` + tableName + ` where id=$1`

	var user entities.User
	err := r.DBPostrS1.QueryRow(context.Background(), query, userid).Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Birthdate, &user.Biography, &user.City)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return user, err
	}

	return user, nil
}
func (r *UserRepository) UserSearch(firstName string, lastName string) ([]entities.User, error) {

	query := `SELECT id, first_name, second_name, birthdate, biography, city FROM ` + tableName + ` where first_name LIKE $1 AND second_name LIKE $2 ORDER BY id`

	rows, err := r.DBPostrS1.Query(context.Background(), query, firstName+"%", lastName+"%")
	if err != nil {
		log.Printf("db error Query: %v\n", err)
		return []entities.User{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.User])

	if err != nil {
		log.Printf("db error CollectRows: %v\n", err)
		return []entities.User{}, err
	}

	return users, nil
}
func (r *UserRepository) UserPasswordGetById(userid int) (entities.User, error) {

	query := `SELECT id, password FROM ` + tableName + ` where id=$1`

	var user entities.User
	err := r.DBPostr.QueryRow(context.Background(), query, userid).Scan(&user.ID, &user.Password)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return user, err
	}

	return user, nil
}
