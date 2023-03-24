package service

import (
	"gin-layout/internal/biz"
	"gin-layout/internal/pkg/copierx"
	"gin-layout/internal/pkg/page"
	"gin-layout/pkg/errResponse"
	"gin-layout/pkg/ginx"
	"github.com/pkg/errors"
)

// UserService .
type UserService struct {
	uc *biz.UcUserUseCase
}

type TestReq struct {
	Id uint64 `form:"id" binding:"omitempty,gte=1"` // id
}

type TestReply struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// Test 获取单条数据
func (s *UserService) Test(ctx *ginx.RequestContext) (any, error) {
	var err error
	req := &TestReq{}

	if err = ctx.Context.ShouldBindQuery(req); err != nil {
		return nil, errResponse.SetCustomizeErrMsgByReason(errResponse.ReasonParamsError, err.Error())
	}

	r := &biz.UcUser{}
	// copy过去筛选参数
	err = errors.WithStack(copierx.Copy(r, req))
	if err != nil {
		return nil, err
	}
	r, err = s.uc.GetTest(ctx.Context, r)
	if err != nil {
		return nil, err
	}

	data := &TestReply{}
	if err = errors.WithStack(copierx.Copy(&data, r)); err != nil {
		return nil, err
	}
	return data, nil
}

type AddTestReq struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
}

// AddTest 添加数据
func (s *UserService) AddTest(ctx *ginx.RequestContext) (any, error) {
	var err error
	req := &AddTestReq{}

	if err = ctx.Context.ShouldBindJSON(req); err != nil {
		return nil, errResponse.SetCustomizeErrMsgByReason(errResponse.ReasonParamsError, err.Error())
	}
	u := &biz.UcUser{}
	if err = errors.WithStack(copierx.Copy(&u, req)); err != nil {
		return nil, err
	}
	return nil, s.uc.AddTest(ctx.Context, u)
}

// TranTest 事务使用示例
func (s *UserService) TranTest(ctx *ginx.RequestContext) (any, error) {
	return nil, s.uc.TranTest(ctx.Context)
}

func (s *UserService) IsAdmin(ctx *ginx.RequestContext) (any, error) {
	return nil, nil
}

type ListTestReq struct {
	PageNum  uint64 `form:"pageNum" binding:"omitempty,gte=1"`
	PageSize uint64 `form:"pageSize" binding:"omitempty,gte=1"`
	Id       uint64 `form:"id" binding:"omitempty,gte=1"`   // id
	Name     string `form:"name" binding:"omitempty,min=1"` // 名称
}

type ListTestReply struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	SerialNumber int    `json:"serial_number"`
}

// ListTest 分页获取多条数据
func (s *UserService) ListTest(ctx *ginx.RequestContext) (any, error) {
	var err error
	req := &ListTestReq{}

	if err = ctx.Context.ShouldBindQuery(req); err != nil {
		return nil, errResponse.SetCustomizeErrMsgByReason(errResponse.ReasonParamsError, err.Error())
	}
	condition := &biz.ListTestRep{
		Page: &page.Page{
			Num:  req.PageNum,
			Size: req.PageSize,
		},
	}
	// copy过去筛选参数
	err = errors.WithStack(copierx.Copy(condition, req))
	if err != nil {
		return nil, err
	}

	r, err := s.uc.ListTest(ctx.Context, condition)
	if err != nil {
		return nil, err
	}
	res := make([]*ListTestReply, 0)
	err = errors.WithStack(copierx.Copy(&res, r))
	if err != nil {
		return nil, err
	}
	return ctx.ReturnList(res, condition.Page.Total), nil
}
