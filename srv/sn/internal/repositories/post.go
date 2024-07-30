package repositories

import (
	"context"

	"github.com/avp365/hl-sn/srv/sn/internal/entities"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var PostRep PostRepository
var tableNamePosts = "posts"

type PostRepository struct {
	DBPostr   *pgxpool.Pool
	DBPostrS1 *pgxpool.Pool
}

func InitPostRepository(dbMaster *pgxpool.Pool, dbs1 *pgxpool.Pool) error {
	PostRep = PostRepository{dbMaster, dbs1}
	return nil
}

func (r *PostRepository) CreatePost(post entities.Post) (int, error) {

	query := `INSERT INTO ` + tableNamePosts + ` (user_id, text, date) VALUES (@UserId, @Text, @Date) returning (id)`

	args := pgx.NamedArgs{
		"UserId": post.UserID,
		"Text":   post.Text,
		"Date":   post.Date,
	}

	var post_id int

	err := r.DBPostr.QueryRow(context.Background(), query, args).Scan(&post_id)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return 0, err
	}

	return post_id, nil
}

func (r *PostRepository) Delete(postId int, userId int) (bool, error) {

	query := `DELETE FROM ` + tableNamePosts + ` where id=$1 and user_id=$2`

	_, err := r.DBPostr.Exec(context.Background(), query, postId, userId)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *PostRepository) PostUpdate(post entities.Post) (int, error) {

	query := `UPDATE ` + tableNamePosts + ` SET text=$1 , date=$2 where id=$3 and user_id=$4 returning id`

	_, err := r.DBPostr.Exec(context.Background(), query, post.Text, post.Date, post.ID, post.UserID)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return 0, err
	}

	return 1, nil
}

func (r *PostRepository) Get(postId int, userId int) (entities.Post, error) {

	var post entities.Post

	query := `SELECT id, user_id, text, date FROM ` + tableNamePosts + ` where id=$1 and user_id=$2`

	err := r.DBPostr.QueryRow(context.Background(), query, postId, userId).Scan(&post.ID, &post.UserID, &post.Text, &post.Date)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return post, err
	}

	return post, nil
}

func (r *PostRepository) Feed(userId int) ([]entities.Post, error) {

	var posts []entities.Post

	query := `SELECT p.id, p.user_id, p.text, p.date FROM friends f JOIN posts p ON f.id_user_2 = p.user_id where f.id_user_1=$1 and id_user_1!=user_id order by date_add DESC limit 1000;`

	rows, err := r.DBPostr.Query(context.Background(), query, userId)

	if err != nil {
		log.Printf("db error: %v\n", err)
		return posts, err
	}

	for rows.Next() {
		var post entities.Post
		rows.Scan(&post.ID, &post.UserID, &post.Text, &post.Date)
		posts = append(posts, post)
	}

	return posts, nil
}
