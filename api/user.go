package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gyu-young-park/go_blog/util"
)

type getAllUserDataResponse struct {
	Data []userDataResponse `json:"data"`
}

func logUserRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Get user api request")
	}
}

func (server *Server) userRouterSetting() {
	userRouter := server.router.Group("/user")
	userRouter.Use(logUserRequest())
	userRouter.GET("/", server.getAllUserData)
	userRouter.GET("/:id", server.getUserData)
	userRouter.POST("/", server.registerUser)
	userRouter.POST("/delete", server.deleteUser)
	userRouter.POST("/update", server.updateUserInfo)
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

type userDataRequest struct {
	Username string `json:"username" binding:"required,min=6,max=10"`
	Email    string `json:"email" binding:"required,email"`
}

type userDataResponse struct {
	Username string
	Email    string
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
	userDataRequest
	Password string `json:"password" binding:"required,min=6"`
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
	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		sendErrorMessage(503, "insert user", err, c)
	}
	user, err := server.store.RegisterUser(req.Username, req.Email, hashedPassword)
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
	userDataRequest
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
	userDataRequest
	PreEmail string `json:"pre_email" binding:"required,email"`
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
	user, err := server.store.UpdateUserInfo(req.PreEmail, req.Email, req.Username)
	if err != nil {
		sendErrorMessage(503, "update user", err, c)
		return
	}
	c.JSON(200, user)
}
