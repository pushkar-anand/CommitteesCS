package model

import (
	"database/sql/driver"
	"time"
)

type Model struct {
	ID        uint      `gorm:"id" schema:"-" json:"id"`
	CreatedAt time.Time `gorm:"created_at" schema:"-"`
	UpdatedAt time.Time `gorm:"updated_at" schema:"-"`
}

type Date struct {
	time.Time
}

const dateFormat = "2006-01-02"

func (d *Date) Value() (driver.Value, error) {
	return d.MarshalText()
}

func (d *Date) Scan(src interface{}) error {
	t, err := time.Parse(dateFormat, src.(string))
	if err != nil {
		return err
	}

	*d = Date{t}

	return nil
}

