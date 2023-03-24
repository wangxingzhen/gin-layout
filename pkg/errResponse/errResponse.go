package errResponse

import (
	"gin-layout/pkg/errors"
)

const ReasonSuccess = "SUCCESS"
const ReasonUnknownError = "UNKNOWN_ERROR"

const ReasonParamsError = "PARAMS_ERROR"
const ReasonUnauthorizedUser = "UNAUTHORIZED_USER"
const ReasonLoginTokenIsExpired = "LOGIN_TOKEN_IS_EXPIRED"
const ReasonLoginPermissionDenied = "LOGIN_PERMISSION_DENIED"
const ReasonUserIsNotFount = "REASON_USER_IS_NOT_FOUNT"
const ReasonDataIsNotFount = "REASON_DATA_IS_NOT_FOUNT"

var reasonMessageAll = map[string]string{
	ReasonSuccess:      "success",
	ReasonUnknownError: "未知错误",

	ReasonParamsError: "请求参数错误",

	ReasonUnauthorizedUser:      "用户未授权",
	ReasonLoginTokenIsExpired:   "登陆信息已失效",
	ReasonLoginPermissionDenied: "无权登陆",
	ReasonUserIsNotFount:        "用户不存在",
	ReasonDataIsNotFount:        "data is not found",
}

var reasonCodeAll = map[string]int{
	ReasonSuccess: 200,

	ReasonUnknownError:          10001,
	ReasonParamsError:           10002,
	ReasonUnauthorizedUser:      10003,
	ReasonLoginTokenIsExpired:   10004,
	ReasonLoginPermissionDenied: 10005,
	ReasonUserIsNotFount:        10006,
	ReasonDataIsNotFount:        10007,
}

//
//// SetCustomizeErrInfo 根据err.Reason返回自定义包装错误
//func SetCustomizeErrInfo(err error) error {
//	e := errors.FromError(err)
//	reason := e.Reason
//	if reason == "" {
//		reason = ReasonUnknownError
//	}
//	if _, ok := reasonCodeAll[reason]; !ok {
//		return err
//	}
//	// 如果是参数错误， 则检查err是否有值， 有则直接返回
//	if e.Reason == ReasonParamsError && e.Message != "" {
//		return errors.New(reasonCodeAll[reason], reason, e.Message)
//	}
//	return SetCustomizeErrInfoByReason(e.Reason)
//}

// SetCustomizeErrInfoByReason 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfoByReason(reason string) error {
	code, message := reasonCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}

// SetCustomizeErrMsgByReason 根据err.Reason返回自定义包装错误
func SetCustomizeErrMsgByReason(reason string, message string) error {
	code := reasonCodeAll[reason]
	return errors.New(code, reason, message)
}

// SetSuccessMsg 返回成功
func SetSuccessMsg() error {
	return SetCustomizeErrInfoByReason(ReasonSuccess)
}
