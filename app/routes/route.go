package routes

import (
	"caker/controllers/cjob"
	"caker/controllers/cmitras"
	"caker/controllers/cnews"
	"caker/controllers/csocmeds"
	"caker/controllers/ctraining"
	"caker/controllers/cusers"

	"github.com/labstack/echo/v4"
)

func SetRoute(server *echo.Echo) {
	server.GET("/user", cusers.GetAllUser)
	server.GET("/user/:id", cusers.GetByID)
	server.POST("/user/register", cusers.Register)
	server.POST("/user/login", cusers.Login)
	server.PUT("/user/:id", cusers.UpdateUser)
	server.DELETE("/user/:id", cusers.Delete)

	server.GET("/socmed", csocmeds.GetAllSocmed)
	server.GET("/socmed/:id", csocmeds.GetByIDSocmed)
	server.POST("/socmed", csocmeds.CreateSocmed)
	server.PUT("/socmed/:id", csocmeds.UpdateSocmed)
	server.DELETE("/socmed/:id", csocmeds.DeleteSocmed)

	server.GET("/mitra", cmitras.GetAllMitra)
	server.GET("/mitra/:id", cmitras.GetByIDMitra)
	server.POST("/mitra/register", cmitras.Register)
	server.POST("/mitra/login", cmitras.Login)
	server.PUT("mitra/:id", cmitras.UpdateMitra)
	server.DELETE("mitra/:id", cmitras.DeleteMitra)

	server.GET("/news", cnews.GetAllNews)
	server.GET("/news/:id", cnews.GetByIDNews)
	server.POST("/news", cnews.CreateNews)
	server.PUT("/news/:id", cnews.UpdateNews)
	server.DELETE("/news/:id", cnews.DeleteNews)

	server.GET("/training", ctraining.GetAllTraining)
	server.GET("/training/:id", ctraining.GetByIDTraining)
	server.POST("/training", ctraining.CreateTraining)
	server.PUT("/training/:id", ctraining.UpdateTraining)
	server.DELETE("/training/:id", ctraining.DeleteTraining)

	server.GET("/job", cjob.GetAllJob)
	server.GET("/job/:id", cjob.GetByIDJob)
	server.POST("/job", cjob.CreateJob)
	server.PUT("/job/:id", cjob.UpdateJob)
	server.DELETE("/job/:id", cjob.DeleteJob)

}
