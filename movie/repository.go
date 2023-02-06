package movie

import (
	"context"

	"gorm.io/gorm"
)

type MovieRepository interface{
	FindAll(c context.Context) (movie []Movie, err error)
	FindById(c context.Context, id int) (movie Movie, err error)
	Save(c context.Context, input Movie) (id int, err error)
	Update(c context.Context, id int, input Movie) (err error)
	Delete(c context.Context, id int) (err error)
}

type movieRepository struct{
	Conn *gorm.DB
}

func NewMovieRepository(Conn *gorm.DB) MovieRepository{
	return &movieRepository{Conn}
}

func(repo *movieRepository) FindAll(c context.Context) (movie []Movie, err error){
	err = repo.Conn.Find(&movie).Error

	if err != nil{
		return []Movie{}, err
	}

	return movie, nil
}
	
func(repo *movieRepository) FindById(c context.Context, id int) (movie Movie, err error){
	err = repo.Conn.Where("id = ?", id).First(&movie).Error

	if err != nil{
		return Movie{}, err
	}

	return movie, nil
}

func(repo *movieRepository) Save(c context.Context, input Movie) (id int, err error){
	err = repo.Conn.Save(&input).Error

	if err != nil{
		return 0, err
	}

	return input.ID, nil
}

func(repo *movieRepository) Update(c context.Context, id int, input Movie) (err error){
	err = repo.Conn.Where("id = ?", id).Updates(&input).Error

	return err
}

func(repo *movieRepository) Delete(c context.Context, id int) (err error){
	var movie Movie

	err = repo.Conn.Where("id = ?", id).Delete(&movie).Error

	return err
}