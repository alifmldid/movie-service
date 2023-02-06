package movie

import (
	"context"
	"time"
)

type MovieUsecase interface{
	GetAll(c context.Context) (movie []Movie, err error)
	GetById(c context.Context, id int) (movie Movie, err error)
	Add(c context.Context, payload Movie) (id int, err error)
	Update(c context.Context, id int, payload Movie) (err error)
	Delete(c context.Context, id int) (err error)
}

type movieUsecase struct{
	movieRepository MovieRepository
}

func NewMovieUsecase(movieRepository MovieRepository) MovieUsecase{
	return &movieUsecase{movieRepository}
}

func (uc *movieUsecase) GetAll(c context.Context) (movie []Movie, err error){
	movie, err = uc.movieRepository.FindAll(c)

	return movie, err
}

func (uc *movieUsecase) GetById(c context.Context, id int) (movie Movie, err error){
	movie, err = uc.movieRepository.FindById(c, id)

	return movie, err
}

func (uc *movieUsecase) Add(c context.Context, payload Movie) (id int, err error){
	var movie Movie

	movie.ID = payload.ID
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.Rating = payload.Rating
	movie.Image = payload.Image
	movie.CreatedAt = time.Now()
	movie.CreatedAt = time.Now()

	id, err = uc.movieRepository.Save(c, movie)

	return id, err
}

func (uc *movieUsecase) Update(c context.Context, id int, payload Movie) (err error){
	var movie = Movie{
		ID: payload.ID,
		Title: payload.Title,
		Description: payload.Description,
		Rating: payload.Rating,
		Image: payload.Image,
		CreatedAt: payload.CreatedAt,
		UpdatedAt: payload.UpdatedAt,
	}

	err = uc.movieRepository.Update(c, id, movie)

	return err
}

func (uc *movieUsecase) Delete(c context.Context, id int) (err error){
	err = uc.movieRepository.Delete(c, id)

	return err
}