package master

import "gorm.io/gorm"

type RegionUseCase struct {
	DB *gorm.DB
}
