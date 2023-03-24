package copierx

import (
	"github.com/golang-module/carbon"
	"github.com/jinzhu/copier"
)

var (
	CarbonToString = []copier.TypeConverter{
		{
			SrcType: carbon.DateTime{},
			DstType: copier.String,
			Fn:      carbonToString,
		},
		{
			SrcType: carbon.Date{},
			DstType: copier.String,
			Fn:      carbonToString,
		},
	}
	StringToCarbon = []copier.TypeConverter{
		{
			SrcType: copier.String,
			DstType: carbon.DateTime{},
			Fn:      stringToCarbonTime,
		},
		{
			SrcType: copier.String,
			DstType: carbon.Date{},
			Fn:      stringToCarbonDate,
		},
	}
)

func carbonToString(src any) (rp any, err error) {
	rp = ""
	if v, ok := src.(carbon.DateTime); ok {
		if !v.IsZero() {
			rp = v.ToDateTimeString()
		}
	}
	if v, ok := src.(carbon.Date); ok {
		if !v.IsZero() {
			rp = v.ToDateString()
		}
	}
	return
}

func stringToCarbonTime(src any) (rp any, err error) {
	if v, ok := src.(string); ok {
		rp = carbon.DateTime{
			Carbon: carbon.Parse(v),
		}
	}
	return
}

func stringToCarbonDate(src any) (rp any, err error) {
	if v, ok := src.(string); ok {
		rp = carbon.Date{
			Carbon: carbon.Parse(v),
		}
	}
	return
}

// Copy copy两个变量 https://github.com/jinzhu/copier/blob/master/README.md
//
//	type A struct {
//		Name string
//	}
//
//	type B struct {
//		Name string
//	}
//
// a:=A{Name:"sss"}
// var b B{}
// err := Copy(&b,a) // 这时b内容：{Name:"sss"}
func Copy(to any, from any) (err error) {
	return copier.CopyWithOption(to, from, copier.Option{Converters: append(CarbonToString, StringToCarbon...)})
}

func CopyWithOption(to any, from any, opt copier.Option) (err error) {
	opt.Converters = append(CarbonToString, StringToCarbon...)
	return copier.CopyWithOption(to, from, opt)
}
