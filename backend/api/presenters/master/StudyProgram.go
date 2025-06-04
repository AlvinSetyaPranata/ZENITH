package master

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StudyProgramPresenter struct {
	StudyProgramService *services.StudyProgramService
	Log                 *zap.SugaredLogger
}

func NewStudyProgramPresenter(StudyProgramService *services.StudyProgramService, log *zap.SugaredLogger) *StudyProgramPresenter {
	return &StudyProgramPresenter{
		StudyProgramService: StudyProgramService,
		Log:                 log,
	}
}

// Core Presenters

func (Presenter *StudyProgramPresenter) CreateStudyProgramPresenter(ctx *fiber.Ctx) (*models.StudyProgramResponseModel, int, string) {

	StudyProgramRequest := new(models.StudyProgramRequestModel)

	newStudyProgramEntity, status, messege := Presenter.StudyProgramService.CreateStudyProgramData(ctx, StudyProgramRequest)

	if status != 201 {
		return nil, status, messege
	}

	facultyResponse := &models.FacultyResponseModel{
		Id:          newStudyProgramEntity.Faculty.Id,
		Name:        newStudyProgramEntity.Faculty.Name,
		DateCreated: newStudyProgramEntity.Faculty.DateCreated,
	}

	response := &models.StudyProgramResponseModel{
		Id:          newStudyProgramEntity.Id,
		Name:        newStudyProgramEntity.Name,
		Faculty:     facultyResponse,
		DateCreated: newStudyProgramEntity.DateCreated,
		DateUpdated: newStudyProgramEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudyProgramPresenter) GetAllStudyProgramsData(ctx *fiber.Ctx) (*[]models.StudyProgramResponseModel, int, string) {

	studyProgramsEntity, status, messege := Presenter.StudyProgramService.GetAllStudyProgramData(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.StudyProgramResponseModel, 0, len(*studyProgramsEntity))

	for _, StudyProgram := range *studyProgramsEntity {

		response = append(response, models.StudyProgramResponseModel{
			Id:   StudyProgram.Id,
			Name: StudyProgram.Name,
			Faculty: &models.FacultyResponseModel{
				Id:          StudyProgram.Faculty.Id,
				Name:        StudyProgram.Faculty.Name,
				DateCreated: StudyProgram.Faculty.DateCreated,
			},
			DateCreated: StudyProgram.DateCreated,
			DateUpdated: StudyProgram.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *StudyProgramPresenter) GetStudyProgramDataById(ctx *fiber.Ctx) (*models.StudyProgramResponseModel, int, string) {

	StudyProgramEntity, status, messege := Presenter.StudyProgramService.GetStudyProgramDataById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StudyProgramResponseModel{
		Id:   StudyProgramEntity.Id,
		Name: StudyProgramEntity.Name,
		Faculty: &models.FacultyResponseModel{
			Id:          StudyProgramEntity.Faculty.Id,
			Name:        StudyProgramEntity.Faculty.Name,
			DateCreated: StudyProgramEntity.Faculty.DateCreated,
		},
		DateCreated: StudyProgramEntity.DateCreated,
		DateUpdated: StudyProgramEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudyProgramPresenter) UpdateStudyProgramData(ctx *fiber.Ctx) (*models.StudyProgramResponseModel, int, string) {

	StudyProgramRequestModel := new(models.StudyProgramRequestModel)

	updatedStudyProgramData, status, messege := Presenter.StudyProgramService.UpdateStudyProgramData(ctx, StudyProgramRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StudyProgramResponseModel{
		Id:   updatedStudyProgramData.Id,
		Name: updatedStudyProgramData.Name,
		Faculty: &models.FacultyResponseModel{
			Id:          updatedStudyProgramData.Faculty.Id,
			Name:        updatedStudyProgramData.Faculty.Name,
			DateCreated: updatedStudyProgramData.Faculty.DateCreated,
		},
		DateCreated: updatedStudyProgramData.DateCreated,
		DateUpdated: updatedStudyProgramData.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudyProgramPresenter) DeleteStudyProgramData(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.StudyProgramService.DeleteStudyProgramData(ctx)

	if status != 204 {
		return status, messege
	}

	return status, ""
}
