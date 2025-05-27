package master

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProvincePresenter struct {
	ProvinceService *services.ProvinceService
	Log             *zap.SugaredLogger
}

func NewProvincePresenter(ProvinceService *services.ProvinceService, log *zap.SugaredLogger) *ProvincePresenter {
	return &ProvincePresenter{
		ProvinceService: ProvinceService,
		Log:             log,
	}
}

// Core Presenters

func (Presenter *ProvincePresenter) CreateProvincePresenter(ctx *fiber.Ctx) (*models.ProvinceResponseModel, int, string) {

	ProvinceRequest := new(models.ProvinceRequestModel)

	newProvinceEntity, status, messege := Presenter.ProvinceService.CreateProvinceData(ctx, ProvinceRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.ProvinceResponseModel{
		Id:          newProvinceEntity.Id,
		Name:        newProvinceEntity.Name,
		DateCreated: newProvinceEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *ProvincePresenter) GetAllProvincesData(ctx *fiber.Ctx) (*[]models.ProvinceResponseModel, int, string) {

	ProvincesEntity, status, messege := Presenter.ProvinceService.GetAllProvinceData(ctx)

	response := make([]models.ProvinceResponseModel, 0, len(*ProvincesEntity))

	for _, Province := range *ProvincesEntity {
		response = append(response, models.ProvinceResponseModel{
			Id:          Province.Id,
			Name:        Province.Name,
			DateCreated: Province.DateCreated,
		})
	}

	return &response, status, messege
}

func (Presenter *ProvincePresenter) GetProvinceDataById(ctx *fiber.Ctx) (*models.ProvinceResponseModel, int, string) {

	ProvinceEntity, status, messege := Presenter.ProvinceService.GetProvinceDataById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.ProvinceResponseModel{
		Id:          ProvinceEntity.Id,
		Name:        ProvinceEntity.Name,
		DateCreated: ProvinceEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *ProvincePresenter) UpdateProvinceData(ctx *fiber.Ctx) (*models.ProvinceResponseModel, int, string) {

	ProvinceRequestModel := new(models.ProvinceRequestModel)

	updatedProvinceData, status, messege := Presenter.ProvinceService.UpdateProvinceData(ctx, ProvinceRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.ProvinceResponseModel{
		Id:          updatedProvinceData.Id,
		Name:        updatedProvinceData.Name,
		DateCreated: updatedProvinceData.DateCreated,
	}

	return response, status, messege
}

func (Presenter *ProvincePresenter) DeleteProvinceData(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.ProvinceService.DeleteProvinceData(ctx)

	if status != 204 {
		return status, messege
	}

	return status, ""
}
