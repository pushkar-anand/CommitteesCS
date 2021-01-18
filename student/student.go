package student

import "committees/db/model"

type Student struct {
	model.Model
	Name 	*string `gorm:"name" schema:"name,required" validate:"required,trim,printascii"`
	Usn 	*string `gorm:"usn" schema:"usn,required" validate:"required,trim,printascii"`
	Email 	*string `gorm:"email" schema:"email,required" validate:"required,trim,printascii"`
	Phone 	*string `gorm:"phone" schema:"phone,required" validate:"required,trim,printascii"`
}

func (f Student) TableName() string {
	return "students"
}
