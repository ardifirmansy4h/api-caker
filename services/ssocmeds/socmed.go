package ssocmeds

import (
	"caker/models/msocmeds"
	"caker/repository"
	rsocmeds "caker/repository/rsocmeds"
)

type SocialMediaService struct {
	repo repository.SocmedRepo
}

func NewSocmedService() SocialMediaService {
	return SocialMediaService{
		repo: &rsocmeds.SocmedRepoIMPL{},
	}
}

func (ss *SocialMediaService) GetAllSocmed() []msocmeds.SocialMedia {
	return ss.repo.GetAllSocmed()
}

func (ss *SocialMediaService) GetByIDSocmed(id string) msocmeds.SocialMedia {
	return ss.repo.GetByIDSocmed(id)
}

func (ss *SocialMediaService) CreateSocmed(input msocmeds.InputSocmed) msocmeds.SocialMedia {
	return ss.repo.CreateSocmed(input)
}

func (ss *SocialMediaService) UpdateSocmed(id string, input msocmeds.InputSocmed) msocmeds.SocialMedia {
	return ss.repo.UpdateSocmed(id, input)
}

func (ss *SocialMediaService) DeleteSocmed(id string) bool {
	return ss.repo.DeleteSocmed(id)
}
