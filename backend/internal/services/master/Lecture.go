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

type LectureService struct {
	LectureRepository *repositories.LectureRepository
	Log               *zap.SugaredLogger
}

func NewLectureService(lectureRepository *repositories.LectureRepository, log *zap.SugaredLogger) *LectureService {
	return &LectureService{
		LectureRepository: lectureRepository,
		Log:               log,
	}
}

// Core repositories

func (Service *LectureService) CreateLectureService(ctx *fiber.Ctx, newLectureRequest *models.LectureRequestModel) (*entities.Lecture, int, string) {

	if err := ctx.BodyParser(newLectureRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	lectureEntity := &entities.Lecture{
		Nidn:           newLectureRequest.Nidn,
		Name:           newLectureRequest.Name,
		GenderId:       newLectureRequest.Gender,
		CityId:         newLectureRequest.City,
		Address:        newLectureRequest.Address,
		ReligionId:     newLectureRequest.Religion,
		ProvinceId:     newLectureRequest.Province,
		FacultyId:      newLectureRequest.Faculty,
		StatusId:       newLectureRequest.Status,
		StudyProgramId: newLectureRequest.StudyProgram,
		UserId:         newLectureRequest.User,
		CountryId:      newLectureRequest.Country,
		DateCreated:    time.Now(),
		DateUpdated:    time.Now(),
	}

	if err := Service.LectureRepository.Create(ctx.UserContext(), lectureEntity); err != nil {
		return nil, 500, err.Error()
	}

	nidn := strconv.FormatUint(uint64(lectureEntity.Nidn), 10)

	if err := Service.LectureRepository.GetById(ctx.UserContext(), lectureEntity, nidn); err != nil {
		return nil, 500, "Oops, something wrong!, try again"
	}

	return lectureEntity, 201, "New lecture has been registered, successfully!"
}

func (Service *LectureService) GetAllLectureService(ctx *fiber.Ctx) (*[]entities.Lecture, int, string) {
	lecturesEntity := new([]entities.Lecture)

	if err := Service.LectureRepository.GetAll(ctx.UserContext(), lecturesEntity); err != nil {
		return nil, 500, err.Error()
	}

	return lecturesEntity, 200, "Data of all registered lecturer"
}

func (Service *LectureService) GetByIdLectureService(ctx *fiber.Ctx) (*entities.Lecture, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Bad Request!"
	}

	lectureEntity := new(entities.Lecture)

	if err := Service.LectureRepository.GetById(ctx.UserContext(), lectureEntity, id); err != nil {

		if errType := utils.HandlerError(err); errType == 404 {

			return nil, 404, "Lecture with given id, is not found!"
		}
		return nil, 500, err.Error()

	}

	return lectureEntity, 200, "Data of lecturer, with given, ID!"
}

func (Service *LectureService) UpdateLectureService(ctx *fiber.Ctx, newLectureRequest *models.LectureRequestModel) (*entities.Lecture, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(newLectureRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	currentLectureEntity := new(entities.Lecture)

	if err := Service.LectureRepository.GetById(ctx.UserContext(), currentLectureEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Lecture with given id, is not found!"
		}
	}

	newLectureEntity := &entities.Lecture{
		Nidn:           currentLectureEntity.Nidn,
		Name:           newLectureRequest.Name,
		GenderId:       newLectureRequest.Gender,
		CityId:         newLectureRequest.City,
		Address:        newLectureRequest.Address,
		ReligionId:     newLectureRequest.Religion,
		ProvinceId:     newLectureRequest.Province,
		FacultyId:      newLectureRequest.Faculty,
		StudyProgramId: newLectureRequest.StudyProgram,
		UserId:         newLectureRequest.User,
		CountryId:      newLectureRequest.Country,
		DateCreated:    currentLectureEntity.DateCreated,
		DateUpdated:    time.Now(),
	}

	if err := Service.LectureRepository.Update(ctx.UserContext(), newLectureEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	nidn := strconv.FormatUint(uint64(newLectureEntity.Nidn), 10)

	if err := Service.LectureRepository.GetById(ctx.UserContext(), newLectureEntity, nidn); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Lecture with given id, is not found!"
		}
	}

	return newLectureEntity, 200, "Lecture with given id, is successfully updated!"
}

func (Service *LectureService) DeleteLectureService(ctx *fiber.Ctx) (int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	currentLectureEntity := new(entities.Lecture)

	if err := Service.LectureRepository.GetById(ctx.UserContext(), currentLectureEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return 404, "Lecture with given id, is not found!"
		}
	}

	if err := Service.LectureRepository.Delete(ctx.UserContext(), currentLectureEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
