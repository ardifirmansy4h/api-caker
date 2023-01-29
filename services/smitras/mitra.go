package smitras

import (
	"caker/models/mmitra"
	"caker/repository"
	"caker/repository/rmitras"
)

type MitraService struct {
	repo repository.MitraRepo
}

func NewMitraService() MitraService {
	return MitraService{
		repo: &rmitras.MitraRepoIMPL{},
	}
}

func (mss *MitraService) GetAllMitra() []mmitra.Mitra {
	return mss.repo.GetAllMitra()
}

func (ms *MitraService) GetByIDMitra(id string) mmitra.Mitra {
	return ms.repo.GetByIDMitra(id)
}

func (ms *MitraService) Register(input mmitra.Register) mmitra.Mitra {
	return ms.repo.Register(input)
}

func (ms *MitraService) UpdateMitra(id string, input mmitra.Register) mmitra.Mitra {
	return ms.repo.UpdateMitra(id, input)
}

func (ms *MitraService) Login(input mmitra.Login) string {
	return ms.repo.Login(input)
}
func (ms *MitraService) DeleteMitra(id string) bool {
	return ms.repo.DeleteMitra(id)
}
