package dao

import (
	"fmt"
	"goweb/model"
	"goweb/router/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}
	return userDao
}

func (me *UserDao) GetUserByNamePassword(stUserName, stPassword string) model.User {
	var user model.User
	fmt.Println(user.Name)
	me.Orm.Model(&user).Where("name = ? and password = ?", stUserName, stPassword).Find(&user)
	return user
}

func (me *UserDao) AddUser(userAddDTO *dto.UserAddDTO) error {
	var user model.User
	userAddDTO.ConvertToModule(&user)
	err := me.Orm.Save(&user).Error
	if err == nil {
		userAddDTO.ID = user.ID
		userAddDTO.Password = ""
	}
	return err
}

func (me *UserDao) CheckUserNameExist(stUserName string) bool {
	var n int64
	me.Orm.Model(&model.User{}).Where("name = ?", stUserName).Count(&n)
	return n > 0
}

func (me *UserDao) GetUserByName(name string) (model.User, error) {
	var user model.User
	err := me.Orm.Model(&user).Where("name = ?", name).Find(&user).Error
	return user, err
}

func (me *UserDao) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := me.Orm.First(&user, id).Error
	return user, err
}

func (me *UserDao) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	err := me.Orm.Model(&model.User{}).Scopes(Paginate(userListDTO.Paginate)).Find(&userList).Offset(-1).Count(&total).Error
	return userList, total, err
}
