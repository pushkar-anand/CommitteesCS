package faculty

import "committees/db"

type Repository struct {
	db *db.DB
}

func (r *Repository) Create(f *Faculty) error {
	return r.db.Conn.Create(f).Error
}

func (r *Repository) Fetch(id uint) (*Faculty, error) {
	f := &Faculty{}

	err := r.db.Conn.Model(&Faculty{}).Where("id = ?", id).Scan(f).Error
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *Repository) Update(f *Faculty) error {
	return r.db.Conn.Model(&Faculty{}).Where("id = ?", f.ID).Save(f).Error
}

func (r *Repository) Delete(f *Faculty) error {
	return r.db.Conn.Delete(f).Error
}

func (r *Repository) FetchAll() ([]*Faculty, error) {
	faculties := make([]*Faculty, 0)

	err := r.db.Conn.Model(&Faculty{}).Find(&faculties).Error
	if err != nil {
		return nil, err
	}

	return faculties, nil
}
