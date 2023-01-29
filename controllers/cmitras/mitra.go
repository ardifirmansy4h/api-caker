package cmitras

import (
	"caker/controllers"
	"caker/models/mmitra"
	"caker/services/smitras"
	"net/http"

	"github.com/labstack/echo/v4"
)

var mitraService smitras.MitraService = smitras.NewMitraService()

func GetAllMitra(c echo.Context) error {
	mitra := mitraService.GetAllMitra()
	return controllers.NewResponse(c, http.StatusOK, "succes", "data ditemukan", mitra)
}

func GetByIDMitra(c echo.Context) error {
	mitraID := c.Param("id")
	mitra := mitraService.GetByIDMitra(mitraID)
	if mitra.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak ada", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data ditemukan", mitra)
}

func Register(c echo.Context) error {
	mitraRegis := new(mmitra.Register)
	if err := c.Bind(mitraRegis); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal daftar", "")
	}
	err := mitraRegis.ValidRegister()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}

	mitra := mitraService.Register(*mitraRegis)
	return controllers.NewResponse(c, http.StatusCreated, "success", "berhasil daftar", mitra)
}

func UpdateMitra(c echo.Context) error {
	mitraUpdate := new(mmitra.Register)
	if err := c.Bind(mitraUpdate); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah", "")
	}
	mitraID := c.Param("id")
	if mitraID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak ada", "")
	}
	mitra := mitraService.UpdateMitra(mitraID, *mitraUpdate)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil diubah", mitra)
}

func Login(c echo.Context) error {
	loginMitra := new(mmitra.Login)
	if err := c.Bind(*loginMitra); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal login", "")
	}
	mitra := mitraService.Login(*loginMitra)
	return controllers.NewResponse(c, http.StatusOK, "success", "berhasil login", mitra)
}

func DeleteMitra(c echo.Context) error {
	mitraID := c.Param("id")
	if mitraID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak ada", "")
	}
	mitra := mitraService.DeleteMitra(mitraID)

	if !mitra {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal dihapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "berhasil dihapus", "")
	}
}
