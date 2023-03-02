package entities

import (
	"gorm.io/gorm"
)

type CV struct {
	gorm.Model
	PersonalInfo PersonalInfo  `json:"personal_info" gorm:"embedded;embeddedPrefix:personal_info_"`
	About_me     string        `json:"about_me"`
	Education    []Education   `json:"education"`
	Experience   []Experience  `json:"experience"`
	Skills       []Skill       `json:"skills"`
	SocialMedias []SocialMedia `json:"social_medias"`
}

type PersonalInfo struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type Education struct {
	gorm.Model
	CVID     uint   `json:"cv_id"`
	Degree   string `json:"degree"`
	Major    string `json:"major"`
	School   string `json:"school"`
	Location string `json:"location"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type Experience struct {
	gorm.Model
	CVID        uint   `json:"cv_id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
}

type Skill struct {
	gorm.Model
	CVID        uint   `json:"cv_id"`
	Name        string `json:"name"`
	Percent     int    `json:"percent"`
	// Description string `json:"description"`
}

type SocialMedia struct {
	gorm.Model
	CVID     uint   `json:"cv_id"`
	PlatForm string `json:"plat_form"`
	Link     string `json:"link"`
}

