package repository

import (
	"caker/models/mjob"
	"caker/models/mmitra"
	"caker/models/mnews"
	"caker/models/msocmeds"
	"caker/models/mtraining"
	"caker/models/musers"
)

type UserRepo interface {
	GetAllUser() []musers.User
	GetByID(id string) musers.User
	Register(input musers.Register) musers.User
	UpdateUser(id string, input musers.Register) musers.User
	Login(input musers.Login) string
	DeleteUser(id string) bool
}

type MitraRepo interface {
	GetAllMitra() []mmitra.Mitra
	GetByIDMitra(id string) mmitra.Mitra
	Register(input mmitra.Register) mmitra.Mitra
	UpdateMitra(id string, input mmitra.Register) mmitra.Mitra
	Login(input mmitra.Login) string
	DeleteMitra(id string) bool
}

type SocmedRepo interface {
	GetAllSocmed() []msocmeds.SocialMedia
	GetByIDSocmed(id string) msocmeds.SocialMedia
	CreateSocmed(input msocmeds.InputSocmed) msocmeds.SocialMedia
	UpdateSocmed(id string, input msocmeds.InputSocmed) msocmeds.SocialMedia
	DeleteSocmed(id string) bool
}

type NewsRepo interface {
	GetAllNews() []mnews.News
	GetByIDNews(id string) mnews.News
	CreateNews(input mnews.InputNews) mnews.News
	UpdateNews(id string, input mnews.UpdateNews) mnews.News
	DeleteNews(id string) bool
}

type TrainingRepo interface {
	GetAllTraining() []mtraining.Training
	GetByIDTraining(id string) mtraining.Training
	CreateTraining(input mtraining.InputTraining) mtraining.Training
	UpdateTraining(id string, input mtraining.InputTraining) mtraining.Training
	DeleteTraining(id string) bool
}

type JobRepo interface {
	GetAllJob() []mjob.Job
	GetByIDJob(id string) mjob.Job
	CreateJob(input mjob.InputJob) mjob.Job
	UpdateJob(id string, input mjob.InputJob) mjob.Job
	DeleteJob(id string) bool
}
