package movie

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_movieRepo_Create(t *testing.T){
	type fields struct{
		Conn *gorm.DB
	}
	type args struct{
		ctx context.Context
		input Movie
	}

	createdTime, _ := time.Parse("2006-01-02 15:04:05", "2022-08-01 10:56:31")
	updatedTime, _ := time.Parse("2006-01-02 15:04:05", "2022-08-13 09:30:23")
	
	tests := []struct{
		name		string
		args		args
		beforeTest	func(sqlmock.Sqlmock)
		want		int
		wantErr		bool
	}{
		{
			name: "fail save movie",
			args: args{
				ctx: context.TODO(),
				input: Movie{ID: 1, Title: "Pengabdi Setan 2 Communion", Description: "adalah sebuah film horor Indonesia.",
				Rating: 7, Image: "", CreatedAt: createdTime, UpdatedAt: updatedTime},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock){
				mockSQL.ExpectBegin()
				mockSQL.ExpectQuery(regexp.QuoteMeta(`INSERT INTO movies (id, title, description, rating, image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);`)).
				WithArgs(1, "Pengabdi Setan 2 Communion", "adalah sebuah film horor Indonesia.", 7, "", createdTime, updatedTime).
				WillReturnError(errors.New("whoops, error"))
				mockSQL.ExpectCommit()
			},
			wantErr: true,
		},
		{
			name: "success save movie",
			args: args{
				ctx: context.TODO(),
				input: Movie{ID: 1, Title: "Pengabdi Setan 2 Communion", Description: "adalah sebuah film horor Indonesia.",
				Rating: 7, Image: "", CreatedAt: createdTime, UpdatedAt: updatedTime},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock){
				mockSQL.ExpectBegin()
				mockSQL.ExpectExec(regexp.QuoteMeta(`INSERT INTO movies (id, title, description, rating, image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);`)).
				WithArgs(1, "Pengabdi Setan 2 Communion", "adalah sebuah film horor Indonesia.", 7, "", createdTime, updatedTime).
				WillReturnResult(sqlmock.NewResult(1, 1))
				mockSQL.ExpectCommit()
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			dialector := postgres.New(postgres.Config{
				DSN:                  "sqlmock_db_0",
				DriverName:           "postgres",
				Conn:                 mockDB,
				PreferSimpleProtocol: true,
			})

			dbCon, err := gorm.Open(dialector, &gorm.Config{})

			u := &movieRepository{
				Conn: dbCon,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Save(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("movieRepo.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movieRepo.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}