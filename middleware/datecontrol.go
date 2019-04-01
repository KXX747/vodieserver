package middleware

import (
	"time"

)

/**
日期时间格式化
 */

//返回毫秒
func DateUnix() int64{
	return time.Now().Unix()
}

//返回纳秒
func DateUnixNano() int64{
	return time.Now().UnixNano()
}

//把时间格式化成字符串(time->string) : time.Now().Format("2006-01-02 15:04:05")
func DateFormatTimeToStringTime() string{

	return time.Now().Format("2006-01-02 15:04:05")
}






