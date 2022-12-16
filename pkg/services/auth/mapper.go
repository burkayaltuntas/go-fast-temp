package auth

import (
	"github.com/burkayaltuntas/go-fast-temp/pkg/data/models"
	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
)

func ToUserDto(u *models.User) *dto.UserDto {
	return &dto.UserDto{
		Id:        u.Id.String(),
		Role:      u.Role,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		IsActive:  u.DeletedAt == nil,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
