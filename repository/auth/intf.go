package auth

import "fyque/model/domain"

type AuthRepository interface {
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}