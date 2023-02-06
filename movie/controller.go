package movie

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController interface{
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type movieController struct{
	movieUsecase MovieUsecase
}

func NewMovieController(movieUsecase MovieUsecase) MovieController{
	return &movieController{movieUsecase}
}

func (controller *movieController) GetAll(c *gin.Context){
	movie, err := controller.movieUsecase.GetAll(c)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": movie,
	})
}

func (controller *movieController) GetById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}
	
	movie, err := controller.movieUsecase.GetById(c, id)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": movie,
	})
}

func (controller *movieController) Add(c *gin.Context){
	var payload Movie
	c.ShouldBindJSON(&payload)

	movieID, err := controller.movieUsecase.Add(c, payload)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}

	type Movie struct{ID int `json:"id"`}

	c.JSON(200, gin.H{
		"message": "success",
		"data": Movie{ID: movieID},
	})
}

func (controller *movieController) Update(c *gin.Context){
	var payload Movie
	c.ShouldBindJSON(&payload)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "success",
		})
	}

	err = controller.movieUsecase.Update(c, id, payload)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}
	
	c.JSON(200, gin.H{
		"message": "success",
		"data": "",
	})
}

func (controller *movieController) Delete(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "success",
		})
	}

	err = controller.movieUsecase.Delete(c, id)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data": "",
		})
	}
	
	c.JSON(200, gin.H{
		"message": "success",
		"data": "",
	})
}