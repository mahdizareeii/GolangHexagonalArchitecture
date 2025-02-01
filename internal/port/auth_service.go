package port

import "go_sample/internal/core"

type AuthService interface {
	Register(email, password, role string) (*core.User, error)
	Login(email, password string) (string, error) //returns jwt token
}
