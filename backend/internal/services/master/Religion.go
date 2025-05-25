package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ReligionService struct {
	ReligionRepository *repositories.ReligionRepository
	Log                *zap.SugaredLogger
}

func NewReligionService(religionRepository *repositories.ReligionRepository, log *zap.SugaredLogger) *ReligionService {
	return &ReligionService{
		ReligionRepository: religionRepository,
		Log:                log,
	}
}

// Core Religion Service

func (service *ReligionService) CreateReligionData(ctx *fiber.Ctx, religionRequest *model.ReligionyModelRequest) (*entities.Religion, int, string) {

	if err := ctx.BodyParser(religionRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	religionEntity := &entities.Religion{
		Name:        religionRequest.Name,
		DateCreated: time.Now(),
	}

	if err := service.ReligionRepository.Create(ctx.UserContext(), religionEntity); err != nil {
		return nil, 500, err.Error()
	}

	return religionEntity, 201, "New Religion data is successfully created!"
}
