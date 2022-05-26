package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) commentRouterSetting() {
	server.router.GET("/comment/:id", server.getCommentDataById)
}

type getCommentDataByIdResponse struct {
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (server *Server) getCommentDataById(c *gin.Context) {
	var res getCommentDataByIdResponse
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("Invalid input key")
		sendErrorMessage(404, "get comment data", err, c)
		return
	}
	comment, err := server.store.GetCommentById(int64(commentId))
	if err != nil {
		sendErrorMessage(404, "get comment data", err, c)
		return
	}
	res.Username = comment.Username
	res.Content = comment.Content
	res.CreatedAt = comment.CreatedAt
	res.UpdatedAt = comment.UpdatedAt
	c.JSON(200, res)
}
