package ginx

import (
	"gin-layout/pkg/errResponse"
	"gin-layout/pkg/errors"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"net/http"
)

// RequestContext 封装gin.Context, 提供更加便捷的方法来处理各种参数等
type RequestContext struct {
	Context *gin.Context
	Request *http.Request
	UserId  uint64
	logger  *logs.Entry
}

// New 从gin.Context构建RequestContext
func New(c *gin.Context) *RequestContext {
	return &RequestContext{
		Context: c,
		Request: c.Request,
		UserId:  c.GetUint64("login_user_id"),
		logger:  c.MustGet("logger").(*logs.Entry),
	}
}

// ErrResponse 返回错误信息
func (rc *RequestContext) ErrResponse(err *errors.Error) {
	rc.Context.JSON(http.StatusOK, gin.H{
		"code": errors.Code(err),
		"msg":  errors.Message(err),
	})
}

// ToResponse 返回数据
func (rc *RequestContext) ToResponse(data any) {
	err := errResponse.SetSuccessMsg()
	rc.Context.JSON(200, gin.H{
		"code": errors.Code(err),
		"msg":  errors.Message(err),
		"data": data,
	})
}

// SuccResponse 成功无数据返回
func (rc *RequestContext) SuccResponse() {
	err := errResponse.SetSuccessMsg()
	rc.Context.JSON(http.StatusOK, gin.H{
		"code": errors.Code(err),
		"msg":  errors.Message(err),
	})
}

// GetLogger return logger
func (rc *RequestContext) GetLogger() *logs.Entry {
	return rc.logger
}

// ReturnList 分页返回格式化数据
func (rc *RequestContext) ReturnList(list any, Total int64) any {
	return struct {
		List  any   `json:"list"`
		Total int64 `json:"total"`
	}{
		List:  list,
		Total: Total,
	}
}
