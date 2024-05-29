package entity

import "time"

type RotatedUrl struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	RotateID  string
	TargetURL string
	Hit       int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"type:timestamp; autoCreateTime; default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"type:timestamp; autoUpdateTime; default:CURRENT_TIMESTAMP()"`
}
