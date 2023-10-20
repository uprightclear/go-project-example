package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-example/internal/app/type/website"
	website2 "go-project-example/internal/app/website"
)

func TestHandler(c *gin.Context) {
	var err error
	var req website.TestReq
	if err = c.ShouldBindJSON(&req); err != nil {
		Resp(c).SetMsg(err.Error()).Abort(ParseFieldsError)
		return
	}

	var resp website.TestResp
	resp, err = website2.Test(req)
	if err != nil {
		Resp(c).SetMsg(err.Error()).Abort(InternalError)
		return
	}

	Resp(c).SetData(resp).Abort(Success)
	return
}
