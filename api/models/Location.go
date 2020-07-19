package models

type Location struct {
	ID           uint64  `gorm:"primary_key;auto_increment" json:"id"`
	LocationName string  `gorm:"size:255;not null;unique" json:"location_name"`
	Latitude     float64 `gorm:"not null;" json:"latitude"`
	Longitude    float64 `gorm:"not null;" json:"longitude"`
}
