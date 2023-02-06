package main

import (
	"movie-services/server"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	server.RegisterMovieService(r)

	r.Run(":8000")
}