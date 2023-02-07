package movie

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_addMovieUseCase_Execute(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewMockMovieRepo(ctrl)
	fmt.Println(c)

	type args struct {
		ctx   context.Context
		input Movie
	}

	createdTime, _ := time.Parse("2022-08-01 10:56:31", "2022-08-01 10:56:31")
	updatedTime, _ := time.Parse("2022-08-13 09:30:23", "2022-08-13 09:30:23")
	
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *MockMovieRepo)
		want       int
		wantErr    bool
	}{
		{
			name: "success creating new user",
			args: args{
				ctx: context.TODO(),
				input: Movie{
					ID: 1,
					Title: "Pengabdi Setan 2 Comunion",
					Description: "adalah sebuah film horor Indonesia",
					Rating: 7,
					Image: "",
					CreatedAt: createdTime,
					UpdatedAt: updatedTime,
				},
			},
			beforeTest: func(userRepo *MockMovieRepo) {
				userRepo.EXPECT().
					Save(
						context.TODO(),
						Movie{
							ID: 1,
							Title: "Pengabdi Setan 2 Comunion",
							Description: "adalah sebuah film horor Indonesia",
							Rating: 7,
							Image: "",
							CreatedAt: createdTime,
							UpdatedAt: updatedTime,
						},
					).
					Return(
						1,
						nil,
					)
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMovieRepo := NewMockMovieRepo(ctrl)

			w := &movieUsecase{
				movieRepository: mockMovieRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockMovieRepo)
			}

			got, err := w.Add(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerUserUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerUserUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}	
}