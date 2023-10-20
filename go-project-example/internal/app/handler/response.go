package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code uint16

const (
	Success            Code = 0
	NoContent          Code = 2040
	ParseFieldsError   Code = 4000
	AuthenticateFailed Code = 4010
	InternalError      Code = 5000
)

var codeMsgMap = map[Code]string{
	Success:            "success",
	NoContent:          "same data, skipped",
	ParseFieldsError:   "parse fields error",
	AuthenticateFailed: "authenticate failed",
	InternalError:      "internal error",
}

type event struct {
	c    *gin.Context
	msg  string
	data interface{}
}

func Resp(c *gin.Context) *event {
	return &event{
		c: c,
	}
}

func (e *event) SetData(i interface{}) *event {
	(*e).data = i
	return e
}

func (e *event) SetMsg(msg string) *event {
	(*e).msg = msg
	return e
}
func (e *event) Abort(c Code) {
	if (*e).msg == "" {
		(*e).msg = codeMsgMap[c]
	}
	if (*e).data == nil {
		(*e).data = gin.H{}
	}
	(*e).c.Set("code", uint(c))
	(*e).c.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"code": c,
			"msg":  (*e).msg,
			"data": (*e).data,
		},
	)
}
