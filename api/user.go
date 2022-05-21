package api

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gyu-young-park/go_blog/util"
)

type getAllUserDataResponse struct {
	Data []userDataResponse `json:"data"`
}

func (server *Server) userRouterSetting() {
	server.router.POST("/user/login", server.loginUser)
	userRouter := server.router.Group("/user")
	userRouter.Use(authMiddleware(server.tokenMaker))
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

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
}

func (server *Server) loginUser(c *gin.Context) {
	var req loginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendErrorMessage(404, "login user", err, c)
		return
	}

	userId, err := server.store.GetUserIdByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorMessage(404, "login user", err, c)
			return
		}
		sendErrorMessage(503, "login user", err, c)
		return
	}

	user, err := server.store.GetUserById(userId)
	if err != nil {
		sendErrorMessage(503, "login user", err, c)
		return
	}

	ok := util.CheckPassword(req.Password, user.Password)
	if !ok {
		sendErrorMessage(401, "login user not match password", err, c)
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		time.Duration(time.Minute),
	)

	res := loginUserResponse{
		AccessToken: accessToken,
		Email:       user.Email,
	}
	c.Set("httpOnly", true)
	c.Set("Secure", true)
	c.JSON(200, res)
}
