package faculty

import "committees/db/model"

type Faculty struct {
	model.Model
	Name        *string `json:"name"`
	Designation *string `json:"designation"`
}

func (f Faculty) TableName() string {
	return "faculties"
}
