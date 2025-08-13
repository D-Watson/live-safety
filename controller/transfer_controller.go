package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"live_safty/consts"
	"live_safty/entity"
	"live_safty/log"
	"live_safty/services"
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
	var err error
	resp, err = services.TransferHttp(ctx, req)
	if err != nil {
		log.Errorf(ctx, "[Transfer] request error--- ", err)
		return
	}
	consts.BuildSuccessResp(c, resp)
}
