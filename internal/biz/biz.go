package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUcUserUseCase,
	///...
)

// Transaction 数据库事务
type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// NewUcUserUseCase 初始化UcUser biz
func NewUcUserUseCase(repo IUcUserRepo, tm Transaction) *UcUserUseCase {
	return &UcUserUseCase{
		repo: repo,
		tm:   tm,
	}
}
