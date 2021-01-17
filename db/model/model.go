package model

import "time"

type Model struct {
	ID        uint      `gorm:"id" schema:"-"`
	CreatedAt time.Time `gorm:"created_at" schema:"-"`
	UpdatedAt time.Time `gorm:"updated_at" schema:"-"`
}
