package movie

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockAddMovieUsecase struct{
	ctrl *gomock.Controller
	recorder *MockAddMovieUsecaseMockRecorder
}

type MockAddMovieUsecaseMockRecorder struct{
	mock *MockAddMovieUsecase
}

func NewMockAddMovieUsecase(ctrl *gomock.Controller) *MockAddMovieUsecase{
	mock := &MockAddMovieUsecase{ctrl: ctrl}
	mock.recorder = &MockAddMovieUsecaseMockRecorder{mock}

	return mock
}

func (m *MockAddMovieUsecase) EXPECT() *MockAddMovieUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockAddMovieUsecase) Execute(ctx context.Context, input Movie) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockAddMovieUsecaseMockRecorder) Execute(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockAddMovieUsecase)(nil).Execute), ctx, input)
}
