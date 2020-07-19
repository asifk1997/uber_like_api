package models

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Booking struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID        uint32    `gorm:"not null" json:"user_id"`
	SourceId      uint64    `gorm:"not null" json:"source_id"`
	DestinationId uint64    `gorm:"not null" json:"destination_id"`
	CabId         uint64    `gorm:"not null" json:"cab_id"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type BookingUserEntry struct {
	Id              uint32
	SourceName      string
	DestinationName string
	CabOwnerName    string
	CreatedAt       time.Time
}

func (p *BookingUserEntry) Validate() error {

	if p.SourceName == "" {
		return errors.New("Source Required")
	}
	if p.DestinationName == "" {
		return errors.New("Destination Required")
	}

	return nil
}

func (p *BookingUserEntry) SaveBooking(db *gorm.DB, uid uint32) (*Booking, error) {
	var err error
	b := &Booking{}
	sourceLocation := &Location{}
	destinationLocation := &Location{}
	cab := &Cab{}
	err = db.Debug().Model(&Location{}).Where("location_name = ?", p.SourceName).Find(&sourceLocation).Error
	if err != nil {
		return &Booking{}, err
	}
	b.SourceId = sourceLocation.ID
	err = db.Debug().Model(&Location{}).Where("location_name = ?", p.DestinationName).Find(&destinationLocation).Error
	if err != nil {
		return &Booking{}, err
	}
	b.DestinationId = destinationLocation.ID
	err = db.Debug().Model(&Location{}).Where("owner_name = ?", p.CabOwnerName).Find(&cab).Error
	if err != nil {
		return &Booking{}, err
	}
	b.CabId = cab.ID
	b.UserID = uid
	err = db.Debug().Model(&Booking{}).Create(&b).Error
	if err != nil {
		return &Booking{}, err
	}
	return b, nil
}

func (p *Booking) FindAllBookingsForAUser(db *gorm.DB, uid uint32) (*[]BookingUserEntry, error) {
	var err error
	Bookings := []BookingUserEntry{}
	err = db.Debug().Raw("SELECT bookings.id, l1.location_name as source_name ,l2.location_name as destination_name ,cabs.owner_name as cab_owner_name ,bookings.created_at as created_at from bookings JOIN locations as l1 ON l1.id = bookings.source_id JOIN locations as l2 ON l2.id = bookings.destination_id JOIN cabs ON cabs.id=bookings.cab_id AND user_id=? ORDER BY bookings.id DESC", uid).Scan(&Bookings).Error
	if err != nil {
		return &[]BookingUserEntry{}, err
	}
	log.Println(Bookings)
	return &Bookings, nil
}

func (p *Booking) FindBookingByID(db *gorm.DB, pid uint64) (*Booking, error) {
	var err error
	err = db.Debug().Model(&Booking{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Booking{}, err
	}
	return p, nil
}
