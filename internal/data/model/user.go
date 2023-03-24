package model

import "gin-layout/internal/biz"

type UcUser struct {
	Model
	Name         string
	SerialNumber int
}

func (u *UcUser) TableName() string {
	return "uc_users"
}

func (u *UcUser) ToDomain() *biz.UcUser {
	return &biz.UcUser{
		Id:           u.ID,
		Name:         u.Name,
		SerialNumber: u.SerialNumber,
	}
}
