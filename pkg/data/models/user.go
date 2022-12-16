package models

import (
	"time"

	"github.com/burkayaltuntas/go-fast-temp/pkg/common"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID   `gorm:"column:Id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string      `gorm:"column:Name;type:varchar(150);not null"`
	Surname   string      `gorm:"column:Surname;type:varchar(150)"`
	Email     string      `gorm:"column:Email;type:varchar(150)"`
	Password  string      `gorm:"column:Password;type:varchar(150)"`
	Role      common.Role `gorm:"column:Role;type:int;not null"`
	CreatedAt time.Time   `gorm:"column:CreatedAt;type:timestamp;not null"`
	UpdatedAt time.Time   `gorm:"column:UpdatedAt;type:timestamp;not null"`
	DeletedAt *time.Time  `gorm:"column:DeletedAt;type:timestamp"`
}

func (User) TableName() string {
	return "User"
}
