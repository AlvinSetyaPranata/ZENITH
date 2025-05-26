package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GenderService struct {
	GenderRepository *repositories.GenderRepository
	Log              *zap.SugaredLogger
}

func NewGenderService(genderRepository *repositories.GenderRepository, log *zap.SugaredLogger) *GenderService {
	return &GenderService{
		GenderRepository: genderRepository,
		Log:              log,
	}
}

// Core services

func (Service *GenderService) CreateGender(ctx *fiber.Ctx, genderRequest *model.GenderRequestModel) (*entities.Gender, int, string) {

	if err := ctx.BodyParser(genderRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	genderEntity := &entities.Gender{
		Name:        genderRequest.Name,
		DateCreated: time.Now(),
	}

	if err := Service.GenderRepository.Create(ctx.UserContext(), genderEntity); err != nil {
		return nil, 500, err.Error()
	}

	return genderEntity, 201, "New gender has been created, successfully!"

}

func (Service *GenderService) GetAllGenderData(ctx *fiber.Ctx) (*[]entities.Gender, int, string) {

	genderEntities := new([]entities.Gender)

	if err := Service.GenderRepository.GetAll(ctx.UserContext(), genderEntities); err != nil {
		return nil, 500, err.Error()
	}

	return genderEntities, 200, "Data of all gender"
}

func (Service *GenderService) GetGenderDataById(ctx *fiber.Ctx) (*entities.Gender, int, string) {

	id := ctx.Params("id", "")
	genderEntity := new(entities.Gender)

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := Service.GenderRepository.GetById(ctx.UserContext(), genderEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return genderEntity, 200, "Gender data with given ID"

}

func (Service *GenderService) UpdateGenderData(ctx *fiber.Ctx, genderRequestModel *model.GenderRequestModel) (*entities.Gender, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(genderRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	currentGenderEntity := new(entities.Gender)

	if err := Service.GenderRepository.GetById(ctx.UserContext(), currentGenderEntity, id); err != nil {
		return nil, 404, "Gender with given ID, is not found!"
	}

	updatedGenderEntity := &entities.Gender{
		Id:          currentGenderEntity.Id,
		Name:        genderRequestModel.Name,
		DateCreated: currentGenderEntity.DateCreated,
	}

	if err := Service.GenderRepository.Update(ctx.UserContext(), updatedGenderEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedGenderEntity, 200, "Gender with given id, has been updated successfully!"
}

func (Service *GenderService) DeleteGenderData(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	genderEntity := new(entities.Gender)

	if err := Service.GenderRepository.Delete(ctx.UserContext(), genderEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
