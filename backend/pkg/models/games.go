package models

import (
	"database/sql"
	"time"
)

// WIP
type Game struct {
	ID          string `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Board       string
	RedUser     User `gorm:"foreignkey:RedUserID"`
	RedUserID   uint
	BlackUser   User `gorm:"foreignkey:BlackUserID"`
	BlackUserID uint
	Resigned    sql.NullInt64
	TimeControl int64
}
