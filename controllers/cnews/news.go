package cnews

import (
	"caker/controllers"
	"caker/models/mnews"
	"caker/services/snews"
	"net/http"

	"github.com/labstack/echo/v4"
)

var newsService snews.NewsService = snews.NewNewsService()

func GetAllNews(c echo.Context) error {
	news := newsService.GetAllNews()
	return controllers.NewResponse(c, http.StatusOK, "success", "all news", news)
}

func GetByIDNews(c echo.Context) error {
	newsID := c.Param("id")
	news := newsService.GetByIDNews(newsID)
	if news.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	} else if newsID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "success", "id tidak ada", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "Data ditemukan", news)
	}
}

func CreateNews(c echo.Context) error {
	cNews := new(mnews.InputNews)
	if err := c.Bind(cNews); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal membuat", "")
	}
	err := cNews.ValidNews()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}

	news := newsService.CreateNews(*cNews)
	return controllers.NewResponse(c, http.StatusCreated, "succes", "data berhasil ditambah", news)

}

func UpdateNews(c echo.Context) error {
	cNews := new(mnews.UpdateNews)
	if err := c.Bind(cNews); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah", "")
	}
	newsID := c.Param("id")
	if newsID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "success", "id tidak ada", "")
	}
	news := newsService.UpdateNews(newsID, *cNews)
	return controllers.NewResponse(c, http.StatusCreated, "success", "berhasil diubah", news)
}

func DeleteNews(c echo.Context) error {
	newsID := c.Param("id")
	if newsID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	dNews := newsService.DeleteNews(newsID)
	if !dNews {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data gagal digapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "data berhasil dihapus", "")
	}
}
