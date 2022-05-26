package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) postRouterSetting() {
	server.router.GET("/post/:id", server.getPostDataById)
}

type getPostDataByIdResponse struct {
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	Tag       string    `json:"tag"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (server *Server) getPostDataById(c *gin.Context) {
	var res getPostDataByIdResponse
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("Invalid input key")
		sendErrorMessage(404, "get post data", err, c)
	}
	post, err := server.store.GetPostById(int64(postId))
	if err != nil {
		sendErrorMessage(404, "get post data", err, c)
		return
	}
	res.Content = post.Content
	res.Email = post.Email
	res.Tag = post.Tag
	res.Email = post.Email
	res.CreatedAt = post.Created_at
	res.UpdatedAt = post.Updated_at

	c.JSON(200, post)
}
