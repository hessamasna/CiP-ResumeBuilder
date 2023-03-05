package dto

type CVDto struct {
	ID             int              `json:"id"`
	UserID         uint             `json:"user_id"`
	PersonalInfo   PersonalInfoDto  `json:"personal_info"`
	About_me       string           `json:"about_me"`
	Educations     []EducationDto   `json:"education"`
	Experiences    []ExperienceDto  `json:"experience"`
	Skills         []SkillDto       `json:"skills"`
	SocialMedias   []SocialMediaDto `json:"social_medias"`
	Name           string           `json:"name"`
	IsPublic       bool             `json:"is_public"`
	FontSize       int              `json:"font_size"`
	FontFamily     string           `json:"font_family"`
	Color          string           `json:"color"`
	TemplateNumber int              `json:"template_number"`
}

type PersonalInfoDto struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type EducationDto struct {
	ID       int    `json:"id"`
	CVID     uint   `json:"cv_id"`
	Degree   string `json:"degree"`
	Major    string `json:"major"`
	School   string `json:"school"`
	Location string `json:"location"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

type ExperienceDto struct {
	ID          int    `json:"id"`
	CVID        uint   `json:"cv_id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
}

type SkillDto struct {
	ID      int    `json:"id"`
	CVID    uint   `json:"cv_id"`
	Name    string `json:"name"`
	Percent int    `json:"percent"`
}

type SocialMediaDto struct {
	ID       int    `json:"id"`
	CVID     uint   `json:"cv_id"`
	PlatForm string `json:"plat_form"`
	Link     string `json:"link"`
}
