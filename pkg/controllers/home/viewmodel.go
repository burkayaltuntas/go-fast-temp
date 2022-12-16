package home

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CityVM struct {
	Id        string     `json:"id"`
	Name      string     `json:"name" validate:"required"`
	PlateCode int        `json:"plate_code" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (p *CityVM) Validate() error {
	return validate.Struct(p)
}
