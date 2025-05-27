package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FacultyService struct {
	Repository *repositories.FacultyRepository
	Log        *zap.SugaredLogger
}

func NewFacultyService(repository *repositories.FacultyRepository, log *zap.SugaredLogger) *FacultyService {
	return &FacultyService{
		Repository: repository,
		Log:        log,
	}
}

// Core services

func (Service *FacultyService) CreateFacultyData(ctx *fiber.Ctx, facultyRequestModel *model.FacultyRequestModel) (*entities.Faculty, int, string) {

	if err := ctx.BodyParser(facultyRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	facultyEntity := &entities.Faculty{
		Name:        facultyRequestModel.Name,
		DateCreated: time.Now(),
	}

	if err := Service.Repository.Create(ctx.UserContext(), facultyEntity); err != nil {
		return nil, 500, err.Error()
	}

	return facultyEntity, 201, "New Faculty data, has been created successfully!"
}

func (Service *FacultyService) GetAllFaculties(ctx *fiber.Ctx) (*[]entities.Faculty, int, string) {
	facultyEntities := new([]entities.Faculty)

	if err := Service.Repository.GetAll(ctx.UserContext(), facultyEntities); err != nil {
		return nil, 500, err.Error()
	}

	return facultyEntities, 200, "Faculty data retrieved successfully!"
}

func (Service *FacultyService) GetFacultyById(ctx *fiber.Ctx) (*entities.Faculty, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	facultyEntity := new(entities.Faculty)

	if err := Service.Repository.GetById(ctx.UserContext(), facultyEntity, id); err != nil {
		return nil, 404, "Faculty not found!"
	}

	return facultyEntity, 200, "Faculty data retrieved successfully!"
}

func (Service *FacultyService) UpdateFacultyData(ctx *fiber.Ctx, facultyRequestModel *model.FacultyRequestModel) (*entities.Faculty, int, string) {

	facultyId := ctx.Params("id", "")

	if facultyId == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(facultyRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current faculty entity

	currentFacultyEntity := new(entities.Faculty)

	if err := Service.Repository.GetById(ctx.UserContext(), currentFacultyEntity, facultyId); err != nil {
		return nil, 404, "Faculty with given id is not found!"
	}

	facultyEntity := &entities.Faculty{
		Id:          currentFacultyEntity.Id,
		Name:        facultyRequestModel.Name,
		DateCreated: time.Now(),
	}

	if err := Service.Repository.Update(ctx.UserContext(), facultyEntity); err != nil {
		return nil, 500, err.Error()
	}

	return facultyEntity, 200, "Faculty data updated successfully!"
}

func (Service *FacultyService) DeleteFacultyData(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	currentFacultyEntity := new(entities.Faculty)

	if err := Service.Repository.GetById(ctx.UserContext(), currentFacultyEntity, id); err != nil {
		return 404, "Faculty with given id is not found!"
	}

	if err := Service.Repository.Delete(ctx.UserContext(), currentFacultyEntity, id); err != nil {
		return 500, "Failed to delete faculty data!"
	}

	return 204, ""
}
