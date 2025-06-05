package master

import (
	"strconv"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StudentService struct {
	StudentRepository *repositories.StudentRepository
	Log               *zap.SugaredLogger
}

func NewStudentService(studentRepository *repositories.StudentRepository, log *zap.SugaredLogger) *StudentService {
	return &StudentService{
		StudentRepository: studentRepository,
		Log:               log,
	}
}

// Core repositories

func (Service *StudentService) CreateStudentService(ctx *fiber.Ctx, newStudentRequest *models.StudentRequestModel) (*entities.Student, int, string) {

	if err := ctx.BodyParser(newStudentRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	studentEntity := &entities.Student{
		Npm:            newStudentRequest.Npm,
		Name:           newStudentRequest.Name,
		GenderId:       newStudentRequest.Gender,
		CityId:         newStudentRequest.City,
		Address:        newStudentRequest.Address,
		ReligionId:     newStudentRequest.Religion,
		ProvinceId:     newStudentRequest.Province,
		FacultyId:      newStudentRequest.Faculty,
		StatusId:       newStudentRequest.Status,
		StudyProgramId: newStudentRequest.StudyProgram,
		UserId:         newStudentRequest.User,
		CountryId:      newStudentRequest.Country,
		DateCreated:    time.Now(),
		DateUpdated:    time.Now(),
	}

	if err := Service.StudentRepository.Create(ctx.UserContext(), studentEntity); err != nil {
		return nil, 500, err.Error()
	}

	npm := strconv.FormatUint(uint64(studentEntity.Npm), 10)

	if err := Service.StudentRepository.GetById(ctx.UserContext(), studentEntity, npm); err != nil {
		return nil, 500, "Oops, something wrong!, try again"
	}

	return studentEntity, 201, "New student has been registered, successfully!"
}

func (Service *StudentService) GetAllStudentService(ctx *fiber.Ctx) (*[]entities.Student, int, string) {
	studentsEntity := new([]entities.Student)

	if err := Service.StudentRepository.GetAll(ctx.UserContext(), studentsEntity); err != nil {
		return nil, 500, err.Error()
	}

	return studentsEntity, 200, "Data of all registered studentr"
}

func (Service *StudentService) GetByIdStudentService(ctx *fiber.Ctx) (*entities.Student, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Bad Request!"
	}

	studentEntity := new(entities.Student)

	if err := Service.StudentRepository.GetById(ctx.UserContext(), studentEntity, id); err != nil {

		if errType := utils.HandlerError(err); errType == 404 {

			return nil, 404, "Student with given id, is not found!"
		}
		return nil, 500, err.Error()

	}

	return studentEntity, 200, "Data of studentr, with given, ID!"
}

func (Service *StudentService) UpdateStudentService(ctx *fiber.Ctx, newStudentRequest *models.StudentRequestModel) (*entities.Student, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(newStudentRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	currentStudentEntity := new(entities.Student)

	if err := Service.StudentRepository.GetById(ctx.UserContext(), currentStudentEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Student with given id, is not found!"
		}
	}

	newStudentEntity := &entities.Student{
		Npm:            currentStudentEntity.Npm,
		Name:           newStudentRequest.Name,
		GenderId:       newStudentRequest.Gender,
		CityId:         newStudentRequest.City,
		Address:        newStudentRequest.Address,
		ReligionId:     newStudentRequest.Religion,
		ProvinceId:     newStudentRequest.Province,
		FacultyId:      newStudentRequest.Faculty,
		StudyProgramId: newStudentRequest.StudyProgram,
		UserId:         newStudentRequest.User,
		CountryId:      newStudentRequest.Country,
		DateCreated:    currentStudentEntity.DateCreated,
		DateUpdated:    time.Now(),
	}

	if err := Service.StudentRepository.Update(ctx.UserContext(), newStudentEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	npm := strconv.FormatUint(uint64(newStudentEntity.Npm), 10)

	if err := Service.StudentRepository.GetById(ctx.UserContext(), newStudentEntity, npm); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Student with given id, is not found!"
		}
	}

	return newStudentEntity, 200, "Student with given id, is successfully updated!"
}

func (Service *StudentService) DeleteStudentService(ctx *fiber.Ctx) (int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	currentStudentEntity := new(entities.Student)

	if err := Service.StudentRepository.GetById(ctx.UserContext(), currentStudentEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return 404, "Student with given id, is not found!"
		}
	}

	if err := Service.StudentRepository.Delete(ctx.UserContext(), currentStudentEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
