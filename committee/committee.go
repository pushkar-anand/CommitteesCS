package committee

import (
	"committees/db/model"
	"committees/faculty"
)

type Committee struct {
	model.Model
	Name    *string            `gorm:"name" schema:"name,required" validate:"required,trim,printascii" json:"name"`
	Members []*faculty.Faculty `gorm:"many2many:committee_members;" schema:"members,required" json:"members"`
}

type Members struct {
	model.Model
	CommitteeID uint `gorm:"committee_id"`
	FacultyID   uint `gorm:"faculty_id"`
}

func (m Members) TableName() string {
	return "committee_members"
}
