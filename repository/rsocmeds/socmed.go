package rsocmeds

import (
	"caker/app/config"
	"caker/models/msocmeds"
)

type SocmedRepoIMPL struct{}

func (sr *SocmedRepoIMPL) GetAllSocmed() []msocmeds.SocialMedia {
	var socmed []msocmeds.SocialMedia
	config.DB.Preload("User").Find(&socmed)
	return socmed
}

func (sr *SocmedRepoIMPL) GetByIDSocmed(id string) msocmeds.SocialMedia {
	var socmed msocmeds.SocialMedia
	config.DB.First(&socmed, "id = ?", id).Preload("User").Find(&socmed)
	return socmed
}

func (sr *SocmedRepoIMPL) CreateSocmed(input msocmeds.InputSocmed) msocmeds.SocialMedia {
	var socmed msocmeds.SocialMedia = msocmeds.SocialMedia{
		UserID:  input.UserID,
		Caption: input.Caption,
		Photo:   input.Photo,
	}

	res := config.DB.Create(&socmed).Preload("User").Find(&socmed)
	res.Last(&socmed)
	return socmed
}

func (sr *SocmedRepoIMPL) UpdateSocmed(id string, input msocmeds.InputSocmed) msocmeds.SocialMedia {
	var socmed msocmeds.SocialMedia = sr.GetByIDSocmed(id)
	socmed.Caption = input.Caption
	config.DB.Save(&socmed)
	return socmed
}

func (sr *SocmedRepoIMPL) DeleteSocmed(id string) bool {
	var socmed msocmeds.SocialMedia = sr.GetByIDSocmed(id)
	res := config.DB.Delete(&socmed)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
