package rusers

import (
	"caker/app/config"
	"caker/models/musers"

	"golang.org/x/crypto/bcrypt"
)

type UserRepoIMPL struct{}

func (ur *UserRepoIMPL) GetAllUser() []musers.User {
	var user []musers.User
	config.DB.Find(&user)
	return user
}

func (ur *UserRepoIMPL) GetByID(id string) musers.User {
	var user musers.User
	config.DB.First(&user, "id = ?", id)
	return user
}

func (ur *UserRepoIMPL) Register(input musers.Register) musers.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	var newUser musers.User = musers.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
		NoHp:     input.NoHp,
		Gender:   input.Gender,
		Photo:    input.Photo,
		Address:  input.Address,
	}
	config.DB.Create(&newUser)
	return newUser
}

func (ur *UserRepoIMPL) UpdateUser(id string, input musers.Register) musers.User {
	var user musers.User = ur.GetByID(id)
	user.Username = input.Username
	user.NoHp = input.NoHp
	user.Photo = input.Photo
	user.Address = input.Address
	res := config.DB.Save(&user)
	res.Last(&user)
	return user
}

func (ur *UserRepoIMPL) Login(input musers.Login) string {
	var user musers.User = musers.User{}
	config.DB.First(&user, "email", input.Email)
	if user.ID == 0 {
		return "Gagal Login"
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "Gagal Login"
	}
	return "Login Berhasil"
}

func (ur *UserRepoIMPL) DeleteUser(id string) bool {
	var user musers.User = ur.GetByID(id)

	res := config.DB.Delete(&user)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
