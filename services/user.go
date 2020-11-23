package services

import (
	"github.com/DmitryKuzmenec/CoolTeacher/model"
	"github.com/DmitryKuzmenec/CoolTeacher/repositories"
)

// UserService - service of user
type UserService struct {
	user *repositories.UserRepo
}

// NewUserService - constructor of UserService
func NewUserService(user *repositories.UserRepo) *UserService {
	return &UserService{
		user: user,
	}
}

// Create - create user
func (s *UserService) Create(user *model.User) (int, error) {
	return s.user.Add(user)
}

// Get - get user
func (s *UserService) Get(id int) (*model.User, error) {
	return s.user.Get(id)
}

// List - get all users
func (s *UserService) List() ([]*model.User, error) {
	return s.user.GetAll()
}

// Delete - delete user
func (s *UserService) Delete(id int) error {
	return s.user.Delete(id)
}

// Update - update user
func (s *UserService) Update(user *model.User) error {
	return s.user.Update(user)
}
