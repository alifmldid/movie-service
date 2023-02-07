package movie

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository MovieRepository
	movie Movie
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	dbCon, err := gorm.Open(dialector, &gorm.Config{})

	require.NoError(s.T(), err)

	s.repository = NewMovieRepository(dbCon)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Create() {
	createdTime, err := time.Parse("2006-01-02 15:04:05", "2022-08-01 10:56:31")
	require.NoError(s.T(), err)

	updatedTime, err := time.Parse("2006-01-02 15:04:05", "2022-08-13 09:30:23")
	require.NoError(s.T(), err)

	var (
		id = 1
		title = "Pengabdi Setan 2 Comunion"
		description = "adalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan."
		rating = 7
		image = ""
	)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "movies" ("title","description","rating","image","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(title, description, float32(rating), image, createdTime, updatedTime, id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))
	s.mock.ExpectCommit()		

	id, err = s.repository.Save(context.TODO(), Movie{ID: id, Title: title, Description: description, Rating: float32(rating), Image: image, CreatedAt: createdTime, UpdatedAt: updatedTime})

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(1, id))
}
