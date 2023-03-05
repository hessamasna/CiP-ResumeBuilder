package repository

import (
	"fmt"
	"userService/dto"
	"userService/entities"
	"userService/errors"
	"userService/initializers"
)

func AddPersonalInfo(piDto dto.PersonalInfoDto) (*dto.PersonalInfoDto, *errors.Base_error) {
	var piEntity entities.PersonalInfo
	err := initializers.Mapper.Map(&piEntity, piDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&piEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	return &piDto, nil
}

func AddEducation(eduDto dto.EducationDto) (*dto.EducationDto, *errors.Base_error) {
	var eduEntity entities.Education
	err := initializers.Mapper.Map(&eduEntity, eduDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&eduEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	return &eduDto, nil
}

func AddExperience(expDto dto.ExperienceDto) (*dto.ExperienceDto, *errors.Base_error) {
	var expEntity entities.Experience
	err := initializers.Mapper.Map(&expEntity, expDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&expEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	return &expDto, nil
}

func AddSkill(skillDto dto.SkillDto) (*dto.SkillDto, *errors.Base_error) {
	var skillEntity entities.Skill
	err := initializers.Mapper.Map(&skillEntity, skillDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&skillEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	return &skillDto, nil
}

func AddSocialMedia(socialMediaDto dto.SocialMediaDto) (*dto.SocialMediaDto, *errors.Base_error) {
	var socialMediaEntity entities.SocialMedia
	err := initializers.Mapper.Map(&socialMediaEntity, socialMediaDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&socialMediaEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	return &socialMediaDto, nil
}

func AddEducationList(eduList []dto.EducationDto, cvID uint) *errors.Base_error {
	var eduEntities []entities.Education
	for _, edu := range eduList {
		eduEntity := entities.Education{
			CVID:     cvID,
			Degree:   edu.Degree,
			Major:    edu.Major,
			School:   edu.School,
			Location: edu.Location,
			Start:    edu.Start,
			End:      edu.End,
		}
		eduEntities = append(eduEntities, eduEntity)
	}

	tx := initializers.DB.Begin()
	if err := tx.Create(&eduEntities).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return errors.New_internal_error(err.Error(), err).Error
	}
	return nil
}

func AddExperienceList(experiences []dto.ExperienceDto, cvID uint) *errors.Base_error {
	tx := initializers.DB.Begin()
	for _, experienceDto := range experiences {
		experience := entities.Experience{
			CVID:        cvID,
			Title:       experienceDto.Title,
			Company:     experienceDto.Company,
			Location:    experienceDto.Location,
			Start:       experienceDto.Start,
			End:         experienceDto.End,
			Description: experienceDto.Description,
		}
		if err := tx.Create(&experience).Error; err != nil {
			tx.Rollback()
			return errors.New_internal_error(err.Error(), err).Error
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New_internal_error(err.Error(), err).Error
	}
	return nil
}

func AddSkillList(skills []dto.SkillDto, cvID uint) *errors.Base_error {
	tx := initializers.DB.Begin()
	for _, skillDto := range skills {
		skill := entities.Skill{
			CVID:    cvID,
			Name:    skillDto.Name,
			Percent: skillDto.Percent,
		}
		if err := tx.Create(&skill).Error; err != nil {
			tx.Rollback()
			return errors.New_internal_error(err.Error(), err).Error
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New_internal_error(err.Error(), err).Error
	}
	return nil
}

func AddSocialMediaList(socialMedias []dto.SocialMediaDto, cvID uint) *errors.Base_error {
	tx := initializers.DB.Begin()
	for _, socialMediaDto := range socialMedias {
		socialMedia := entities.SocialMedia{
			CVID:     cvID,
			Platform: socialMediaDto.PlatForm,
			Link:     socialMediaDto.Link,
		}
		if err := tx.Create(&socialMedia).Error; err != nil {
			tx.Rollback()
			return errors.New_internal_error(err.Error(), err).Error
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New_internal_error(err.Error(), err).Error
	}
	return nil
}

func AddCv(cvDto dto.CVDto) (*dto.CVDto, *errors.Base_error) {
	// fmt.Println("cvDto: ",  fmt.Sprintf("%v", cvDto))
	var cvEntity entities.CV
	err := initializers.Mapper.Map(&cvEntity, cvDto)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map dto to entity", err).Error
	}
	tx := initializers.DB.Begin()
	if err := tx.Create(&cvEntity).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err.Error())
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	if err := AddEducationList(cvDto.Educations, cvEntity.ID); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := AddExperienceList(cvDto.Experiences, cvEntity.ID); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := AddSkillList(cvDto.Skills, cvEntity.ID); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := AddSocialMediaList(cvDto.SocialMedias, cvEntity.ID); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &cvDto, nil
}


func GetCVsByUserId(userId int) ([]dto.CVDto, *errors.Base_error) {
	var cvEntities []entities.CV
	err := initializers.DB.Where("user_id = ?", userId).Find(&cvEntities).Error
	if err != nil {
		return nil, errors.New_internal_error(err.Error(), err).Error
	}
	// fmt.Print(cvEntities)

	var cvDtos []dto.CVDto
	for _, cvEntity := range cvEntities {
		cvid := cvEntity.ID
		skills, err := GetAllSkillsByCVID(cvid)
		if err != nil {
			return nil, err
		}
		educations, err := GetAllEducationsByCVID(cvid)
		if err != nil {
			return nil, err
		}
		experiences, err := GetAllExperiencesByCVID(cvid)
		if err != nil {
			return nil, err
		}
		social_medias, err := GetAllSocialMediasByCVID(cvid)
		if err != nil {
			return nil, err
		}
		var cvDto dto.CVDto
		er1 := initializers.Mapper.Map(&cvDto, cvEntity)
		if er1 != nil {
			return nil, errors.New_internal_error("Failed to map entities to dtos", er1).Error
		}
		cvDto.Educations = educations
		cvDto.Skills = skills
		cvDto.Experiences = experiences
		cvDto.SocialMedias = social_medias
		cvDtos = append(cvDtos, cvDto)
	}

	return cvDtos, nil
}

func GetCvById(id int) (*dto.CVDto, *errors.Base_error) {
	var cv entities.CV
	error := initializers.DB.Where("id = ?", id).First(&cv).Error
	if cv.ID == 0 {
		error_message := fmt.Sprintf("Cv with id : %d does not exist.", id)
		return nil, errors.New_entity_does_not_exist_error(error_message, nil).Error
	}
	if error != nil {
		error_message := fmt.Sprintf("Error getting CV with ID %d", id)
		return nil, errors.New_internal_error(error_message, error).Error
	}

	var cvDto dto.CVDto
	if err := initializers.Mapper.Map(&cvDto, &cv); err != nil {
		error_message := fmt.Sprintf("Error mapping CV with ID %d to DTO", id)
		return nil, errors.New_internal_error(error_message, err).Error
	}

	skills, err := GetAllSkillsByCVID(cv.ID)
	if err != nil {
		return nil, err
	}
	educations, err := GetAllEducationsByCVID(cv.ID)
	if err != nil {
		return nil, err
	}
	experiences, err := GetAllExperiencesByCVID(cv.ID)
	if err != nil {
		return nil, err
	}
	social_medias, err := GetAllSocialMediasByCVID(cv.ID)
	if err != nil {
		return nil, err
	}

	cvDto.Educations = educations
	cvDto.Skills = skills
	cvDto.Experiences = experiences
	cvDto.SocialMedias = social_medias
	return &cvDto, nil
}

func GetCVsByUserIdWithPagination(userID uint, page, size int) ([]dto.CVDto, *errors.Base_error) {
	var cvEntities []entities.CV
	var cvDtos []dto.CVDto

	// retrieve CV entities for the given user ID, paginated
	err := initializers.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset((page - 1) * size).
		Limit(size).
		Find(&cvEntities).Error

	if err != nil {
		return nil, errors.New_internal_error("Failed to retrieve CVs", err).Error
	}

	// map CV entities to DTOs
	for _, cvEntity := range cvEntities {
		var cvDto dto.CVDto
		err := initializers.Mapper.Map(&cvDto, cvEntity)
		if err != nil {
			//TODO log error
			return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
		}
		cvDtos = append(cvDtos, cvDto)
	}

	return cvDtos, nil

}

func GetAllEducationsByCVID(cvid uint) ([]dto.EducationDto, *errors.Base_error) {
	var educations []entities.Education
	if err := initializers.DB.Where("cv_id = ?", cvid).Find(&educations).Error; err != nil {
		return nil, errors.New_internal_error("Failed to get educations for CV", err).Error
	}

	educationDtos := make([]dto.EducationDto, len(educations))
	if err := initializers.Mapper.Map(&educationDtos, educations); err != nil {
		return nil, errors.New_internal_error("Failed to map educations to dtos", err).Error
	}

	return educationDtos, nil
}

func GetAllExperiencesByCVID(cvid uint) ([]dto.ExperienceDto, *errors.Base_error) {
	var experiences []entities.Experience
	if err := initializers.DB.Where("cv_id = ?", cvid).Find(&experiences).Error; err != nil {
		return nil, errors.New_internal_error("Failed to get experiences for CV", err).Error
	}

	experienceDtos := make([]dto.ExperienceDto, len(experiences))
	if err := initializers.Mapper.Map(&experienceDtos, experiences); err != nil {
		return nil, errors.New_internal_error("Failed to map experiences to dtos", err).Error
	}

	return experienceDtos, nil
}

func GetAllSkillsByCVID(cvid uint) ([]dto.SkillDto, *errors.Base_error) {
	var skills []entities.Skill
	if err := initializers.DB.Where("cv_id = ?", cvid).Find(&skills).Error; err != nil {
		return nil, errors.New_internal_error("Failed to get skills for CV", err).Error
	}

	skillDtos := make([]dto.SkillDto, len(skills))
	if err := initializers.Mapper.Map(&skillDtos, skills); err != nil {
		return nil, errors.New_internal_error("Failed to map skills to dtos", err).Error
	}

	return skillDtos, nil
}

func GetAllSocialMediasByCVID(cvid uint) ([]dto.SocialMediaDto, *errors.Base_error) {
	var socialMedias []entities.SocialMedia
	if err := initializers.DB.Where("cv_id = ?", cvid).Find(&socialMedias).Error; err != nil {
		return nil, errors.New_internal_error("Failed to get social medias for CV", err).Error
	}

	socialMediaDtos := make([]dto.SocialMediaDto, len(socialMedias))
	if err := initializers.Mapper.Map(&socialMediaDtos, socialMedias); err != nil {
		return nil, errors.New_internal_error("Failed to map social medias to dtos", err).Error
	}

	return socialMediaDtos, nil
}
