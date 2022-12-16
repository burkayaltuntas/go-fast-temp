package home

import (
	"github.com/burkayaltuntas/go-fast-temp/pkg/data/models"
	dto "github.com/burkayaltuntas/go-fast-temp/pkg/dto"
)

func toCityDto(model *models.City) *dto.CityDto {

	return &dto.CityDto{
		Id:        model.Id.String(),
		Name:      model.Name,
		PlateCode: model.PlateCode,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
