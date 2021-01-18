package student

import "committees/db"

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(f *Student) error {
	return r.db.Conn.Create(f).Error
}

func (r *Repository) Fetch(id uint) (*Student, error) {
	f := &Student{}

	err := r.db.Conn.Model(&Student{}).Where("id = ?", id).Scan(f).Error
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *Repository) Update(f *Student) error {
	return r.db.Conn.Model(&Student{}).Where("id = ?", f.ID).Save(f).Error
}

func (r *Repository) Delete(f *Student) error {
	return r.db.Conn.Delete(f).Error
}

func (r *Repository) FetchAll() ([]*Student, error) {
	students := make([]*Student, 0)

	err := r.db.Conn.Model(&Student{}).Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}
