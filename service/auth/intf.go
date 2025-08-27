package auth

import "fyque/model/web/auth"

type AuthService interface {
	Register(req auth.RegisterRequest) (string, error)
	Login(req auth.LoginRequest)(string, error)
}