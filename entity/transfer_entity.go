package entity

type TransferRequest struct {
	Crypto       int    `json:"crypto" binding:"required"`       //1=加密，2=解密
	Role         int32  `json:"role" binding:"required"`         //1=前端 2=服务端
	TransferData string `json:"transferData" binding:"required"` //原始数据
}

type TransferResponse struct {
	TransferData string `json:"transferData"` //处理后的数据
}
