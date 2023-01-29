package cusers

import (
	"caker/controllers"
	"caker/models/musers"
	"caker/services/susers"
	"net/http"

	"github.com/labstack/echo/v4"
)

var userService susers.UserService = susers.NewUserService()

func GetAllUser(c echo.Context) error {
	users := userService.GetAllUser()
	return controllers.NewResponse(c, http.StatusAccepted, "succes", "all user", users)
}

func GetByID(c echo.Context) error {
	userID := c.Param("id")
	users := userService.GetByID(userID)
	if users.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak ada", "")
	}
	return controllers.NewResponse(c, http.StatusAccepted, "succes", "data ditemukan", users)
}

func Register(c echo.Context) error {
	userRegis := new(musers.Register)
	if err := c.Bind(userRegis); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal daftar", "")
	}
	err := userRegis.ValidInputRegister()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}
	user := userService.Register(*userRegis)
	return controllers.NewResponse(c, http.StatusCreated, "succes", "berhasil daftar", user)
}

func UpdateUser(c echo.Context) error {
	userUpdate := new(musers.Register)
	if err := c.Bind(userUpdate); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah", "")
	}
	userID := c.Param("id")
	if userID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	user := userService.UpdateUser(userID, *userUpdate)
	return controllers.NewResponse(c, http.StatusCreated, "success", "berhasil update", user)
}

func Login(c echo.Context) error {
	userLogin := new(musers.Login)
	if err := c.Bind(userLogin); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal login", "")
	}
	login := userService.Login(*userLogin)
	return controllers.NewResponse(c, http.StatusBadRequest, "succes", "berhasil login", login)
}

func Delete(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	user := userService.DeleteUser(userID)
	if !user {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal dihapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "berhasil dihapus", "")
	}
}
