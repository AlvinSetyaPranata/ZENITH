package master

import (
	authModels "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
)

type LecturePresenter struct {
	LectureService *services.LectureService
}

func NewLecturePresenter(lectureService *services.LectureService) *LecturePresenter {
	return &LecturePresenter{
		LectureService: lectureService,
	}
}

// Core presenter

func (Presenter *LecturePresenter) CreateLecturePresenter(ctx *fiber.Ctx) (*models.LectureResponseModel, int, string) {

	lectureRequetModel := new(models.LectureRequestModel)

	lectureEntity, status, messege := Presenter.LectureService.CreateLectureService(ctx, lectureRequetModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.LectureResponseModel{
		Nidn:    lectureEntity.Nidn,
		Name:    lectureEntity.Name,
		Address: lectureEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          lectureEntity.Gender.Id,
			Name:        lectureEntity.Gender.Name,
			DateCreated: lectureEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          lectureEntity.City.Id,
			Name:        lectureEntity.City.Name,
			DateCreated: lectureEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          lectureEntity.Religion.Id,
			Name:        lectureEntity.Religion.Name,
			DateCreated: lectureEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          lectureEntity.Country.Id,
			Name:        lectureEntity.Country.Name,
			DateCreated: lectureEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          lectureEntity.Province.Id,
			Name:        lectureEntity.Province.Name,
			DateCreated: lectureEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          lectureEntity.Status.Id,
			Name:        lectureEntity.Status.Name,
			DateCreated: lectureEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          lectureEntity.User.Id,
			RoleId:      uint(lectureEntity.User.RoleId),
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},
		DateCreated: lectureEntity.DateCreated,
		DateUpdated: lectureEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *LecturePresenter) GetAllLecturePresenter(ctx *fiber.Ctx) (*[]models.LectureResponseModel, int, string) {
	lectureEntities, status, messege := Presenter.LectureService.GetAllLectureService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.LectureResponseModel, 0, len(*lectureEntities))

	for _, lecture := range *lectureEntities {
		response = append(response, models.LectureResponseModel{
			Nidn:    lecture.Nidn,
			Name:    lecture.Name,
			Address: lecture.Address,
			Gender: models.GenderResponseModel{
				Id:          lecture.Gender.Id,
				Name:        lecture.Gender.Name,
				DateCreated: lecture.Gender.DateCreated,
			},
			City: models.CityModelResponse{
				Id:          lecture.City.Id,
				Name:        lecture.City.Name,
				DateCreated: lecture.City.DateCreated,
			},
			Religion: models.ReligionModelResponse{
				Id:          lecture.Religion.Id,
				Name:        lecture.Religion.Name,
				DateCreated: lecture.Religion.DateCreated,
			},
			Country: models.CountryModelResponse{
				Id:          lecture.Country.Id,
				Name:        lecture.Country.Name,
				DateCreated: lecture.Country.DateCreated,
			},
			Province: models.ProvinceResponseModel{
				Id:          lecture.Province.Id,
				Name:        lecture.Province.Name,
				DateCreated: lecture.Province.DateCreated,
			},
			Status: models.StatusModelResponse{
				Id:          lecture.Status.Id,
				Name:        lecture.Status.Name,
				DateCreated: lecture.Status.DateCreated,
			},
			User: authModels.UserResponseModel{
				Id:          lecture.User.Id,
				RoleId:      uint(lecture.User.RoleId),
				DateCreated: lecture.DateCreated,
				DateUpdated: lecture.DateUpdated,
			},

			Faculty: models.FacultyResponseModel{
				Id:          lecture.Faculty.Id,
				Name:        lecture.Faculty.Name,
				DateCreated: lecture.DateCreated,
			},
			StudyProgram: models.StudyProgramResponseModel{
				Id:          lecture.Faculty.Id,
				Name:        lecture.Faculty.Name,
				DateCreated: lecture.DateCreated,
				DateUpdated: lecture.DateUpdated,
			},
			DateCreated: lecture.DateCreated,
			DateUpdated: lecture.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *LecturePresenter) GetLectureByIdPresenter(ctx *fiber.Ctx) (*models.LectureResponseModel, int, string) {

	lectureEntity, status, messege := Presenter.LectureService.GetByIdLectureService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.LectureResponseModel{
		Nidn:    lectureEntity.Nidn,
		Name:    lectureEntity.Name,
		Address: lectureEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          lectureEntity.Gender.Id,
			Name:        lectureEntity.Gender.Name,
			DateCreated: lectureEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          lectureEntity.City.Id,
			Name:        lectureEntity.City.Name,
			DateCreated: lectureEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          lectureEntity.Religion.Id,
			Name:        lectureEntity.Religion.Name,
			DateCreated: lectureEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          lectureEntity.Country.Id,
			Name:        lectureEntity.Country.Name,
			DateCreated: lectureEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          lectureEntity.Province.Id,
			Name:        lectureEntity.Province.Name,
			DateCreated: lectureEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          lectureEntity.Status.Id,
			Name:        lectureEntity.Status.Name,
			DateCreated: lectureEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          lectureEntity.User.Id,
			RoleId:      uint(lectureEntity.User.RoleId),
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},
		DateCreated: lectureEntity.DateCreated,
		DateUpdated: lectureEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *LecturePresenter) UpdateLecturePresenter(ctx *fiber.Ctx) (*models.LectureResponseModel, int, string) {

	lectureRequestModel := new(models.LectureRequestModel)

	lectureEntity, status, messege := Presenter.LectureService.UpdateLectureService(ctx, lectureRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.LectureResponseModel{
		Nidn:    lectureEntity.Nidn,
		Name:    lectureEntity.Name,
		Address: lectureEntity.Address,
		Gender: models.GenderResponseModel{
			Id:          lectureEntity.Gender.Id,
			Name:        lectureEntity.Gender.Name,
			DateCreated: lectureEntity.Gender.DateCreated,
		},
		City: models.CityModelResponse{
			Id:          lectureEntity.City.Id,
			Name:        lectureEntity.City.Name,
			DateCreated: lectureEntity.City.DateCreated,
		},
		Religion: models.ReligionModelResponse{
			Id:          lectureEntity.Religion.Id,
			Name:        lectureEntity.Religion.Name,
			DateCreated: lectureEntity.Religion.DateCreated,
		},
		Country: models.CountryModelResponse{
			Id:          lectureEntity.Country.Id,
			Name:        lectureEntity.Country.Name,
			DateCreated: lectureEntity.Country.DateCreated,
		},
		Province: models.ProvinceResponseModel{
			Id:          lectureEntity.Province.Id,
			Name:        lectureEntity.Province.Name,
			DateCreated: lectureEntity.Province.DateCreated,
		},
		Status: models.StatusModelResponse{
			Id:          lectureEntity.Status.Id,
			Name:        lectureEntity.Status.Name,
			DateCreated: lectureEntity.Status.DateCreated,
		},
		User: authModels.UserResponseModel{
			Id:          lectureEntity.User.Id,
			RoleId:      uint(lectureEntity.User.RoleId),
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},

		Faculty: models.FacultyResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
		},
		StudyProgram: models.StudyProgramResponseModel{
			Id:          lectureEntity.Faculty.Id,
			Name:        lectureEntity.Faculty.Name,
			DateCreated: lectureEntity.DateCreated,
			DateUpdated: lectureEntity.DateUpdated,
		},
		DateCreated: lectureEntity.DateCreated,
		DateUpdated: lectureEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *LecturePresenter) DeleteLecturePresenter(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.LectureService.DeleteLectureService(ctx)

	if status != 200 {
		return status, messege
	}

	return status, messege
}
