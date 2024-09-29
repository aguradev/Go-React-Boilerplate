package repository

import "gorm.io/gorm"

type UserRepoInterface interface {
	Create()
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create() {

}
