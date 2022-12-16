package dto

import "time"

type CityDto struct {
	Id        string
	Name      string
	PlateCode int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
