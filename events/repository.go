package events

import "committees/db"

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(f *Event) error {
	return r.db.Conn.Create(f).Error
}

func (r *Repository) Fetch(id uint) (*Event, error) {
	f := &Event{}

	err := r.db.Conn.Model(&Event{}).Where("id = ?", id).Scan(f).Error
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *Repository) Update(f *Event) error {
	return r.db.Conn.Model(&Event{}).Where("id = ?", f.ID).Save(f).Error
}

func (r *Repository) Delete(f *Event) error {
	return r.db.Conn.Delete(f).Error
}

func (r *Repository) FetchAll() ([]*Event, error) {
	Events := make([]*Event, 0)

	err := r.db.Conn.Model(&Event{}).Find(&Events).Error
	if err != nil {
		return nil, err
	}

	return Events, nil
}
