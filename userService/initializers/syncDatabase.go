package initializers

import "userService/entities"

func SyncDatabase() {
	DB.AutoMigrate(&entities.CV{})
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Education{})
	DB.AutoMigrate(&entities.PersonalInfo{})
	DB.AutoMigrate(&entities.Experience{})
	DB.AutoMigrate(&entities.Skill{})
	DB.AutoMigrate(&entities.SocialMedia{})
}
