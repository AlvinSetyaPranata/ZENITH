package master

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FacultyPresenter struct {
	FacultyService *services.FacultyService
	Log            *zap.SugaredLogger
}

func NewFacultyPresenter(facultyService *services.FacultyService, log *zap.SugaredLogger) *FacultyPresenter {
	return &FacultyPresenter{
		FacultyService: facultyService,
		Log:            log,
	}
}

// Core Presenters

func (Presenter *FacultyPresenter) CreateFacultyPresenter(ctx *fiber.Ctx) (*models.FacultyResponseModel, int, string) {

	facultyRequest := new(models.FacultyRequestModel)

	newFacultyEntity, status, messege := Presenter.FacultyService.CreateFacultyData(ctx, facultyRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.FacultyResponseModel{
		Id:          newFacultyEntity.Id,
		Name:        newFacultyEntity.Name,
		DateCreated: newFacultyEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *FacultyPresenter) GetAllFacultiesData(ctx *fiber.Ctx) (*[]models.FacultyResponseModel, int, string) {

	facultiesEntity, status, messege := Presenter.FacultyService.GetAllFaculties(ctx)

	response := make([]models.FacultyResponseModel, 0, len(*facultiesEntity))

	for _, faculty := range *facultiesEntity {
		response = append(response, models.FacultyResponseModel{
			Id:          faculty.Id,
			Name:        faculty.Name,
			DateCreated: faculty.DateCreated,
		})
	}

	return &response, status, messege
}

func (Presenter *FacultyPresenter) GetFacultyDataById(ctx *fiber.Ctx) (*models.FacultyResponseModel, int, string) {

	facultyEntity, status, messege := Presenter.FacultyService.GetFacultyById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.FacultyResponseModel{
		Id:          facultyEntity.Id,
		Name:        facultyEntity.Name,
		DateCreated: facultyEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *FacultyPresenter) UpdateFacultyData(ctx *fiber.Ctx) (*models.FacultyResponseModel, int, string) {

	facultyRequestModel := new(models.FacultyRequestModel)

	updatedFacultyData, status, messege := Presenter.FacultyService.UpdateFacultyData(ctx, facultyRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.FacultyResponseModel{
		Id:          updatedFacultyData.Id,
		Name:        updatedFacultyData.Name,
		DateCreated: updatedFacultyData.DateCreated,
	}

	return response, status, messege
}

func (Presenter *FacultyPresenter) DeleteFacultyData(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.FacultyService.DeleteFacultyData(ctx)

	if status != 204 {
		return status, messege
	}

	return status, ""
}
