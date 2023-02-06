package server

import (
	"movie-services/config"
	"movie-services/movie"

	"github.com/gin-gonic/gin"
)

func RegisterMovieService(r *gin.Engine){
	db := config.GetDBCon()

	movieRepo := movie.NewMovieRepository(db)
	movieUsecase := movie.NewMovieUsecase(movieRepo)
	movieController := movie.NewMovieController(movieUsecase)

	registerMovieRoute(r, movieController)
}