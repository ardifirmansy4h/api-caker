package rnews

import (
	"caker/app/config"
	"caker/models/mnews"
)

type NewsRepoIMPL struct{}

func (nr *NewsRepoIMPL) GetAllNews() []mnews.News {
	var news []mnews.News
	config.DB.Preload("Mitra").Find(&news)
	return news
}

func (nr *NewsRepoIMPL) GetByIDNews(id string) mnews.News {
	var news mnews.News
	config.DB.First(&news, "id = ?", id).Preload("Mitra").Find(&news)
	return news
}

func (nr *NewsRepoIMPL) CreateNews(input mnews.InputNews) mnews.News {
	var news mnews.News = mnews.News{
		MitraID: input.MitraID,
		Title:   input.Title,
		Blog:    input.Blog,
		Photo:   input.Photo,
	}
	res := config.DB.Create(&news).Preload("Mitra").Find(&news)
	res.Last(&news)
	return news
}

func (nr *NewsRepoIMPL) UpdateNews(id string, input mnews.UpdateNews) mnews.News {
	var news mnews.News = nr.GetByIDNews(id)
	news.Title = input.Title
	news.Blog = input.Blog
	news.Photo = input.Photo
	res := config.DB.Save(&news).Preload("Mitra").Find(&news)
	res.Last(&news)
	return news
}

func (nr *NewsRepoIMPL) DeleteNews(id string) bool {
	var news mnews.News = nr.GetByIDNews(id)
	res := config.DB.Delete(&news)
	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
