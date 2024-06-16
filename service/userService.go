// Package service provides business logic services for the application.
//
// Services encapsulate the logic for interacting with the data layer (DAOs)
// and performing specific business operations. They provide a layer of
// abstraction between the controllers and the DAOs, making the code more
// modular and maintainable.
package service

import (
	"errors"
	"goweb/dao"
	"goweb/dto"
	"goweb/model"
	"goweb/utils"
)

// userService is a global variable to store the singleton instance of UserService.
var userService *UserService

// UserService represents a service for user-related operations.
type UserService struct {
	// BaseService provides base functionality for services.
	BaseService
	// Dao is the DAO for user data.
	Dao *dao.User
}

// NewUserService creates a new instance of UserService.
// It ensures that only one instance of the service exists using a singleton pattern.
func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUser(),
		}
	}
	return userService
}

// Login attempts to log in a user with the given username and password.
// It checks if the user exists and if the password matches the hashed password
// stored in the database. If successful, it generates a JWT token for the user.
func (me *UserService) Login(userDTO dto.UserLogin) (model.User, string, error) {
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

// AddUser adds a new user to the database.
// It checks if the username already exists and if not, it saves the user data.
func (me *UserService) AddUser(userAddDTO *dto.UserAdd) error {
	if me.Dao.CheckUserNameExist(userAddDTO.Name) {
		return errors.New("user name exist")
	}
	return me.Dao.AddUser(userAddDTO)
}

// GetUserById retrieves a user from the database by their ID.
func (me *UserService) GetUserById(commIDDTO *dto.CommID) (model.User, error) {
	return me.Dao.GetUserById(commIDDTO.ID)
}

// GetUserList retrieves a list of users from the database.
func (me *UserService) GetUserList(userListDTO *dto.UserList) ([]model.User, int64, error) {
	return me.Dao.GetUserList(userListDTO)
}
