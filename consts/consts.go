package consts

import (
	"time"
)

const (
	LIVE_FRONTEND_REQ = 1
	LIVE_BACKEND_REQ  = 2
)
const (
	LIVE_TOKEN_FRONT_END = "live:token:frontend"
	LIVE_TOKEN_BACK_END  = "live:token:backend"
)

const (
	LIVE_ENCRYPT = 1
	LIVE_DECRYPT = 2
)

const (
	BIT_SIZE    = 2048
	EXPIRE_TIME = 35 * time.Minute
)
