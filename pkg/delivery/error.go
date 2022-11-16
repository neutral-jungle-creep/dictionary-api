package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Exception struct {
	funcName string
	ctx      *gin.Context
	httpSt   int
	err      error
}

func NewException(funcName string, ctx *gin.Context, httpSt int, err error) *Exception {
	return &Exception{
		funcName: funcName,
		ctx:      ctx,
		httpSt:   httpSt,
		err:      err,
	}
}

func (e *Exception) ExceptionResp() {
	logrus.Errorf("error in %s func, status: %d, error: %s", e.funcName, e.httpSt, e.err.Error())
	e.ctx.AbortWithStatusJSON(e.httpSt, e.err.Error())
}
