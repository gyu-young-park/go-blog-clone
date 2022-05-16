package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllUserDataResponse struct {
	Data []userDataResponse `json:"data"`
}

func (server *Server) getAllUserData(c *gin.Context) {
	var res getAllUserDataResponse
	users, err := server.store.GetAllUserData()
	if err != nil {
		log.Printf("[DB]failed to get all user data[%v]\n", err)
		c.JSON(503, errResponse(err))
		return
	}
	for _, user := range users {
		userDataRes := userDataResponse{
			Username: user.Name,
			Email:    user.Email,
		}
		res.Data = append(res.Data, userDataRes)
	}
	c.JSON(200, res)
	return
}

type userDataResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (server *Server) getUserData(c *gin.Context) {
	var res userDataResponse
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("failed to get user[%v]\n", err.Error())
		c.JSON(404, errResponse(err))
		return
	}
	user, err := server.store.GetUserById(int64(id))
	if err != nil {
		log.Printf("[DB]failed to get user[%v]\n", err.Error())
		c.JSON(404, errResponse(err))
		return
	}
	res.Email = user.Email
	res.Username = user.Name

	c.JSON(200, res)
	return
}

type registerUserRequest struct {
	userDataResponse
	Password string `json:"password"`
}

type registerUserResponse struct {
	userDataResponse
}

func (server *Server) registerUser(c *gin.Context) {
	var req registerUserRequest
	var res registerUserResponse
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("failed to bind user input data[%v]\n", err.Error())
		c.JSON(404, errResponse(err))
		return
	}
	if req.Username == "" || req.Password == "" || req.Email == "" {
		err := fmt.Errorf("Invalid input key")
		log.Printf("failed to bind user input data[%v]\n", err)
		c.JSON(404, errResponse(err))
		return
	}
	user, err := server.store.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		log.Printf("failed to inset user [%v]\n", err.Error())
		c.JSON(503, errResponse(err))
		return
	}

	res.Username = user.Name
	res.Email = user.Email
	c.JSON(200, res)
	return
}
