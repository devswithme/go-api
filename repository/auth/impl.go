package auth

import (
	"fyque/model/domain"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (r *AuthRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (r *AuthRepositoryImpl) Create(user *domain.User) error {
	return r.db.Create(user).Error
}


