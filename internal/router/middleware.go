package router

import (
	"gin-layout/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	logs "github.com/sirupsen/logrus"
)

// VerifyLogin 登陆校验
func VerifyLogin(user *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// example ... start
		//
		//resErr := errResponse.SetCustomizeErrInfoByReason(errResponse.ReasonLoginTokenIsExpired)
		//
		//if debug {
		//	c.Set("login_user_id", uint64(9999))
		//	c.Next()
		//	return
		//}
		//if {No login} {
		//	rc := ginx.New(c)
		//	re := errors.FromError(resErr)
		//	rc.ErrResponse(re)
		//	c.Abort()
		//	return
		//} else {
		//	c.Set("login_user_id", uint64(xxx))
		//}
		//
		// example ... end
		c.Next()
	}
}

// GenLogger generate request logger
func GenLogger(logger *logs.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger.WithFields(logs.Fields{
			"request_id": uuid.NewString(),
		}))
		c.Next()
	}
}

// GenGinLogger change gin.DefaultWriter logger
func GenGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := gin.LoggerWithWriter(c.MustGet("logger").(*logs.Entry).WriterLevel(logs.InfoLevel))
		writer(c)
		c.Next()
	}
}

// GenGinRecover gin recover
func GenGinRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := gin.RecoveryWithWriter(c.MustGet("logger").(*logs.Entry).WriterLevel(logs.ErrorLevel))
		r(c)
		c.Next()
	}
}
