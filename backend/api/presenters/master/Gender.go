package master

import (
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GenderPresenter struct {
	GenderService *services.GenderService
	Log           *zap.SugaredLogger
}

func NewGenderPresenter(genderService *services.GenderService, log *zap.SugaredLogger) *GenderPresenter {
	return &GenderPresenter{
		GenderService: genderService,
		Log:           log,
	}
}

// Core Presenters

func (Presenter *GenderPresenter) CreateGenderData(ctx *fiber.Ctx) (*model.GenderResponseModel, int, string) {

	genderRequestModel := new(model.GenderRequestModel)

	genderEntity, status, messege := Presenter.GenderService.CreateGender(ctx, genderRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &model.GenderResponseModel{
		Id:          genderEntity.Id,
		Name:        genderEntity.Name,
		DateCreated: genderEntity.DateCreated,
	}

	return response, status, messege

}

func (Presenter *GenderPresenter) GetAllGenderData(ctx *fiber.Ctx) (*[]model.GenderResponseModel, int, string) {
	genderEntities, status, messege := Presenter.GenderService.GetAllGenderData(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]model.GenderResponseModel, 0, len(*genderEntities))

	for _, genderEntity := range *genderEntities {
		response = append(response, model.GenderResponseModel{
			Id:          genderEntity.Id,
			Name:        genderEntity.Name,
			DateCreated: genderEntity.DateCreated,
		})
	}

	return &response, status, messege
}

func (Presenter *GenderPresenter) GetGenderDataById(ctx *fiber.Ctx) (*model.GenderResponseModel, int, string) {
	genderEntity, status, messege := Presenter.GenderService.GetGenderDataById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &model.GenderResponseModel{
		Id:          genderEntity.Id,
		Name:        genderEntity.Name,
		DateCreated: genderEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *GenderPresenter) UpdateGenderData(ctx *fiber.Ctx) (*model.GenderResponseModel, int, string) {
	genderRequestModel := new(model.GenderRequestModel)

	genderEntity, status, messege := Presenter.GenderService.UpdateGenderData(ctx, genderRequestModel)
	if status != 200 {
		return nil, status, messege
	}
	response := &model.GenderResponseModel{
		Id:          genderEntity.Id,
		Name:        genderEntity.Name,
		DateCreated: genderEntity.DateCreated,
	}
	return response, status, messege
}

func (Presenter *GenderPresenter) DeleteGenderData(ctx *fiber.Ctx) (int, string) {
	status, messege := Presenter.GenderService.DeleteGenderData(ctx)
	if status != 200 {
		return status, messege
	}
	return status, messege
}
