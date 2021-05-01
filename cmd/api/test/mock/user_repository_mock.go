package mock

import (
	"github.com/PubApi/cmd/api/app/domain/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) Save(user model.User) (err error) {
	args := mock.Called(user)
	return args.Error(0)
}
