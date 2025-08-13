package consts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_OK     = 200
	PARAMS_ERR  = 2001
	ENCODE_ERR  = 4001
	DECODE_ERR  = 4002
	GET_KEY_ERR = 4000
)

type BaseResp struct {
	ErrCode    int         `json:"errCode"`
	ErrMessage string      `json:"errMessage"`
	Data       interface{} `json:"data"`
}

func BuildRespWithCode(c *gin.Context, code int, resp interface{}) {
	res := &BaseResp{
		ErrCode:    code,
		ErrMessage: "",
		Data:       resp,
	}
	c.AsciiJSON(http.StatusBadRequest, res)
}
func BuildSuccessResp(c *gin.Context, resp interface{}) {
	res := &BaseResp{
		ErrCode: HTTP_OK,
		Data:    resp,
	}
	c.AsciiJSON(http.StatusOK, res)
}
