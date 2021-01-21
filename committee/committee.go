package committee

import (
	"committees/db/model"
	"committees/faculty"
	"time"
)

type Committee struct {
	model.Model
	Name         *string            `gorm:"name" schema:"name,required" validate:"required,trim,printascii" json:"name"`
	Description  *string            `gorm:"description" schema:"description,required" validate:"required,trim,printascii" json:"description"`
	CreationDate time.Time          `gorm:"creation_date" schema:"creation_date,required" validate:"required,trim,printascii" json:"creation_date"`
	Members      []*faculty.Faculty `gorm:"many2many:committee_members;" schema:"members,required" json:"members"`
}

type Members struct {
	model.Model
	CommitteeID uint `gorm:"committee_id"`
	FacultyID   uint `gorm:"faculty_id"`
}

func (m Members) TableName() string {
	return "committee_members"
}
