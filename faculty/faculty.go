package faculty

import "committees/db/model"

type Faculty struct {
	model.Model
	Name        *string `gorm:"name" schema:"name,required" validate:"required,trim,printascii" json:"name"`
	Designation *string `gorm:"designation" schema:"designation,required" validate:"required,trim,printascii" json:"designation"`
}

func (f Faculty) TableName() string {
	return "faculties"
}
