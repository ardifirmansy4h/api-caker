package ctraining

import (
	"caker/controllers"
	"caker/models/mtraining"
	"caker/services/straining"
	"net/http"

	"github.com/labstack/echo/v4"
)

var trainingService straining.TrainingService = straining.NewTrainingService()

func GetAllTraining(c echo.Context) error {
	training := trainingService.GetAllTraining()
	return controllers.NewResponse(c, http.StatusOK, "success", "all training", training)
}

func GetByIDTraining(c echo.Context) error {
	trainingID := c.Param("id")
	if trainingID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	training := trainingService.GetByIDTraining(trainingID)
	if training.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data ditemukan", training)
}

func CreateTraining(c echo.Context) error {
	cTraining := new(mtraining.InputTraining)
	if err := c.Bind(cTraining); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal membuat", "")
	}
	err := cTraining.ValidTraining()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}

	training := trainingService.CreateTraining(*cTraining)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil ditambah", training)
}

func UpdateTraining(c echo.Context) error {
	uTraining := new(mtraining.InputTraining)
	if err := c.Bind(uTraining); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah", "")
	}
	trainingID := c.Param("id")
	if trainingID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	training := trainingService.UpdateTraining(trainingID, *uTraining)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil diubah", training)
}

func DeleteTraining(c echo.Context) error {
	trainingID := c.Param("id")
	if trainingID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	dtraining := trainingService.DeleteTraining(trainingID)
	if !dtraining {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data gagal digapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "data berhasil dihapus", "")
	}
}
