package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	PasswordHash string
	PhoneNumber  string `gorm:"unique;not null"`
	Gender       string
	FirstName    string
	LastName     string
}

type CV struct {
	gorm.Model
	UserID         uint         `json:"user_id" gorm:"not null"`
	PersonalInfo   PersonalInfo `json:"personal_info" gorm:"embedded;embeddedPrefix:personal_info_"`
	AboutMe        string       `json:"about_me"`
	Title          string       `json:"title" gorm:"not null"`
	IsPublic       bool         `json:"is_public" gorm:"not null;default:false"`
	FontSize       int          `json:"font_size" gorm:"not null"`
	FontFamily     string       `json:"font_family" gorm:"not null"`
	Color          string       `json:"color" gorm:"not null"`
	TemplateNumber int          `json:"template_number" gorm:"not null"`
	JobTitle       string       `json:"job_title"`
	Image          string       `json:"image"`

	// Education     []Education   `json:"education" `
	// Experience    []Experience  `json:"experience" `
	// Skills        []Skill       `json:"skills" `
	// SocialMedias  []SocialMedia `json:"social_medias" `

}

type PersonalInfo struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type Education struct {
	gorm.Model
	CVID     uint   `json:"cv_id " gorm:"polymorphic:Owner;"`
	Degree   string `json:"degree"`
	Major    string `json:"major"`
	School   string `json:"school"`
	Location string `json:"location"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type Experience struct {
	gorm.Model
	CVID        uint   `json:"cv_id" gorm:"polymorphic:Owner;"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
}

type Skill struct {
	gorm.Model
	CVID    uint   `json:"cv_id" gorm:"polymorphic:Owner;"`
	Name    string `json:"name"`
	Percent int    `json:"percent"`
}

type SocialMedia struct {
	gorm.Model
	CVID     uint   `json:"cv_id" gorm:"polymorphic:Owner; "`
	Platform string `json:"platform"`
	Link     string `json:"link"`
}
