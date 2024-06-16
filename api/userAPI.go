package api

import (
	"fmt"
	"goweb/dto"
	"goweb/global"
	"goweb/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Error code definitions
const (
	ErrCodeAddUser     = 10001
	ErrCodeGetUserById = 10002
	ErrCodeGetUserList = 10003
	ErrCodeLogin       = 10004
)

// UserApi struct for user API operations
type UserApi struct {
	// Base provides basic API operations
	Base
	// Service provides user-related services
	Service *service.UserService
}

// NewUserApi creates a new instance of UserApi
func NewUserApi() UserApi {
	return UserApi{
		Base:    NewBase(),
		Service: service.NewUserService(),
	}
}

// @Tag User Management
// @Summary User Login
// @Description User login API endpoint
// @Param name formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Login successful"
// @Failure 10004 {string} string "Login failed"
// @Router /api/v1/public/user/login [post]
func (me UserApi) Login(ctx *gin.Context) {
	// Create a login DTO and Build request parameters
	var userLoginDTO dto.UserLogin
	if err := me.Base.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userLoginDTO}); err != nil {
		return
	}

	// Call the user service to perform login
	user, token, err := me.Service.Login(userLoginDTO)

	// Handle login errors
	if err != nil {
		me.Base.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ErrCodeLogin,
			Msg:    err.Error(),
		})
		return
	}

	// Store the token in Redis
	_ = global.RedisClient.Set(strconv.Itoa(int(user.ID)), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)

	// Return the login success information
	me.Base.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

// @Tag User Management
// @Summary Add User
// @Description API endpoint to add a new user
// @Param userAdd formData dto.UserAdd true "User add information"
// @Success 200 {string} string "User added successfully"
// @Failure 10001 {string} string "Failed to add user"
// @Router /api/v1/public/user/add [post]
func (me UserApi) AddUser(ctx *gin.Context) {
	// Create a user add DTO
	var userAddDTO dto.UserAdd
	// Build request parameters
	if err := me.Base.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userAddDTO}); err != nil {
		return
	}
	// Call the user service to add the user
	err := me.Service.AddUser(&userAddDTO)
	// Handle user addition errors
	if err != nil {
		me.Base.ServerFail(ResponseJson{
			Code: ErrCodeAddUser,
			Msg:  err.Error(),
		})
		return
	}

	// Return the user add success information
	me.Base.OK(ResponseJson{
		Data: userAddDTO,
	})
}

// @Tag User Management
// @Summary Get User by ID
// @Description Get user information by user ID
// @Param id path int true "User ID"
// @Success 200 {object} model.User "User information"
// @Failure 10002 {string} string "Get user by ID failed"
// @Router /api/v1/public/user/{id} [get]
func (me UserApi) GetUserById(ctx *gin.Context) {
	// Create a CommID DTO and Build request parameters from URL
	var commIDDTO dto.CommID
	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &commIDDTO, BindParamsFromUrl: true}); err != nil {
		return
	}

	// Call the user service to get user by ID
	user, err := me.Service.GetUserById(&commIDDTO)

	// Handle errors
	if err != nil {
		me.ServerFail(ResponseJson{
			Code: ErrCodeGetUserById,
			Msg:  err.Error(),
		})
		return
	}

	// Log the success response
	fmt.Println("ok", user)

	// Return the user information
	me.OK(ResponseJson{
		Data: user,
	})
}

// @Tag User Management
// @Summary Get User List
// @Description Get a list of users
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} []model.User "User list"
// @Failure 10003 {string} string "Get user list failed"
// @Router /api/v1/public/users [get]
func (me UserApi) GetUserList(ctx *gin.Context) {
	// Create a UserList DTO and Build request parameters
	var userListDTO dto.UserList
	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userListDTO}); err != nil {
		return
	}

	// Call the user service to get user list
	userList, total, err := me.Service.GetUserList(&userListDTO)

	// Handle errors
	if err != nil {
		me.ServerFail(ResponseJson{
			Code: ErrCodeGetUserList,
			Msg:  err.Error(),
		})
		return
	}

	// Return the user list and total count
	me.OK(ResponseJson{
		Data:  userList,
		Total: total,
	})
}
