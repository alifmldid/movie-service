package server

import (
	"movie-services/movie"

	"github.com/gin-gonic/gin"
)

func registerMovieRoute(r *gin.Engine, controller movie.MovieController){
	r.GET("/Movies", controller.GetAll)
	r.GET("/Movies/:id", controller.GetById)
	r.POST("/Movies", controller.Add)
	r.PATCH("/Movies/:id", controller.Update)
	r.DELETE("/Movies/:id", controller.Delete)
}