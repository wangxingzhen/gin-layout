package biz

import (
	"context"
	"gin-layout/internal/pkg/page"
	"gin-layout/internal/pkg/validate"
	"gin-layout/pkg/errResponse"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UcUser struct {
	Id           uint64
	Name         string `validate:"required,min=1,max=20" label:"名称"`
	SerialNumber int
}

// IUcUserRepo 操作用户表 。。。 【在biz层规定data层要实现的功能】
type IUcUserRepo interface {
	CreateUcUser(ctx context.Context, a *UcUser) error
	GetUcUserById(ctx context.Context, id uint64) (*UcUser, error)
	GetUcUserNum(ctx context.Context) (int, error)
	GetUcUserMaxId(ctx context.Context) (uint64, error)
	SaveUcUserSerialNumber(ctx context.Context, a *UcUser) error
	UserList(ctx context.Context, condition *ListTestRep) ([]*UcUser, error)
}

type UcUserUseCase struct {
	repo IUcUserRepo
	tm   Transaction
}

func (u *UcUserUseCase) GetTest(ctx context.Context, user *UcUser) (*UcUser, error) {
	if user.Id < 1 {
		return &UcUser{}, nil
	}
	res, err := u.repo.GetUcUserById(ctx, user.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonDataIsNotFount)
	}
	return res, err
}

func (u *UcUserUseCase) AddTest(ctx context.Context, user *UcUser) error {
	if err := errors.WithStack(validate.ValidateStruct(user)); err != nil {
		return err
	}
	return u.repo.CreateUcUser(ctx, user)
}

// TranTest 事务使用 (示例)
func (u *UcUserUseCase) TranTest(ctx context.Context) error {
	err := u.tm.InTx(ctx, func(ctx context.Context) error {
		// 获取现在有多少条
		n, err := u.repo.GetUcUserNum(ctx)
		if err != nil {
			return err
		}
		// 获取最大id
		maxId, err := u.repo.GetUcUserMaxId(ctx)
		if err != nil {
			return err
		}
		// 修改最大的那一条数据中SerialNumber字段
		err = u.repo.SaveUcUserSerialNumber(ctx, &UcUser{
			Id:           maxId,
			SerialNumber: n,
		})
		return err
	})
	return err
}

// ListTestRep 查询用户列表
type ListTestRep struct {
	Page *page.Page
	Id   uint64
	Name string
}

// ListTest 用户列表
func (u *UcUserUseCase) ListTest(ctx context.Context, condition *ListTestRep) ([]*UcUser, error) {
	return u.repo.UserList(ctx, condition)
}
