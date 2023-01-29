package susers

import (
	"caker/models/musers"
	"caker/repository"
	rusers "caker/repository/rusers"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService() UserService {
	return UserService{
		repo: &rusers.UserRepoIMPL{},
	}
}

func (us *UserService) GetAllUser() []musers.User {
	return us.repo.GetAllUser()
}

func (us *UserService) GetByID(id string) musers.User {
	return us.repo.GetByID(id)
}

func (us *UserService) UpdateUser(id string, input musers.Register) musers.User {
	return us.repo.UpdateUser(id, input)
}

func (us *UserService) Register(input musers.Register) musers.User {
	return us.repo.Register(input)
}

func (us *UserService) Login(input musers.Login) string {
	return us.repo.Login(input)
}

func (us *UserService) DeleteUser(id string) bool {
	return us.repo.DeleteUser(id)
}
