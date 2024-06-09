package dto

import (
	"goweb/model"
)

//type UserLoginDTO struct {
//	Name     string `json:"name" binding:"required,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
//	Password string `json:"password" binding:"required" message:"密码不能为空"`
//}

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" binding:"required" message:"密码不能为空"`
}

func (this *UserAddDTO) ConvertToModule(user *model.User) {
	user.Name = this.Name
	user.RealName = this.RealName
	user.Mobile = this.Mobile
	user.Name = this.Name
	user.Email = this.Email
	user.Password = this.Password
}

type UserListDTO struct {
	Paginate
}
