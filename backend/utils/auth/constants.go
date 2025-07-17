package auth

import (
	"time"
)

var REFRESH_TOKEN_EXPIRED_TIME = time.Now().Add(7 * 24 * time.Hour).Unix()
var ACCESS_TOKEN_EXPIRED_TIME = time.Now().Add(15 * time.Minute).Unix()
