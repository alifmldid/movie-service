package movie

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockMovieRepo struct{
	ctrl		*gomock.Controller
	recorder	*MockMovieRepoMockRecorder
}

type MockMovieRepoMockRecorder struct{
	mock *MockMovieRepo
}

func NewMockMovieRepo(ctrl *gomock.Controller) *MockMovieRepo{
	mock := &MockMovieRepo{ctrl: ctrl}
	mock.recorder = &MockMovieRepoMockRecorder{mock}

	return mock
}

func (m *MockMovieRepo) EXPECT() *MockMovieRepoMockRecorder{
	return m.recorder
}

func (m *MockMovieRepo) FindAll(ctx context.Context) ([]Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMovieRepoMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockMovieRepo)(nil).FindAll), ctx)
}

func (m *MockMovieRepo) FindById(ctx context.Context, id int) (Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMovieRepoMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockMovieRepo)(nil).FindById), ctx, id)
}

func (m *MockMovieRepo) Save(ctx context.Context, input Movie) (int, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMovieRepoMockRecorder) Save(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMovieRepo)(nil).Save), ctx, input)
}

func (m *MockMovieRepo) Update(ctx context.Context, id int, input Movie) (error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockMovieRepoMockRecorder) Update(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovieRepo)(nil).Update), ctx, id, input)
}


func (m *MockMovieRepo) Delete(ctx context.Context, id int) error{
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockMovieRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockMovieRepo)(nil).Delete), ctx, id)
}
