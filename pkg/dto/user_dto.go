package dto

import (
	"time"

	"github.com/burkayaltuntas/go-fast-temp/pkg/common"
)

type UserDto struct {
	Id           string      `json:"id"`
	Role         common.Role `json:"role"`
	RoleName     string      `json:"roleName"`
	Email        string      `json:"email"`
	Password     string      `json:"-"`
	Name         string      `json:"name"`
	IsLead       bool        `json:"isLead"`
	ContractorId string      `json:"contractorId"`
	Surname      string      `json:"surname"`
	IsActive     bool        `json:"isActive"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
