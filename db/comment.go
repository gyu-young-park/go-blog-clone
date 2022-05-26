package db

import (
	"log"
	"time"
)

type Comment struct {
	CommentId int64
	PostId    int64
	Username  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// INSERT INTO Comment(postId, user_name, CONTENT) VALUES(2, "gyu", "hello");
func (store *Store) GetCommentById(commentId int64) (Comment, error) {
	var comment Comment
	err := store.db.QueryRow("SELECT postId, user_name, content, created_at, updated_at FROM Comment WHERE commentId = ?", commentId).Scan(&comment.PostId, &comment.Username, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		log.Printf("failed to get comment data[%v]", err.Error())
		return comment, err
	}
	comment.CommentId = commentId
	return comment, nil
}
