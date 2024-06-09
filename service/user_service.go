package service

import (
	"errors"
	"goweb/dao"
	"goweb/model"
	"goweb/router/dto"
	"goweb/utils"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (me*UserService) Login(userDTO dto.UserLoginDTO) (model.User, string, error) {
	var errRes error
	var token string
	user, err := me.Dao.GetUserByName(userDTO.Name)
	if err != nil || !utils.CompareHashAndPassword(user.Password, userDTO.Password) {
		errRes = errors.New("invalid username or password")
	} else {
		token, _ = utils.GenerateToken(user.ID, user.Name)
	}

	return user, token, errRes
}

func (me *UserService) AddUser(userAddDTO *dto.UserAddDTO) error {
	if me.Dao.CheckUserNameExist(userAddDTO.Name) {
		return errors.New("user name exist")
	}
	return me.Dao.AddUser(userAddDTO)
}

func (me *UserService) GetUserById(commIDDTO *dto.CommIDDTO) (model.User, error) {
	return me.Dao.GetUserById(commIDDTO.ID)
}

func (me *UserService) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return me.Dao.GetUserList(userListDTO)
}
