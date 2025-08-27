package auth

import (
	"errors"
	"fyque/helper"
	"fyque/model/domain"
	m "fyque/model/web/auth"
	"fyque/repository/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo auth.AuthRepository
}

func NewAuthService(authRepo auth.AuthRepository) AuthService{
	return &AuthServiceImpl{
		authRepo: authRepo,
	}
}

func (s *AuthServiceImpl) Register(req m.RegisterRequest) (string, error){
	user, err := s.authRepo.FindByUsername(req.Username)

	if err == nil && user != nil  {
		return "", errors.New("username exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user = &domain.User {
		Username: req.Username,
		Password: string(hashed),
	}
	
	if err := s.authRepo.Create(user); err != nil {
		return "", err
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthServiceImpl) Login(req m.LoginRequest) (string, error){
	user, err := s.authRepo.FindByUsername(req.Username)

	if err != nil && user == nil {
		return "", errors.New("invalid credentials")
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}