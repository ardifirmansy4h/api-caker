package config

import (
	"caker/models/mjob"
	"caker/models/mmitra"
	"caker/models/mnews"
	"caker/models/msocmeds"
	"caker/models/mtraining"
	"caker/models/musers"
)

func MigrateDB() {
	DB.AutoMigrate(&musers.User{}, &msocmeds.SocialMedia{}, &mmitra.Mitra{}, mnews.News{}, &mtraining.Training{}, &mjob.Job{})
}
