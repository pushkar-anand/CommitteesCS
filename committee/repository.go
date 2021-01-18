package committee

import (
	"committees/db"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(c *Committee) error {
	return r.db.Conn.Create(c).Error
}

func (r *Repository) Fetch(id uint) (*Committee, error) {
	f := &Committee{}

	err := r.db.Conn.Model(&Committee{}).Preload("Members").Where("id = ?", id).Scan(f).Error
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *Repository) Update(c *Committee) error {
	return r.db.Conn.Model(&Committee{}).Where("id = ?", c.ID).Save(c).Error
}

func (r *Repository) Delete(c *Committee) error {
	return r.db.Conn.Delete(c).Error
}

func (r *Repository) FetchAll() ([]*Committee, error) {
	faculties := make([]*Committee, 0)

	err := r.db.Conn.Model(&Committee{}).Preload("Members").Find(&faculties).Error
	if err != nil {
		return nil, err
	}

	return faculties, nil
}

func (r *Repository) AddFacultyToCommittee(committeeID, facultyID uint) error {
	m := &Members{
		CommitteeID: committeeID,
		FacultyID:   facultyID,
	}

	return r.db.Conn.Create(m).Error
}
