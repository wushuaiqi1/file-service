package utils

import (
	"strconv"
	"time"
)

func SystemCurrentMills() int64 {
	return time.Now().UnixMilli()
}

func Int64ToString(param int64) string {
	return strconv.FormatInt(param, 10)
}
