package master

import (
	"fmt"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
)

type SubjectService struct {
	SubjectRepository *repositories.SubjectRepository
}

func NewSubjectService(subjectRepository *repositories.SubjectRepository) *SubjectService {
	return &SubjectService{
		SubjectRepository: subjectRepository,
	}
}

// Services

func (Service *SubjectService) CreateSubjectService(ctx *fiber.Ctx, subjectRequestModel *models.SubjectRequestModel) (*entities.Subject, int, string) {
	if err := ctx.BodyParser(subjectRequestModel); err != nil {
		fmt.Println(err.Error())
		return nil, 400, "Invalid Request!"
	}

	// Create new subject entity and passed into repository

	newSubjectEntity := &entities.Subject{
		Name:          subjectRequestModel.Name,
		LectureId:     subjectRequestModel.Lecture,
		SubjectTimeId: subjectRequestModel.SubjectTime,
		DateCreated:   time.Now(),
		DateUpdated:   time.Now(),
	}

	if err := Service.SubjectRepository.Create(ctx.UserContext(), newSubjectEntity); err != nil {
		return nil, 500, err.Error()
	}

	return newSubjectEntity, 201, "New subject has been registered successfully!"
}

func (Service *SubjectService) GetAllSubjectService(ctx *fiber.Ctx) (*[]entities.Subject, int, string) {
	subjectEntities := new([]entities.Subject)

	if err := Service.SubjectRepository.GetAll(ctx.UserContext(), subjectEntities); err != nil {
		return nil, 500, err.Error()
	}

	return subjectEntities, 200, "Data of all registered subjects"

}

func (Service *SubjectService) GetSubjectByIdService(ctx *fiber.Ctx) (*entities.Subject, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	subjectEntity := new(entities.Subject)

	if err := Service.SubjectRepository.GetById(ctx.UserContext(), subjectEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return subjectEntity, 200, "Data of subject, with given id!"

}

func (Service *SubjectService) UpdateSubjectService(ctx *fiber.Ctx, subjectRequetModel *models.SubjectRequestModel) (*entities.Subject, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(subjectRequetModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current subject with given id
	currentSubjectEntity := new(entities.Subject)

	if err := Service.SubjectRepository.GetById(ctx.UserContext(), currentSubjectEntity, id); err != nil {
		return nil, 404, "Subject with given id, is not found!"
	}

	newSubjectEntity := &entities.Subject{
		Id:            currentSubjectEntity.Id,
		Name:          subjectRequetModel.Name,
		LectureId:     subjectRequetModel.Lecture,
		SubjectTimeId: subjectRequetModel.SubjectTime,
		DateCreated:   currentSubjectEntity.DateCreated,
		DateUpdated:   time.Now(),
	}

	if err := Service.SubjectRepository.Update(ctx.UserContext(), newSubjectEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return newSubjectEntity, 200, "Subject with given id, has been updated successfully!"
}

func (Service *SubjectService) DeleteSubjectService(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	currentSubject := new(entities.Subject)

	if err := Service.SubjectRepository.GetById(ctx.UserContext(), currentSubject, id); err != nil {
		return 404, "Subject with given id, is not found!"
	}

	if err := Service.SubjectRepository.Delete(ctx.UserContext(), currentSubject, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
