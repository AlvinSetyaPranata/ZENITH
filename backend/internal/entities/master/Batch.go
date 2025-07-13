package entities

import "time"

type Batch struct {
	Id          uint      `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;unique"`
	StartPeriod time.Time `gorm:"column:start_period;uniqueIndex:idx_batch_periods"`
	EndPeriod   time.Time `gorm:"column:end_period;uniqueIndex:idx_batch_periods"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
	DateUpdated time.Time `gorm:"column:date_updated"`
}
