package repositories

import (
	"context"
	"errors"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

var tableNameFriends = "friends"

func (r *UserRepository) GetFriends(userId int) ([]int, error) {

	query := `SELECT id_user_1, id_user_2 FROM ` + tableNameFriends + ` where id_user_1=$1 or id_user_2=$1`

	rows, err := r.DBPostr.Query(context.Background(), query, userId)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return nil, err
	}

	var friendIds []int

	for rows.Next() {

		var id1, id2 int

		rows.Scan(&id1, &id2)

		if id1 != userId {
			friendIds = append(friendIds, id1)
		}

		if id2 != userId {
			friendIds = append(friendIds, id2)
		}

	}

	return friendIds, nil
}
func (r *UserRepository) friendSetCheck(friendSet entities.FriendSet) (int, error) {

	query := `SELECT id FROM ` + tableNameFriends + ` where id_user_1=$1 and id_user_2=$2 or id_user_1=$2 and id_user_2=$1`

	var setId int

	err := r.DBPostr.QueryRow(context.Background(), query, friendSet.IdUser1, friendSet.IdUser2).Scan(&setId)

	if err != nil {

		return 0, err
	}

	return setId, nil

}
func (r *UserRepository) FriendSet(friendSet entities.FriendSet) (int, error) {

	if friendSet.IdUser1 == friendSet.IdUser2 {

		return 0, errors.New("User id = friend id")
	}

	setId, err := r.friendSetCheck(friendSet)

	if err == nil {

		return setId, nil
	}

	query := `INSERT INTO ` + tableNameFriends + ` (id_user_1, id_user_2, date_add) VALUES (@IdUser1, @IdUser2, @DateAdd) returning (id)`

	args := pgx.NamedArgs{
		"IdUser1": friendSet.IdUser1,
		"IdUser2": friendSet.IdUser2,
		"DateAdd": friendSet.DateAdd,
	}

	err = r.DBPostr.QueryRow(context.Background(), query, args).Scan(&setId)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return 0, err
	}

	return setId, nil
}
func (r *UserRepository) FriendDelete(friendSet entities.FriendSet) (bool, error) {

	setId, err := r.friendSetCheck(friendSet)

	if err != nil {

		return true, nil
	}

	query := `DELETE FROM ` + tableNameFriends + ` where id=$1`

	_, err = r.DBPostr.Query(context.Background(), query, setId)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return false, err
	}

	return true, nil
}
