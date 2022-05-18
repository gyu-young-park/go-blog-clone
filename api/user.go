package api

import (
	"fmt"
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
		sendErrorMessage(503, "get all user data", err, c)
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
		sendErrorMessage(503, "get user", err, c)
		return
	}
	user, err := server.store.GetUserById(int64(id))
	if err != nil {
		sendErrorMessage(503, "get user", err, c)
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
		sendErrorMessage(404, "bind user input data", err, c)
		return
	}
	if req.Username == "" || req.Password == "" || req.Email == "" {
		err := fmt.Errorf("Invalid input key")
		sendErrorMessage(404, "bind user input data", err, c)
		return
	}
	user, err := server.store.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		sendErrorMessage(503, "insert user", err, c)
		return
	}

	res.Username = user.Name
	res.Email = user.Email
	c.JSON(200, res)
	return
}

type deleteUserRequest struct {
	userDataResponse
}

type deleteUserResponse struct {
	Result string `json:"result"`
}

func (server *Server) deleteUser(c *gin.Context) {
	var req deleteUserRequest
	if err := c.ShouldBind(&req); err != nil {
		sendErrorMessage(404, "delete user", err, c)
		return
	}
	id, err := server.store.GetUserIdByEmail(req.Email)
	if err != nil {
		sendErrorMessage(503, "delete user", err, c)
		return
	}
	err = server.store.DeleteUser(id)
	if err != nil {
		sendErrorMessage(503, "delete user", err, c)
		return
	}
	var res deleteUserResponse
	res.Result = "success"
	c.JSON(200, res)
}

type updateUserInfoRequest struct {
	PreEmail string `json:"pre_email"`
	Email    string `json:"email"`
	Name     string `json:"username"`
}

type updateUserInfoResponse struct {
	userDataResponse
}

func (server *Server) updateUserInfo(c *gin.Context) {
	var req updateUserInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		sendErrorMessage(404, "update user", err, c)
		return
	}
	user, err := server.store.UpdateUserInfo(req.PreEmail, req.Email, req.Name)
	if err != nil {
		sendErrorMessage(503, "update user", err, c)
		return
	}
	c.JSON(200, user)
}
