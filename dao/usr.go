package dao

import (
	"fmt"
	"goweb/dto"
	"goweb/model"
)

var userDao *User

type User struct {
	BaseDao
}

// NewUser creates a new instance of User DAO
func NewUser() *User {
	if userDao == nil {
		userDao = &User{NewBase()}
	}
	return userDao
}

// GetUserByNamePassword retrieves a user by username and password
func (me *User) GetUserByNamePassword(stUserName, stPassword string) model.User {
	var user model.User
	fmt.Println(user.Name) // This line seems unnecessary, you can remove it
	me.Orm.Model(&user).Where("name = ? and password = ?", stUserName, stPassword).Find(&user)
	return user
}

// AddUser adds a new user to the database
func (me *User) AddUser(userAddDTO *dto.UserAdd) error {
	var user model.User
	userAddDTO.ConvertToModule(&user) // Convert DTO to model
	err := me.Orm.Save(&user).Error
	if err == nil {
		userAddDTO.ID = user.ID  // Update ID in DTO
		userAddDTO.Password = "" // Clear password in DTO
	}
	return err
}

// CheckUserNameExist checks if a username already exists
func (me *User) CheckUserNameExist(stUserName string) bool {
	var n int64
	me.Orm.Model(&model.User{}).Where("name = ?", stUserName).Count(&n)
	return n > 0
}

// GetUserByName retrieves a user by username
func (me *User) GetUserByName(name string) (model.User, error) {
	var user model.User
	err := me.Orm.Model(&user).Where("name = ?", name).Find(&user).Error
	return user, err
}

// GetUserById retrieves a user by ID
func (me *User) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := me.Orm.First(&user, id).Error // Use First to retrieve by ID
	return user, err
}

// GetUserList retrieves a list of users with pagination
func (me *User) GetUserList(userListDTO *dto.UserList) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	err := me.Orm.Model(&model.User{}).Scopes(Paginate(userListDTO.Paginate)).Find(&userList).Offset(-1).Count(&total).Error
	return userList, total, err
}
