package controller

import (
	"context"

	"github.com/D-Watson/live-safety/consts"
	"github.com/D-Watson/live-safety/entity"
	"github.com/D-Watson/live-safety/log"
	"github.com/D-Watson/live-safety/services"
	"github.com/gin-gonic/gin"
)

func TransferData(c *gin.Context) {
	req := &entity.TransferRequest{}
	ctx := context.Background()
	resp := &entity.TransferResponse{}
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf(ctx, "[HTTP] analyze req error, err=", err)
		consts.BuildRespWithCode(c, consts.PARAMS_ERR, resp)
		return
	}
	//参数校验
	if !services.VerifyTransferParams(ctx, req) {
		consts.BuildRespWithCode(c, consts.PARAMS_ERR, resp)
		return
	}
	resp, code := services.TransferHttp(ctx, req)
	if code > 0 {
		log.Errorf(ctx, "[Transfer] request error--- ", code)
		consts.BuildRespWithCode(c, code, resp)
		return
	}
	consts.BuildSuccessResp(c, resp)
}
