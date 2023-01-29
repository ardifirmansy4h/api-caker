package snews

import (
	"caker/models/mnews"
	"caker/repository"
	"caker/repository/rnews"
)

type NewsService struct {
	repo repository.NewsRepo
}

func NewNewsService() NewsService {
	return NewsService{
		repo: &rnews.NewsRepoIMPL{},
	}
}

func (ns *NewsService) GetAllNews() []mnews.News {
	return ns.repo.GetAllNews()
}

func (ns *NewsService) GetByIDNews(id string) mnews.News {
	return ns.repo.GetByIDNews(id)
}

func (ns *NewsService) CreateNews(input mnews.InputNews) mnews.News {
	return ns.repo.CreateNews(input)
}

func (ns *NewsService) UpdateNews(id string, input mnews.UpdateNews) mnews.News {
	return ns.repo.UpdateNews(id, input)
}

func (ns *NewsService) DeleteNews(id string) bool {
	return ns.repo.DeleteNews(id)
}
