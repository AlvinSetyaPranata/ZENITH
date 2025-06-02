package master

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SubjectTimePresenter struct {
	SubjectTimeService *services.SubjectTimeService
	Log                *zap.SugaredLogger
}

func NewSubjectTimePresenter(subjectTimeService *services.SubjectTimeService, log *zap.SugaredLogger) *SubjectTimePresenter {
	return &SubjectTimePresenter{
		SubjectTimeService: subjectTimeService,
		Log:                log,
	}
}

// Core Presenters

func (Presenter *SubjectTimePresenter) CreateSubjectTimePresenter(ctx *fiber.Ctx) (*models.SubjectTimeResponseModel, int, string) {
	subjectTimeRequestModel := new(models.SubjectTimeRequestModel)

	subjectTimeEntity, status, messege := Presenter.SubjectTimeService.CreateNewSubjectTime(ctx, subjectTimeRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.SubjectTimeResponseModel{
		Id:          subjectTimeEntity.Id,
		Name:        subjectTimeEntity.Name,
		TimeStart:   subjectTimeEntity.TimeStart,
		TimeEnd:     subjectTimeEntity.TimeEnd,
		DateCreated: subjectTimeEntity.DateCreated,
		DateUpdated: subjectTimeEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *SubjectTimePresenter) GetAllSubjectTimePresenter(ctx *fiber.Ctx) (*[]models.SubjectTimeResponseModel, int, string) {
	subjectTimeEntities, status, messege := Presenter.SubjectTimeService.GetAllSubjectTime(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.SubjectTimeResponseModel, 0, len(*subjectTimeEntities))

	for _, subjectTime := range *subjectTimeEntities {
		response = append(response, models.SubjectTimeResponseModel{
			Id:          subjectTime.Id,
			Name:        subjectTime.Name,
			TimeStart:   subjectTime.TimeStart,
			TimeEnd:     subjectTime.TimeEnd,
			DateCreated: subjectTime.DateCreated,
			DateUpdated: subjectTime.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *SubjectTimePresenter) GetSubjectTimeById(ctx *fiber.Ctx) (*models.SubjectTimeResponseModel, int, string) {
	subjectTimeEntity, status, messege := Presenter.SubjectTimeService.GetSubjectTimeById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.SubjectTimeResponseModel{
		Id:          subjectTimeEntity.Id,
		Name:        subjectTimeEntity.Name,
		TimeStart:   subjectTimeEntity.TimeStart,
		TimeEnd:     subjectTimeEntity.TimeEnd,
		DateCreated: subjectTimeEntity.DateCreated,
		DateUpdated: subjectTimeEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *SubjectTimePresenter) UpdateSubjectTime(ctx *fiber.Ctx) (*models.SubjectTimeResponseModel, int, string) {

	subjectTimeRequestModel := new(models.SubjectTimeRequestModel)

	subjectTimeEntity, status, messege := Presenter.SubjectTimeService.UpdateSubjectTime(ctx, subjectTimeRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.SubjectTimeResponseModel{
		Id:          subjectTimeEntity.Id,
		Name:        subjectTimeEntity.Name,
		TimeStart:   subjectTimeEntity.TimeStart,
		TimeEnd:     subjectTimeEntity.TimeEnd,
		DateCreated: subjectTimeEntity.DateCreated,
		DateUpdated: subjectTimeEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *SubjectTimePresenter) DeleteSubjectTime(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.SubjectTimeService.DeleteSubjectTime(ctx)

	if status != 200 {
		return status, messege
	}

	return status, ""
}
