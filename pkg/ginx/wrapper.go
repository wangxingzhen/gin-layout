package ginx

import (
	"fmt"
	"gin-layout/pkg"
	"gin-layout/pkg/errors"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"strings"
)

// RequestHandler api层看到的handler
type RequestHandler func(*RequestContext)

// RequestFilter 在RequestHandler做一层filter
type RequestFilter func(RequestHandler) RequestHandler

// ReturnJSON writes json to response
func ReturnJSON(c *RequestContext, data any, err error) {
	if err != nil {
		se := errors.FromError(err)
		c.ErrResponse(se)
		return
	}

	if data == nil {
		c.SuccResponse()
		return
	}
	c.ToResponse(data)
	return
}

// api is a base wrapper which makes metrics, initializes notice logger and packs JSON response.
func api(request APIHandler, serviceName string, funcName string) RequestHandler {
	return func(rc *RequestContext) {
		defer func() {
			if e := recover(); e != nil {
				panic(e)
			}
		}()
		response, err := request(rc)
		if err != nil {
			rc.GetLogger().Errorf(fmt.Sprintf("%+v", err))
		}

		ReturnJSON(rc, response, err)
	}
}

// APIHandler 正常情况下的API的格式
type APIHandler func(*RequestContext) (any, error)

// API 将 APIRequestHandler 和 filters封装成为 gin#handler
func API(f APIHandler, filters ...RequestFilter) gin.HandlerFunc {
	// 通过反射获取函数名称
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	fullNameStr := strings.Split(fullName, ".")
	funcName := fullNameStr[len(fullNameStr)-1]
	ServiceName := fullNameStr[len(fullNameStr)-2]

	handler := api(f, ServiceName, funcName)
	for i := range filters {
		w := filters[len(filters)-1-i]
		handler = w(handler)
	}

	return func(c *gin.Context) {
		rc := New(c)
		rc.GetLogger().
			WithField(fmt.Sprintf("%sHeader", funcName), pkg.ToJSON(c.Request.Header)).
			Info()
		handler(rc)
	}
}
