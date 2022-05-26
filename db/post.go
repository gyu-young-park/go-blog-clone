package db

import (
	"log"
	"time"
)

type Post struct {
	PostId     int64
	UserId     int64
	Email      string
	Content    string
	Tag        string
	Created_at time.Time
	Updated_at time.Time
}

// INSERT INTO Post (userId, email, content, tag) VALUES(1,"ewqe","ewqe","eqw");
func (store *Store) GetPostById(postId int64) (Post, error) {
	var post Post
	err := store.db.QueryRow("SELECT userId, email, content, tag, created_at, updated_at FROM Post WHERE postId = ?", postId).Scan(&post.UserId, &post.Email, &post.Content, &post.Tag, &post.Created_at, &post.Updated_at)
	if err != nil {
		log.Printf("failed to get post data[%v]", err.Error())
		return post, err
	}
	post.PostId = postId
	return post, nil
}
