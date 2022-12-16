package home

import "github.com/burkayaltuntas/go-fast-temp/pkg/dto"

type HomeRepository interface {
	GetCityList() []dto.CityDto
}
