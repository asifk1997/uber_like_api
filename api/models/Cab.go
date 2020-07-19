package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Cab struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	OwnerName  string    `gorm:"size:255;not null;unique" json:"cab_owner"`
	CarDetails string    `gorm:"size:255;not null;" json:"car_details"`
	LocationId uint64    `gorm:"not null;" json:"location_id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserCarrentLocationAndRange struct {
	CurrentLocation string
	NearbyRange     uint
}

func (u *Cab) FindNearByCabs(db *gorm.DB, getCab UserCarrentLocationAndRange) (*[]Cab, error) {
	var err error
	cabs := []Cab{}
	location := Location{}
	err = db.Debug().Model(&Location{}).Where("location_name = ?", getCab.CurrentLocation).Limit(100).Find(&location).Error
	if err != nil {
		return &[]Cab{}, err
	}
	log.Println(location)
	log.Println(getCab)
	err = db.Debug().Raw("SELECT * from cabs INNER JOIN locations ON locations.id = cabs.location_id AND calculate_distance(locations.latitude,locations.longitude,?,?,'K')<=?", location.Latitude, location.Longitude, getCab.NearbyRange).Limit(100).Scan(&cabs).Error
	if err != nil {
		return &[]Cab{}, err
	}
	log.Println(cabs)
	return &cabs, err
}
