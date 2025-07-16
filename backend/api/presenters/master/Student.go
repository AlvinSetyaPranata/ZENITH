package master

import (
	"fmt"

	authModels "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
)

type StudentPresenter struct {
	StudentService *services.StudentService
}

func NewStudentPresenter(studentService *services.StudentService) *StudentPresenter {
	return &StudentPresenter{
		StudentService: studentService,
	}
}

// Core presenter

func (Presenter *StudentPresenter) CreateStudentPresenter(ctx *fiber.Ctx) (*models.StudentResponseModel, int, string) {

	studentRequetModel := new(models.StudentRequestModel)

	studentEntity, status, messege := Presenter.StudentService.CreateStudentService(ctx, studentRequetModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.StudentResponseModel{
		Npm:     studentEntity.Npm,
		Name:    studentEntity.Name,
		Address: studentEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          studentEntity.Gender.Id,
			Name:        studentEntity.Gender.Name,
			DateCreated: studentEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          studentEntity.City.Id,
			Name:        studentEntity.City.Name,
			DateCreated: studentEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          studentEntity.Religion.Id,
			Name:        studentEntity.Religion.Name,
			DateCreated: studentEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          studentEntity.Country.Id,
			Name:        studentEntity.Country.Name,
			DateCreated: studentEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          studentEntity.Province.Id,
			Name:        studentEntity.Province.Name,
			DateCreated: studentEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          studentEntity.Status.Id,
			Name:        studentEntity.Status.Name,
			DateCreated: studentEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          studentEntity.User.Id,
			RoleId:      uint(studentEntity.User.RoleId),
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},
		DateCreated: studentEntity.DateCreated,
		DateUpdated: studentEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudentPresenter) GetAllStudentPresenter(ctx *fiber.Ctx) (*[]models.StudentResponseModel, int, string) {
	studentEntities, status, messege := Presenter.StudentService.GetAllStudentService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.StudentResponseModel, 0, len(*studentEntities))

	for _, student := range *studentEntities {
		response = append(response, models.StudentResponseModel{
			Npm:     student.Npm,
			Name:    student.Name,
			Address: student.Address,
			Gender: models.GenderResponseModel{
				Id:          student.Gender.Id,
				Name:        student.Gender.Name,
				DateCreated: student.Gender.DateCreated,
			},
			City: models.CityModelResponse{
				Id:          student.City.Id,
				Name:        student.City.Name,
				DateCreated: student.City.DateCreated,
			},
			Religion: models.ReligionModelResponse{
				Id:          student.Religion.Id,
				Name:        student.Religion.Name,
				DateCreated: student.Religion.DateCreated,
			},
			Country: models.CountryModelResponse{
				Id:          student.Country.Id,
				Name:        student.Country.Name,
				DateCreated: student.Country.DateCreated,
			},
			Province: models.ProvinceResponseModel{
				Id:          student.Province.Id,
				Name:        student.Province.Name,
				DateCreated: student.Province.DateCreated,
			},
			Status: models.StatusModelResponse{
				Id:          student.Status.Id,
				Name:        student.Status.Name,
				DateCreated: student.Status.DateCreated,
			},
			User: authModels.UserResponseModel{
				Id:          student.User.Id,
				Email:       student.User.Email,
				RoleId:      uint(student.User.RoleId),
				DateCreated: student.DateCreated,
				DateUpdated: student.DateUpdated,
			},

			Faculty: models.FacultyResponseModel{
				Id:          student.Faculty.Id,
				Name:        student.Faculty.Name,
				DateCreated: student.DateCreated,
			},
			StudyProgram: models.StudyProgramResponseModel{
				Id:          student.Faculty.Id,
				Name:        student.Faculty.Name,
				DateCreated: student.DateCreated,
				DateUpdated: student.DateUpdated,
			},
			DateCreated: student.DateCreated,
			DateUpdated: student.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *StudentPresenter) GetStudentByIdPresenter(ctx *fiber.Ctx) (*models.StudentResponseModel, int, string) {

	fmt.Println("Goes here")

	studentEntity, status, messege := Presenter.StudentService.GetByIdStudentService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StudentResponseModel{
		Npm:     studentEntity.Npm,
		Name:    studentEntity.Name,
		Address: studentEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          studentEntity.Gender.Id,
			Name:        studentEntity.Gender.Name,
			DateCreated: studentEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          studentEntity.City.Id,
			Name:        studentEntity.City.Name,
			DateCreated: studentEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          studentEntity.Religion.Id,
			Name:        studentEntity.Religion.Name,
			DateCreated: studentEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          studentEntity.Country.Id,
			Name:        studentEntity.Country.Name,
			DateCreated: studentEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          studentEntity.Province.Id,
			Name:        studentEntity.Province.Name,
			DateCreated: studentEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          studentEntity.Status.Id,
			Name:        studentEntity.Status.Name,
			DateCreated: studentEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          studentEntity.User.Id,
			RoleId:      uint(studentEntity.User.RoleId),
			RoleName:    studentEntity.User.Role.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},
		DateCreated: studentEntity.DateCreated,
		DateUpdated: studentEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudentPresenter) GetStudentByUserIdPresenter(ctx *fiber.Ctx) (*models.StudentResponseModel, int, string) {

	studentEntity, status, messege := Presenter.StudentService.GetStudentByUserIdService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StudentResponseModel{
		Npm:     studentEntity.Npm,
		Name:    studentEntity.Name,
		Address: studentEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          studentEntity.Gender.Id,
			Name:        studentEntity.Gender.Name,
			DateCreated: studentEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          studentEntity.City.Id,
			Name:        studentEntity.City.Name,
			DateCreated: studentEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          studentEntity.Religion.Id,
			Name:        studentEntity.Religion.Name,
			DateCreated: studentEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          studentEntity.Country.Id,
			Name:        studentEntity.Country.Name,
			DateCreated: studentEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          studentEntity.Province.Id,
			Name:        studentEntity.Province.Name,
			DateCreated: studentEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          studentEntity.Status.Id,
			Name:        studentEntity.Status.Name,
			DateCreated: studentEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          studentEntity.User.Id,
			RoleId:      uint(studentEntity.User.RoleId),
			RoleName:    studentEntity.User.Role.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},
		DateCreated: studentEntity.DateCreated,
		DateUpdated: studentEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudentPresenter) UpdateStudentPresenter(ctx *fiber.Ctx) (*models.StudentResponseModel, int, string) {

	studentRequestModel := new(models.StudentRequestModel)

	studentEntity, status, messege := Presenter.StudentService.UpdateStudentService(ctx, studentRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.StudentResponseModel{
		Npm:     studentEntity.Npm,
		Name:    studentEntity.Name,
		Address: studentEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          studentEntity.Gender.Id,
			Name:        studentEntity.Gender.Name,
			DateCreated: studentEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          studentEntity.City.Id,
			Name:        studentEntity.City.Name,
			DateCreated: studentEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          studentEntity.Religion.Id,
			Name:        studentEntity.Religion.Name,
			DateCreated: studentEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          studentEntity.Country.Id,
			Name:        studentEntity.Country.Name,
			DateCreated: studentEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          studentEntity.Province.Id,
			Name:        studentEntity.Province.Name,
			DateCreated: studentEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          studentEntity.Status.Id,
			Name:        studentEntity.Status.Name,
			DateCreated: studentEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          studentEntity.User.Id,
			RoleId:      uint(studentEntity.User.RoleId),
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          studentEntity.Faculty.Id,
			Name:        studentEntity.Faculty.Name,
			DateCreated: studentEntity.DateCreated,
			DateUpdated: studentEntity.DateUpdated,
		},
		DateCreated: studentEntity.DateCreated,
		DateUpdated: studentEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *StudentPresenter) DeleteStudentPresenter(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.StudentService.DeleteStudentService(ctx)

	if status != 200 {
		return status, messege
	}

	return status, messege
}
