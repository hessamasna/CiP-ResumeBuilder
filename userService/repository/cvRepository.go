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
	if  err := AddEducationList(cvDto.Educations , cvEntity.ID); err != nil{
		tx.Rollback()
		return nil  , err
	}
	if err := AddExperienceList(cvDto.Experiences , cvEntity.ID); err != nil{
		tx.Rollback()
		return nil  , err
	}
	if err:= AddSkillList(cvDto.Skills , cvEntity.ID); err != nil{
		tx.Rollback()
		return nil , err
	}
	if err := AddSocialMediaList(cvDto.SocialMedias , cvEntity.ID); err != nil{
		tx.Rollback()
		return nil , err
	}
	return &cvDto, nil
}



func FindCVById(id int) (*dto.CVDto, *errors.Base_error) {
	var cvEntity entities.CV
	cv := initializers.DB.Preload("PersonalInfo").Preload("Education").Preload("Experience").Preload("Skills").Preload("SocialMedias").First(&cvEntity, id)
	// initializers.DB.First(&cvEntity, "ID = ?", id)
	if cv.Error != nil {
		// fmt.Print(cv.Error)
		return nil, nil
	}
	if cvEntity.ID == 0 {
		return nil, nil
	}
	var result dto.CVDto
	err := initializers.Mapper.Map(&result, cvEntity)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}
	return &result, nil
}

func FindAllCVByUserId(id int) (*dto.CVDto, *errors.Base_error) {
	var cvEntity entities.CV
	initializers.DB.First(&cvEntity, "user_id = ?", id)
	if cvEntity.ID == 0 {
		return nil, nil
	}
	var result dto.CVDto
	err := initializers.Mapper.Map(&result, cvEntity)
	if err != nil {
		//TODO log error
		return nil, errors.New_internal_error("Failed to map entity to dto", err).Error
	}
	return &result, nil
}
