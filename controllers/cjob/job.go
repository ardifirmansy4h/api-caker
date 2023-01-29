package cjob

import (
	"caker/controllers"
	"caker/models/mjob"
	"caker/services/sjob"
	"net/http"

	"github.com/labstack/echo/v4"
)

var jobService sjob.JobService = sjob.NewJobRepo()

func GetAllJob(c echo.Context) error {
	job := jobService.GetAllJob()
	return controllers.NewResponse(c, http.StatusOK, "success", "all job", job)
}

func GetByIDJob(c echo.Context) error {
	jobID := c.Param("id")
	if jobID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak boleh kosong", "")
	}
	job := jobService.GetByIDJob(jobID)
	if job.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak boleh kosong", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data ditemukan", job)
}

func CreateJob(c echo.Context) error {
	cJob := new(mjob.InputJob)
	if err := c.Bind(cJob); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal membuat", "")
	}
	err := cJob.ValidJob()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}

	job := jobService.CreateJob(*cJob)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil dibuat", job)
}

func UpdateJob(c echo.Context) error {
	uJob := new(mjob.InputJob)
	if err := c.Bind(uJob); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah data", "")
	}
	jobID := c.Param("id")
	if jobID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "gagal mengubah data", "")
	}
	job := jobService.UpdateJob(jobID, *uJob)
	return controllers.NewResponse(c, http.StatusCreated, "success", "data berhasil diubah", job)
}

func DeleteJob(c echo.Context) error {
	jobID := c.Param("id")
	if jobID == "" {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "id tidak ada", "")
	}
	dJob := jobService.DeleteJob(jobID)
	if !dJob {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data gagal digapus", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "data berhasil dihapus", "")
	}
}
