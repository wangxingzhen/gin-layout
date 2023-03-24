package data

import (
	"context"
	"fmt"
	"gin-layout/internal/biz"
	"gin-layout/internal/data/model"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

type ucUserRepo struct {
	data *Data
	sg   *singleflight.Group
}

func NewUcUserRepo(data *Data) biz.IUcUserRepo {
	return &ucUserRepo{
		data: data,
		sg:   &singleflight.Group{},
	}
}

func (r *ucUserRepo) CreateUcUser(ctx context.Context, user *biz.UcUser) error {
	var u model.UcUser
	u.Name = user.Name

	if err := errors.WithStack(r.data.DB(ctx).Save(&u).Error); err != nil {
		return err
	}
	return nil
}

func (r *ucUserRepo) GetUcUserById(ctx context.Context, id uint64) (*biz.UcUser, error) {
	res, err, _ := r.sg.Do(fmt.Sprintf("GetUcUserById_%d", id), func() (any, error) {
		var c model.UcUser
		if err := errors.WithStack(r.data.DB(ctx).Where("id = ?", id).Take(&c).Error); err != nil {
			return nil, err
		}
		return &c, nil
	})
	if err != nil {
		return nil, err
	}
	return res.(*model.UcUser).ToDomain(), nil
}

func (r *ucUserRepo) GetUcUserNum(ctx context.Context) (int, error) {
	var num int64
	err := errors.WithStack(r.data.DB(ctx).Model(&model.UcUser{}).Count(&num).Error)
	return int(num), err
}

func (r *ucUserRepo) SaveUcUserSerialNumber(ctx context.Context, user *biz.UcUser) error {
	return errors.WithStack(r.data.DB(ctx).Model(&model.UcUser{}).
		Where("id = ?", user.Id).
		Update("serial_number", user.SerialNumber).
		Error)
}

func (r *ucUserRepo) GetUcUserMaxId(ctx context.Context) (uint64, error) {
	var user model.UcUser
	err := errors.WithStack(r.data.DB(ctx).Model(&model.UcUser{}).
		Order("id DESC").Select("id").
		Take(&user).Error)
	return user.ID, err
}

func (r *ucUserRepo) UserList(ctx context.Context, condition *biz.ListTestRep) (res []*biz.UcUser, err error) {
	roles := make([]*model.UcUser, 0)
	db := r.data.DB(ctx)
	db = db.Model(&model.UcUser{}).Order("id DESC")
	// 条件筛选
	if condition.Id > 0 {
		db = db.Where("id = ?", condition.Id)
	}
	if len(condition.Name) > 0 {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", condition.Name))
	}

	err = errors.WithStack(condition.Page.
		WithContext(ctx).
		Query(db).
		Find(&roles),
	)
	if err != nil {
		return
	}
	res = make([]*biz.UcUser, 0)
	for _, v := range roles {
		res = append(res, v.ToDomain())
	}
	return
}
