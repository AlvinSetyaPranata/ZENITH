package master

import (
	"fmt"

	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
)

type SubjectPresenter struct {
	SubjectService *services.SubjectService
}

func NewSubjectPresenter(subjectService *services.SubjectService) *SubjectPresenter {
	return &SubjectPresenter{
		SubjectService: subjectService,
	}
}

// Core presenters

func (Presenter *SubjectPresenter) CreateSubjectPresenter(ctx *fiber.Ctx) (*models.SubjectResponseModel, int, string) {

	subjectRequest := new(models.SubjectRequestModel)

	subjectEntity, status, messege := Presenter.SubjectService.CreateSubjectService(ctx, subjectRequest)

	if status != 201 {
		return nil, status, messege
	}

	fmt.Println(subjectEntity)

	response := &models.SubjectResponseModel{
		Id:      subjectEntity.Id,
		Name:    subjectEntity.Name,
		Lecture: subjectEntity.LectureId,
	}

	return response, status, messege
}

func (Presenter *SubjectPresenter) GetAllSubjectPresenter(ctx *fiber.Ctx) (*[]models.SubjectResponseModel, int, string) {

	subjectEntities, status, messege := Presenter.SubjectService.GetAllSubjectService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.SubjectResponseModel, 0, len(*subjectEntities))

	for _, subject := range *subjectEntities {
		response = append(response, models.SubjectResponseModel{
			Id:      subject.Id,
			Name:    subject.Name,
			Lecture: subject.LectureId,
		})
	}

	return &response, status, messege
}

func (Presenter *SubjectPresenter) GetSubjectByIdPresenter(ctx *fiber.Ctx) (*models.SubjectResponseModel, int, string) {

	subjectEntity, status, messege := Presenter.SubjectService.GetSubjectByIdService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.SubjectResponseModel{
		Id:      subjectEntity.Id,
		Name:    subjectEntity.Name,
		Lecture: subjectEntity.LectureId,
	}

	return response, status, messege
}
func (Presenter *SubjectPresenter) UpdateSubjectPresenter(ctx *fiber.Ctx) (*models.SubjectResponseModel, int, string) {

	subjectRequest := new(models.SubjectRequestModel)

	subjectEntity, status, messege := Presenter.SubjectService.UpdateSubjectService(ctx, subjectRequest)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.SubjectResponseModel{
		Id:      subjectEntity.Id,
		Name:    subjectEntity.Name,
		Lecture: subjectEntity.LectureId,
	}

	return response, status, messege
}

func (Presenter *SubjectPresenter) DeleteSubjectPresenter(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.SubjectService.DeleteSubjectService(ctx)

	if status != 200 {
		return status, messege
	}

	return status, ""
}
