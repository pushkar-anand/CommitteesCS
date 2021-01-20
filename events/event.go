package events

import (
	"committees/db/model"
)

type Event struct {
	model.Model
	Name      			*string `gorm:"name" schema:"name,required" validate:"required,trim,printascii"`
	StartDate 			model.Date `gorm:"start_date" schema:"start_date,required" validate:"required,trim,printascii"`
	EndDate     		model.Date `gorm:"end_date" schema:"end_date,required" validate:"required,trim,printascii"`
	TotalExpenditure    *string `gorm:"total_expenditure" schema:"total_expenditure,required" validate:"required,trim,printascii"`
}

func (f Event) TableName() string {
	return "events"
}
