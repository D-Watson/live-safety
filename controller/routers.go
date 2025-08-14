package controller

import (
	"context"
	"net/http"

	"github.com/D-Watson/live-safety/conf"
	"github.com/D-Watson/live-safety/log"
	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func InitRouters() {
	c := gin.Default()
	c.POST("/live/safety/transfer", TransferData)
	c.Use(cors())
	err := c.Run(conf.GlobalConfig.Server.Http.Host)
	if err != nil {
		log.Errorf(context.Background(), "[http]run server error--", err)
		return
	}
}
