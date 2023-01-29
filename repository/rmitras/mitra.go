package rmitras

import (
	"caker/app/config"
	"caker/models/mmitra"

	"golang.org/x/crypto/bcrypt"
)

type MitraRepoIMPL struct{}

func (mi *MitraRepoIMPL) GetAllMitra() []mmitra.Mitra {
	var mitraa []mmitra.Mitra
	config.DB.Find(&mitraa)
	return mitraa
}

func (mi *MitraRepoIMPL) GetByIDMitra(id string) mmitra.Mitra {
	var mitra mmitra.Mitra
	config.DB.First(&mitra, "id = ? ", id)
	return mitra
}

func (mi *MitraRepoIMPL) Register(input mmitra.Register) mmitra.Mitra {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	var mitra mmitra.Mitra = mmitra.Mitra{
		Name:     input.Name,
		Photo:    input.Photo,
		Email:    input.Email,
		Password: string(password),
		Nohp:     input.Nohp,
		Address:  input.Address,
		Founder:  input.Founder,
	}
	res := config.DB.Create(&mitra)
	res.Last(&mitra)
	return mitra
}

func (mi *MitraRepoIMPL) UpdateMitra(id string, input mmitra.Register) mmitra.Mitra {
	var mitra mmitra.Mitra = mi.GetByIDMitra(id)
	mitra.Name = input.Name
	mitra.Photo = input.Photo
	mitra.Nohp = input.Nohp
	mitra.Address = input.Address
	mitra.Founder = input.Founder
	res := config.DB.Save(&mitra)
	res.Last(&mitra)
	return mitra
}

func (mi *MitraRepoIMPL) Login(input mmitra.Login) string {
	var mitra mmitra.Mitra = mmitra.Mitra{}
	config.DB.First(&mitra, "email", input.Email)
	if mitra.ID == 0 {
		return "Gagal login"
	}
	err := bcrypt.CompareHashAndPassword([]byte(mitra.Password), []byte(input.Password))
	if err != nil {
		return "gagal login"
	}
	return "Login Berhasil"
}

func (mi *MitraRepoIMPL) DeleteMitra(id string) bool {
	var mitra mmitra.Mitra = mi.GetByIDMitra(id)
	res := config.DB.Delete(&mitra)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
