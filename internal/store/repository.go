package store

import "BeatBoulevard/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	Find(id int) (*model.User, error)
}
