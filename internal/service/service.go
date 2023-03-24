package service

import (
	"gin-layout/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewUserService,
)

func NewUserService(userUseCase *biz.UcUserUseCase) *UserService {
	return &UserService{
		uc: userUseCase,
	}
}
