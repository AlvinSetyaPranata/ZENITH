package master

import (
	"fmt"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SubjectTimeService struct {
	SubjectTimeRepository *repositories.SubjectTimeRepository
	Log                   *zap.SugaredLogger
}

func NewSubjectTimeService(subjectTimeRepository *repositories.SubjectTimeRepository, log *zap.SugaredLogger) *SubjectTimeService {
	return &SubjectTimeService{
		SubjectTimeRepository: subjectTimeRepository,
		Log:                   log,
	}
}

// Core Services

func (Service *SubjectTimeService) CreateNewSubjectTime(ctx *fiber.Ctx, subjectTimeRequestModel *models.SubjectTimeRequestModel) (*entities.SubjectTime, int, string) {
	if err := ctx.BodyParser(subjectTimeRequestModel); err != nil {
		fmt.Println(err.Error())
		return nil, 400, "Invalid Request"
	}

	subjectTimeEntity := &entities.SubjectTime{
		Name:        subjectTimeRequestModel.Name,
		TimeStart:   subjectTimeRequestModel.TimeStart,
		TimeEnd:     subjectTimeRequestModel.TimeEnd,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := Service.SubjectTimeRepository.Create(ctx.UserContext(), subjectTimeEntity); err != nil {

		if errType := utils.HandlerError(err); errType == 409 {
			return nil, 409, "SubjectTime with given name is already exists"
		}

		return nil, 500, err.Error()
	}

	return subjectTimeEntity, 201, "New subjectTime has been added successfully!"
}

func (Service *SubjectTimeService) GetAllSubjectTime(ctx *fiber.Ctx) (*[]entities.SubjectTime, int, string) {

	subjectTimeEntities := new([]entities.SubjectTime)

	if err := Service.SubjectTimeRepository.GetAll(ctx.UserContext(), subjectTimeEntities); err != nil {
		return nil, 500, err.Error()
	}

	return subjectTimeEntities, 200, "Data of all subjectTime"
}

func (Service *SubjectTimeService) GetSubjectTimeById(ctx *fiber.Ctx) (*entities.SubjectTime, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	subjectTimeEntitiy := new(entities.SubjectTime)

	if err := Service.SubjectTimeRepository.GetById(ctx.UserContext(), subjectTimeEntitiy, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "SubjectTime with given id, is not found"
		}

		return nil, 500, err.Error()
	}

	return subjectTimeEntitiy, 200, "Data of subjectTime with given id"
}

func (Service *SubjectTimeService) UpdateSubjectTime(ctx *fiber.Ctx, subjectTimeRequestModel *models.SubjectTimeRequestModel) (*entities.SubjectTime, int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(subjectTimeRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current subjectTime entitiy
	currentSubjectTime := new(entities.SubjectTime)

	if err := Service.SubjectTimeRepository.GetById(ctx.UserContext(), currentSubjectTime, id); err != nil {
		return nil, 404, "SubjectTime with given id, is not found!"
	}

	newSubjectTimeEntity := &entities.SubjectTime{
		Id:          currentSubjectTime.Id,
		Name:        subjectTimeRequestModel.Name,
		TimeStart:   subjectTimeRequestModel.TimeStart,
		TimeEnd:     subjectTimeRequestModel.TimeEnd,
		DateCreated: currentSubjectTime.DateCreated,
		DateUpdated: time.Now(),
	}

	if err := Service.SubjectTimeRepository.Update(ctx.UserContext(), newSubjectTimeEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return newSubjectTimeEntity, 200, "SubjectTime with given id, has been updated successfully!"
}

func (Service *SubjectTimeService) DeleteSubjectTime(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	deletedSubjectTimeEntity := new(entities.SubjectTime)

	// Search for subjectTime with given id

	if err := Service.SubjectTimeRepository.GetById(ctx.UserContext(), deletedSubjectTimeEntity, id); err != nil {
		return 404, "SubjectTime with given id, is not found!"
	}

	if err := Service.SubjectTimeRepository.Delete(ctx.UserContext(), deletedSubjectTimeEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
