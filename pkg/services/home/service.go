package home

import (
	"github.com/burkayaltuntas/go-fast-temp/pkg/data/models"
	"github.com/burkayaltuntas/go-fast-temp/pkg/dto"
	"gorm.io/gorm"
)

var _db *gorm.DB

type HomeService struct {
}

func NewHomeService(db *gorm.DB) *HomeService {
	_db = db
	return &HomeService{}
}

func (s *HomeService) GetCityList() []dto.CityDto {
	cities := []models.City{}
	dtoList := []dto.CityDto{}

	_db.Find(&cities)

	for _, v := range cities {
		dtoList = append(dtoList, *toCityDto(&v))
	}

	return dtoList
}
