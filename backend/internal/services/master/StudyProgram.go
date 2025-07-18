package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StudyProgramService struct {
	StudyProgramRepository *repositories.StudyProgramRepository
	FacultyRepository      *repositories.FacultyRepository
	Log                    *zap.SugaredLogger
}

func NewStudyProgramService(studyProgramRepository *repositories.StudyProgramRepository, facultyRepository *repositories.FacultyRepository, log *zap.SugaredLogger) *StudyProgramService {
	return &StudyProgramService{
		StudyProgramRepository: studyProgramRepository,
		FacultyRepository:      facultyRepository,
		Log:                    log,
	}
}

// Core services

func (Service *StudyProgramService) CreateStudyProgramData(ctx *fiber.Ctx, studyProgramRequest *model.StudyProgramRequestModel) (*entities.StudyProgram, int, string) {

	if err := ctx.BodyParser(studyProgramRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current faculty
	facultyEntity := new(entities.Faculty)

	if err := Service.FacultyRepository.GetById(ctx.UserContext(), facultyEntity, studyProgramRequest.Faculty); err != nil {
		return nil, 404, "Faculty with given ID, is not found!"
	}

	studyProgramEntity := &entities.StudyProgram{
		Name:        studyProgramRequest.Name,
		Faculty:     facultyEntity,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := Service.StudyProgramRepository.Create(ctx.UserContext(), studyProgramEntity); err != nil {
		return nil, 500, err.Error()
	}

	return studyProgramEntity, 201, "New study program has been created, successfully!"

}

func (Service *StudyProgramService) GetAllStudyProgramData(ctx *fiber.Ctx) (*[]entities.StudyProgram, int, string) {

	studyProgramEntities := new([]entities.StudyProgram)

	if err := Service.StudyProgramRepository.GetAll(ctx.UserContext(), studyProgramEntities); err != nil {
		return nil, 500, err.Error()
	}

	return studyProgramEntities, 200, "Data of all study program"
}

func (Service *StudyProgramService) GetStudyProgramDataById(ctx *fiber.Ctx) (*entities.StudyProgram, int, string) {

	id := ctx.Params("id", "")
	studyProgramEntity := new(entities.StudyProgram)

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := Service.StudyProgramRepository.GetById(ctx.UserContext(), studyProgramEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return studyProgramEntity, 200, "Study program data with given ID"

}

func (Service *StudyProgramService) UpdateStudyProgramData(ctx *fiber.Ctx, studyProgramRequestModel *model.StudyProgramRequestModel) (*entities.StudyProgram, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(studyProgramRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current study program entity

	currentstudyProgramEntity := new(entities.StudyProgram)

	if err := Service.StudyProgramRepository.GetById(ctx.UserContext(), currentstudyProgramEntity, id); err != nil {
		return nil, 404, "StudyProgram with given ID, is not found!"
	}

	// Get the requested faculty entitiy

	updatedFacultyEntity := new(entities.Faculty)
	if err := Service.FacultyRepository.GetById(ctx.UserContext(), updatedFacultyEntity, studyProgramRequestModel.Faculty); err != nil {
		return nil, 404, "Faculty with given ID, is not found!"
	}

	// Update study program entity

	updatedstudyProgramEntity := &entities.StudyProgram{
		Id:          currentstudyProgramEntity.Id,
		Name:        studyProgramRequestModel.Name,
		Faculty:     updatedFacultyEntity,
		DateCreated: currentstudyProgramEntity.DateCreated,
		DateUpdated: time.Now(),
	}

	// Update study program

	if err := Service.StudyProgramRepository.Update(ctx.UserContext(), updatedstudyProgramEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedstudyProgramEntity, 200, "StudyProgram with given id, has been updated successfully!"
}

func (Service *StudyProgramService) DeleteStudyProgramData(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	studyProgramEntity := new(entities.StudyProgram)

	if err := Service.StudyProgramRepository.GetById(ctx.UserContext(), studyProgramEntity, id); err != nil {
		return 404, "Study Program with given id is not found"
	}

	if err := Service.StudyProgramRepository.Delete(ctx.UserContext(), studyProgramEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
