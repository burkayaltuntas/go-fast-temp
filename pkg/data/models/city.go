package models

import (
	"time"

	"github.com/google/uuid"
)

type City struct {
	Id        uuid.UUID  `gorm:"column:Id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string     `gorm:"column:Name;type:varchar(150);not null"`
	PlateCode int        `gorm:"column:PlateCode;type:int;not null"`
	CreatedAt time.Time  `gorm:"column:CreatedAt;type:timestamp;not null"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt;type:timestamp;not null"`
	DeletedAt *time.Time `gorm:"column:DeletedAt;type:timestamp"`
}

func (City) TableName() string {
	return "City"
}
