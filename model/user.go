package model

import (
	"gorm.io/gorm"
	"goweb/utils"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	RealName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Mobile   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"128;not null"`
}

type LoginUser struct {
	ID   string
	Name string
}

func (me *User) Encrypt() error {
	hash, err := utils.Encrypt(me.Password)
	if err == nil {
		me.Password = hash
	}
	return err
}

func (me *User) BeforeCreate(orm *gorm.DB) error {
	return me.Encrypt()
}
