package repositories

import (
	"github.com/DmitryKuzmenec/CoolTeacher/client/database"
	"github.com/DmitryKuzmenec/CoolTeacher/model"
)

// UserRepo - type of repository
type UserRepo struct {
	db *database.DB
}

// NewUserRepo - constructor of UserRepo
func NewUserRepo(db *database.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Add - add new user
func (r *UserRepo) Add(user *model.User) (int, error) {
	if res := r.db.Create(user); res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

// Get - get user by id
func (r *UserRepo) Get(id int) (*model.User, error) {
	user := model.User{ID: id}
	if res := r.db.First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// Delete - delete user
func (r *UserRepo) Delete(id int) error {
	if res := r.db.Delete(&model.User{ID: id}); res.Error != nil {
		return res.Error
	}
	return nil
}

// Update - update user
func (r *UserRepo) Update(user *model.User) error {
	if res := r.db.Save(user); res.Error != nil {
		return res.Error
	}
	return nil
}

// GetAll - get all users
func (r *UserRepo) GetAll() ([]*model.User, error) {
	users := []*model.User{}
	if res := r.db.Model(&model.User{}).Find(&users); res.Error != nil {
		return users, res.Error
	}
	return users, nil
}
