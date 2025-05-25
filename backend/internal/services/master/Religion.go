package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ReligionService struct {
	ReligionRepository *repositories.ReligionRepository
	Log                *zap.SugaredLogger
}

func NewReligionService(religionRepository *repositories.ReligionRepository, log *zap.SugaredLogger) *ReligionService {
	return &ReligionService{
		ReligionRepository: religionRepository,
		Log:                log,
	}
}

// Core Religion Service

func (service *ReligionService) CreateReligionData(ctx *fiber.Ctx, religionRequest *model.ReligionyModelRequest) (*entities.Religion, int, string) {

	if err := ctx.BodyParser(religionRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	religionEntity := &entities.Religion{
		Name:        religionRequest.Name,
		DateCreated: time.Now(),
	}

	if err := service.ReligionRepository.Create(ctx.UserContext(), religionEntity); err != nil {
		return nil, 500, err.Error()
	}

	return religionEntity, 201, "New Religion data is successfully created!"
}

func (service *ReligionService) GetAllReligionData(ctx *fiber.Ctx) (*[]entities.Religion, int, string) {

	religionEntities := new([]entities.Religion)

	if err := service.ReligionRepository.GetAll(ctx.UserContext(), religionEntities); err != nil {
		return nil, 500, err.Error()
	}

	return religionEntities, 200, "Religions Data"

}

func (service *ReligionService) GetReligionByID(ctx *fiber.Ctx) (*entities.Religion, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	religionEntity := new(entities.Religion)

	if err := service.ReligionRepository.GetByID(ctx.UserContext(), religionEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return religionEntity, 200, "Religion data by ID"

}

func (service *ReligionService) Update(ctx *fiber.Ctx) (*entities.Religion, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	newReligionData := new(entities.Religion)
	updatedReligionData := new(entities.Religion)

	if err := ctx.BodyParser(newReligionData); err != nil {
		return nil, 400, "Invalid Request"
	}

	if err := service.ReligionRepository.Update(ctx.UserContext(), newReligionData, updatedReligionData, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedReligionData, 200, "Religion data with given id, has been updated successfully!"

}

func (service *ReligionService) Delete(ctx *fiber.Ctx) (*entities.Religion, int, string) {
	id := ctx.Params("id", "")
	deletedEntity := new(entities.Religion)

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := service.ReligionRepository.Delete(ctx.UserContext(), deletedEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return deletedEntity, 200, "Religion data with given id, has been deleted successfully"
}
