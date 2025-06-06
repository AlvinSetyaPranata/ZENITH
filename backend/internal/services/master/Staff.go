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

type StaffService struct {
	StaffRepository *repositories.StaffRepository
	Log             *zap.SugaredLogger
}

func NewStaffService(staffRepository *repositories.StaffRepository, log *zap.SugaredLogger) *StaffService {
	return &StaffService{
		StaffRepository: staffRepository,
		Log:             log,
	}
}

// Core repositories

func (Service *StaffService) CreateStaffService(ctx *fiber.Ctx, newStaffRequest *models.StaffRequestModel) (*entities.Staff, int, string) {

	if err := ctx.BodyParser(newStaffRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	staffEntity := &entities.Staff{
		Npwp:           newStaffRequest.Npwp,
		Name:           newStaffRequest.Name,
		GenderId:       newStaffRequest.Gender,
		CityId:         newStaffRequest.City,
		Address:        newStaffRequest.Address,
		ReligionId:     newStaffRequest.Religion,
		ProvinceId:     newStaffRequest.Province,
		FacultyId:      newStaffRequest.Faculty,
		StatusId:       newStaffRequest.Status,
		StudyProgramId: newStaffRequest.StudyProgram,
		UserId:         newStaffRequest.User,
		CountryId:      newStaffRequest.Country,
		DateCreated:    time.Now(),
		DateUpdated:    time.Now(),
	}

	if err := Service.StaffRepository.Create(ctx.UserContext(), staffEntity); err != nil {
		return nil, 500, err.Error()
	}

	npwp := strconv.FormatUint(uint64(staffEntity.Npwp), 10)

	if err := Service.StaffRepository.GetById(ctx.UserContext(), staffEntity, npwp); err != nil {
		return nil, 500, "Oops, something wrong!, try again"
	}

	return staffEntity, 201, "New staff has been registered, successfully!"
}

func (Service *StaffService) GetAllStaffService(ctx *fiber.Ctx) (*[]entities.Staff, int, string) {
	staffsEntity := new([]entities.Staff)

	if err := Service.StaffRepository.GetAll(ctx.UserContext(), staffsEntity); err != nil {
		return nil, 500, err.Error()
	}

	return staffsEntity, 200, "Data of all registered staffr"
}

func (Service *StaffService) GetByIdStaffService(ctx *fiber.Ctx) (*entities.Staff, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Bad Request!"
	}

	staffEntity := new(entities.Staff)

	if err := Service.StaffRepository.GetById(ctx.UserContext(), staffEntity, id); err != nil {

		if errType := utils.HandlerError(err); errType == 404 {

			return nil, 404, "Staff with given id, is not found!"
		}
		return nil, 500, err.Error()

	}

	return staffEntity, 200, "Data of staffr, with given, ID!"
}

func (Service *StaffService) UpdateStaffService(ctx *fiber.Ctx, newStaffRequest *models.StaffRequestModel) (*entities.Staff, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(newStaffRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	currentStaffEntity := new(entities.Staff)

	if err := Service.StaffRepository.GetById(ctx.UserContext(), currentStaffEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Staff with given id, is not found!"
		}
	}

	newStaffEntity := &entities.Staff{
		Npwp:           currentStaffEntity.Npwp,
		Name:           newStaffRequest.Name,
		GenderId:       newStaffRequest.Gender,
		CityId:         newStaffRequest.City,
		Address:        newStaffRequest.Address,
		ReligionId:     newStaffRequest.Religion,
		ProvinceId:     newStaffRequest.Province,
		FacultyId:      newStaffRequest.Faculty,
		StudyProgramId: newStaffRequest.StudyProgram,
		UserId:         newStaffRequest.User,
		CountryId:      newStaffRequest.Country,
		DateCreated:    currentStaffEntity.DateCreated,
		DateUpdated:    time.Now(),
	}

	if err := Service.StaffRepository.Update(ctx.UserContext(), newStaffEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	npwp := strconv.FormatUint(uint64(newStaffEntity.Npwp), 10)

	if err := Service.StaffRepository.GetById(ctx.UserContext(), newStaffEntity, npwp); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Staff with given id, is not found!"
		}
	}

	return newStaffEntity, 200, "Staff with given id, is successfully updated!"
}

func (Service *StaffService) DeleteStaffService(ctx *fiber.Ctx) (int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	currentStaffEntity := new(entities.Staff)

	if err := Service.StaffRepository.GetById(ctx.UserContext(), currentStaffEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return 404, "Staff with given id, is not found!"
		}
	}

	if err := Service.StaffRepository.Delete(ctx.UserContext(), currentStaffEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
