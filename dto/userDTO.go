/*
Package dto defines data transfer objects (DTOs) for receiving data from web requests.

DTOs are used to encapsulate data that is sent from the web client to the server.
They provide a structured way to represent the data, making it easier to validate
and process on the server side.
*/
package dto

import (
	"goweb/model"
)

//type UserLoginDTO struct {
//	Name     string `json:"name" binding:"required,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
//	Password string `json:"password" binding:"required" message:"密码不能为空"`
//}

// UserLogin represents the data structure for user login requests.
type UserLogin struct {
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

// UserAdd represents the data structure for adding a new user.
type UserAdd struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" binding:"required" message:"密码不能为空"`
}

// UserList represents the data structure for retrieving a list of users.
type UserList struct {
	Paginate
}

// ConvertToModule converts the UserAdd DTO to a User model.
func (me *UserAdd) ConvertToModule(user *model.User) {
	user.Name = me.Name
	user.RealName = me.RealName
	user.Mobile = me.Mobile
	user.Name = me.Name
	user.Email = me.Email
	user.Password = me.Password
}
