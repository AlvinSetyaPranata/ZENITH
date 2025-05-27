package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProvinceService struct {
	ProvinceRepository *repositories.ProvinceRepository
	Log                *zap.SugaredLogger
}

func NewProvinceService(provinceRepository *repositories.ProvinceRepository, log *zap.SugaredLogger) *ProvinceService {
	return &ProvinceService{
		ProvinceRepository: provinceRepository,
		Log:                log,
	}
}

// Core services

func (Service *ProvinceService) CreateProvinceData(ctx *fiber.Ctx, provinceRequest *model.ProvinceRequestModel) (*entities.Province, int, string) {

	if err := ctx.BodyParser(provinceRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	provinceEntity := &entities.Province{
		Name:        provinceRequest.Name,
		DateCreated: time.Now(),
	}

	if err := Service.ProvinceRepository.Create(ctx.UserContext(), provinceEntity); err != nil {
		return nil, 500, err.Error()
	}

	return provinceEntity, 201, "New province has been created, successfully!"

}

func (Service *ProvinceService) GetAllProvinceData(ctx *fiber.Ctx) (*[]entities.Province, int, string) {

	provinceEntities := new([]entities.Province)

	if err := Service.ProvinceRepository.GetAll(ctx.UserContext(), provinceEntities); err != nil {
		return nil, 500, err.Error()
	}

	return provinceEntities, 200, "Data of all Province"
}

func (Service *ProvinceService) GetProvinceDataById(ctx *fiber.Ctx) (*entities.Province, int, string) {

	id := ctx.Params("id", "")
	provinceEntity := new(entities.Province)

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := Service.ProvinceRepository.GetById(ctx.UserContext(), provinceEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return provinceEntity, 200, "Province data with given ID"

}

func (Service *ProvinceService) UpdateProvinceData(ctx *fiber.Ctx, provinceRequestModel *model.ProvinceRequestModel) (*entities.Province, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(provinceRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	currentprovinceEntity := new(entities.Province)

	if err := Service.ProvinceRepository.GetById(ctx.UserContext(), currentprovinceEntity, id); err != nil {
		return nil, 404, "Province with given ID, is not found!"
	}

	updatedprovinceEntity := &entities.Province{
		Id:          currentprovinceEntity.Id,
		Name:        provinceRequestModel.Name,
		DateCreated: currentprovinceEntity.DateCreated,
	}

	if err := Service.ProvinceRepository.Update(ctx.UserContext(), updatedprovinceEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedprovinceEntity, 200, "Province with given id, has been updated successfully!"
}

func (Service *ProvinceService) DeleteProvinceData(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	provinceEntity := new(entities.Province)

	if err := Service.ProvinceRepository.Delete(ctx.UserContext(), provinceEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
