package validate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// ValidateStruct 验证数据
func ValidateStruct(model any) error {
	//验证
	validate := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	err := validate.Struct(model)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return errors.New(err.Error())
		}
	}
	return nil
}
