package auth

import (
	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
)

type AuthRepository interface {
	Login(email string, password string) (*dto.UserDto, string, error)
}
