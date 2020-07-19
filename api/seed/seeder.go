package seed

import (
	"log"

	"github.com/asifk1997/uber_like_api/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Varun",
		Email:    "varun@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Rahul",
		Email:    "rahul@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Tanmay",
		Email:    "tanmay@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Abhishek",
		Email:    "abhishek@gmail.com",
		Password: "password",
	},
}

var locations = []models.Location{
	models.Location{
		LocationName: "Indore",
		Latitude:     77.2,
		Longitude:    99.2,
	},
	models.Location{
		LocationName: "Delhi",
		Latitude:     77.5,
		Longitude:    99.33,
	},
	models.Location{
		LocationName: "Mumbai",
		Latitude:     77.2,
		Longitude:    99.32,
	},
}

var cabs = []models.Cab{
	models.Cab{
		OwnerName:  "Aakash",
		CarDetails: "Toyota Innova",
		LocationId: 2,
	},
	models.Cab{
		OwnerName:  "Revant",
		CarDetails: "Maruti Swift",
		LocationId: 3,
	},
	models.Cab{
		OwnerName:  "Naveed",
		CarDetails: "Honda City",
		LocationId: 1,
	},
}

var bookings = []models.Booking{
	models.Booking{
		UserID:        1,
		SourceId:      1,
		DestinationId: 2,
		CabId:         1,
	},
	models.Booking{
		UserID:        1,
		SourceId:      2,
		DestinationId: 1,
		CabId:         1,
	},
	models.Booking{
		UserID:        2,
		SourceId:      2,
		DestinationId: 1,
		CabId:         1,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Booking{}, &models.Cab{}, &models.Location{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Location{}, &models.Cab{}, &models.Booking{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Cab{}).AddForeignKey("location_id", "locations(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Booking{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Booking{}).AddForeignKey("source_id", "locations(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Booking{}).AddForeignKey("destination_id", "locations(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Booking{}).AddForeignKey("cab_id", "cabs(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range locations {
		err = db.Debug().Model(&models.Location{}).Create(&locations[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range cabs {
		err = db.Debug().Model(&models.Cab{}).Create(&cabs[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
	for i, _ := range bookings {
		err = db.Debug().Model(&models.Booking{}).Create(&bookings[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	// for i, _ := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}
	// 	posts[i].AuthorID = users[i].ID

	// 	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed posts table: %v", err)
	// 	}
	// }
}
