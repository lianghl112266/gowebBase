package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goweb/global"
	"goweb/router/dto"
	"goweb/service"
	"net/http"
	"strconv"
	"time"
)

const (
	ErrCodeUser        = 10001
	ErrCodeGetUserById = 10002
	ErrCodeGetUserList = 10003
	ErrCodeLogin       = 10004
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tag 用户管理
// @Summary user login
// @Description
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (me *UserApi) Login(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO

	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userLoginDTO}).GetError(); err != nil {
		return
	}

	user, token, err := me.Service.Login(userLoginDTO)

	if err != nil {
		me.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ErrCodeLogin,
			Msg:    err.Error(),
		})
		return
	}

	//可以让服务器提前下线token，而不必必须等待token过期
	_ = global.RedisClient.Set(strconv.Itoa(int(user.ID)), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)

	me.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (me *UserApi) AddUser(ctx *gin.Context) {
	var userAddDTO dto.UserAddDTO
	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userAddDTO}).GetError(); err != nil {
		return
	}
	err := me.Service.AddUser(&userAddDTO)
	if err != nil {
		me.ServerFail(ResponseJson{
			Code: ErrCodeUser,
			Msg:  err.Error(),
		})
		return
	}

	me.OK(ResponseJson{
		Data: userAddDTO,
	})
}

func (me *UserApi) GetUserById(ctx *gin.Context) {
	var commIDDTO dto.CommIDDTO
	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &commIDDTO, BindParamsFromUrl: true}).GetError(); err != nil {
		return
	}
	user, err := me.Service.GetUserById(&commIDDTO)
	if err != nil {
		me.ServerFail(ResponseJson{
			Code: ErrCodeGetUserById,
			Msg:  err.Error(),
		})
		return
	}
	fmt.Println("ok", user)
	me.OK(ResponseJson{
		Data: user,
	})
}

func (me *UserApi) GetUserList(ctx *gin.Context) {
	var userListDTO dto.UserListDTO
	if err := me.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userListDTO}).GetError(); err != nil {
		return
	}
	userList, total, err := me.Service.GetUserList(&userListDTO)
	if err != nil {
		me.ServerFail(ResponseJson{
			Code: ErrCodeGetUserList,
			Msg:  err.Error(),
		})
		return
	}

	me.OK(ResponseJson{
		Data:  userList,
		Total: total,
	})
}
