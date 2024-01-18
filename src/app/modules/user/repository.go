package user

import "gorm.io/gorm"

type Repository interface {
	Save(user *User)
	FindByEmail(email string) (*User, error)
	FindById(id int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user *User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Find(&user, "id=?", id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
