package master

import (
	authModels "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
)

type StaffPresenter struct {
	StaffService *services.StaffService
}

func NewStaffPresenter(staffService *services.StaffService) *StaffPresenter {
	return &StaffPresenter{
		StaffService: staffService,
	}
}

// Core presenter

func (Presenter *StaffPresenter) CreateStaffPresenter(ctx *fiber.Ctx) (*models.StaffResponseModel, int, string) {

	staffRequetModel := new(models.StaffRequestModel)

	staffEntity, status, messege := Presenter.StaffService.CreateStaffService(ctx, staffRequetModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.StaffResponseModel{
		Npwp:    staffEntity.Npwp,
		Name:    staffEntity.Name,
		Address: staffEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          staffEntity.Gender.Id,
			Name:        staffEntity.Gender.Name,
			DateCreated: staffEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          staffEntity.City.Id,
			Name:        staffEntity.City.Name,
			DateCreated: staffEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          staffEntity.Religion.Id,
			Name:        staffEntity.Religion.Name,
			DateCreated: staffEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          staffEntity.Country.Id,
			Name:        staffEntity.Country.Name,
			DateCreated: staffEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          staffEntity.Province.Id,
			Name:        staffEntity.Province.Name,
			DateCreated: staffEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          staffEntity.Status.Id,
			Name:        staffEntity.Status.Name,
			DateCreated: staffEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          staffEntity.User.Id,
			RoleId:      uint(staffEntity.User.RoleId),
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},
		DateCreated: staffEntity.DateCreated,
		DateUpdated: staffEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StaffPresenter) GetAllStaffPresenter(ctx *fiber.Ctx) (*[]models.StaffResponseModel, int, string) {
	staffEntities, status, messege := Presenter.StaffService.GetAllStaffService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.StaffResponseModel, 0, len(*staffEntities))

	for _, staff := range *staffEntities {
		response = append(response, models.StaffResponseModel{
			Npwp:    staff.Npwp,
			Name:    staff.Name,
			Address: staff.Address,
			Gender: models.GenderResponseModel{
				Id:          staff.Gender.Id,
				Name:        staff.Gender.Name,
				DateCreated: staff.Gender.DateCreated,
			},
			City: models.CityModelResponse{
				Id:          staff.City.Id,
				Name:        staff.City.Name,
				DateCreated: staff.City.DateCreated,
			},
			Religion: models.ReligionModelResponse{
				Id:          staff.Religion.Id,
				Name:        staff.Religion.Name,
				DateCreated: staff.Religion.DateCreated,
			},
			Country: models.CountryModelResponse{
				Id:          staff.Country.Id,
				Name:        staff.Country.Name,
				DateCreated: staff.Country.DateCreated,
			},
			Province: models.ProvinceResponseModel{
				Id:          staff.Province.Id,
				Name:        staff.Province.Name,
				DateCreated: staff.Province.DateCreated,
			},
			Status: models.StatusModelResponse{
				Id:          staff.Status.Id,
				Name:        staff.Status.Name,
				DateCreated: staff.Status.DateCreated,
			},
			User: authModels.UserResponseModel{
				Id:          staff.User.Id,
				RoleId:      uint(staff.User.RoleId),
				DateCreated: staff.DateCreated,
				DateUpdated: staff.DateUpdated,
			},

			Faculty: models.FacultyResponseModel{
				Id:          staff.Faculty.Id,
				Name:        staff.Faculty.Name,
				DateCreated: staff.DateCreated,
			},
			StudyProgram: models.StudyProgramResponseModel{
				Id:          staff.Faculty.Id,
				Name:        staff.Faculty.Name,
				DateCreated: staff.DateCreated,
				DateUpdated: staff.DateUpdated,
			},
			DateCreated: staff.DateCreated,
			DateUpdated: staff.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *StaffPresenter) GetStaffByIdPresenter(ctx *fiber.Ctx) (*models.StaffResponseModel, int, string) {

	staffEntity, status, messege := Presenter.StaffService.GetByIdStaffService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StaffResponseModel{
		Npwp:    staffEntity.Npwp,
		Name:    staffEntity.Name,
		Address: staffEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          staffEntity.Gender.Id,
			Name:        staffEntity.Gender.Name,
			DateCreated: staffEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          staffEntity.City.Id,
			Name:        staffEntity.City.Name,
			DateCreated: staffEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          staffEntity.Religion.Id,
			Name:        staffEntity.Religion.Name,
			DateCreated: staffEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          staffEntity.Country.Id,
			Name:        staffEntity.Country.Name,
			DateCreated: staffEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          staffEntity.Province.Id,
			Name:        staffEntity.Province.Name,
			DateCreated: staffEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          staffEntity.Status.Id,
			Name:        staffEntity.Status.Name,
			DateCreated: staffEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          staffEntity.User.Id,
			RoleId:      uint(staffEntity.User.RoleId),
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},
		DateCreated: staffEntity.DateCreated,
		DateUpdated: staffEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StaffPresenter) UpdateStaffPresenter(ctx *fiber.Ctx) (*models.StaffResponseModel, int, string) {

	staffRequestModel := new(models.StaffRequestModel)

	staffEntity, status, messege := Presenter.StaffService.UpdateStaffService(ctx, staffRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StaffResponseModel{
		Npwp:    staffEntity.Npwp,
		Name:    staffEntity.Name,
		Address: staffEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          staffEntity.Gender.Id,
			Name:        staffEntity.Gender.Name,
			DateCreated: staffEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          staffEntity.City.Id,
			Name:        staffEntity.City.Name,
			DateCreated: staffEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          staffEntity.Religion.Id,
			Name:        staffEntity.Religion.Name,
			DateCreated: staffEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          staffEntity.Country.Id,
			Name:        staffEntity.Country.Name,
			DateCreated: staffEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          staffEntity.Province.Id,
			Name:        staffEntity.Province.Name,
			DateCreated: staffEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          staffEntity.Status.Id,
			Name:        staffEntity.Status.Name,
			DateCreated: staffEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          staffEntity.User.Id,
			RoleId:      uint(staffEntity.User.RoleId),
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          staffEntity.Faculty.Id,
			Name:        staffEntity.Faculty.Name,
			DateCreated: staffEntity.DateCreated,
			DateUpdated: staffEntity.DateUpdated,
		},
		DateCreated: staffEntity.DateCreated,
		DateUpdated: staffEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StaffPresenter) DeleteStaffPresenter(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.StaffService.DeleteStaffService(ctx)

	if status != 200 {
		return status, messege
	}

	return status, messege
}
