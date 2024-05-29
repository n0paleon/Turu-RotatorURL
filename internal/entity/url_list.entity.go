package entity

import "time"

type UrlList struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   string
	TargetURL string
	HitPerDay int
	HitToday  int `gorm:"default:0"`
	TargetHit int
	TotalHit  int       `gorm:"default:0"`
	Rotate    bool      `gorm:"column:rotate; default:true"`
	CreatedAt time.Time `gorm:"type:timestamp; autoCreateTime; default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `gorm:"type:timestamp; autoUpdateTime; default:CURRENT_TIMESTAMP()"`
}
