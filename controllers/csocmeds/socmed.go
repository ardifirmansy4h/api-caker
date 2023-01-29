package csocmeds

import (
	"caker/controllers"
	"caker/models/msocmeds"
	"caker/services/ssocmeds"
	"net/http"

	"github.com/labstack/echo/v4"
)

var socmedService ssocmeds.SocialMediaService = ssocmeds.NewSocmedService()

func GetAllSocmed(c echo.Context) error {
	socmed := socmedService.GetAllSocmed()
	return controllers.NewResponse(c, http.StatusOK, "success", "all social media", socmed)
}

func GetByIDSocmed(c echo.Context) error {
	socmedID := c.Param("id")
	socmed := socmedService.GetByIDSocmed(socmedID)
	if socmed.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	return controllers.NewResponse(c, http.StatusAccepted, "success", "data ditemukan", socmed)
}

func CreateSocmed(c echo.Context) error {
	cSocmed := new(msocmeds.InputSocmed)
	if err := c.Bind(cSocmed); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal menambahkan data", "")
	}
	err := cSocmed.ValidSocmed()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}

	socmed := socmedService.CreateSocmed(*cSocmed)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil ditambahkan", socmed)
}

func UpdateSocmed(c echo.Context) error {
	uSocmed := new(msocmeds.InputSocmed)
	if err := c.Bind(uSocmed); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah data", "")
	}
	socmedID := c.Param("id")
	if socmedID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	socmed := socmedService.UpdateSocmed(socmedID, *uSocmed)
	return controllers.NewResponse(c, http.StatusBadRequest, "success", "data berhasil diubah", socmed)
}

func DeleteSocmed(c echo.Context) error {
	socmedID := c.Param("id")
	if socmedID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	dsocmed := socmedService.DeleteSocmed(socmedID)
	if !dsocmed {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data gagal digapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "data berhasil dihapus", "")
	}
}
