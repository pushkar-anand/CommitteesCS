package events

import (
	"committees/db/model"
	"time"
)

type Event struct {
	model.Model
	Name             *string   `gorm:"name" schema:"name,required" validate:"required,trim,printascii"`
	StartDate        time.Time `gorm:"start_date" schema:"start_date,required" validate:"required,trim,printascii"`
	EndDate          time.Time `gorm:"end_date" schema:"end_date,required" validate:"required,trim,printascii"`
	TotalExpenditure *string   `gorm:"total_expenditure" schema:"total_expenditure,required" validate:"required,trim,printascii"`
}

func (f Event) TableName() string {
	return "events"
}
