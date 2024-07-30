package post

import (
	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/repositories"
)

func CreatePost(post entities.Post) (int, error) {
	return repositories.PostRep.CreatePost(post)
}
func Delete(postId int, userId int) (bool, error) {
	return repositories.PostRep.Delete(postId, userId)
}
func PostUpdate(post entities.Post) (int, error) {
	return repositories.PostRep.PostUpdate(post)
}
func Get(postId int, userId int) (entities.Post, error) {
	return repositories.PostRep.Get(postId, userId)
}
func Feed(userId int) ([]entities.Post, error) {
	return repositories.PostRep.Feed(userId)
}
