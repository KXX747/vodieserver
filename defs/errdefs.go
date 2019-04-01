package defs

import "net/http"

var(
	//参数异常
	ErrorReuqestBodyParseFailed =ErrResponse{HttpCode:400,Error:Err{Error:"request not connact",ErrorCode:"0001"}}
	//api接口未权限认证
	ErrorNotAuthUser =ErrResponse{HttpCode:401,Error:Err{Error:"user auth failed",ErrorCode:"0002"}}

	//mysql连接失败
	ErrorDB =ErrResponse{HttpCode:401,Error:Err{Error:"connect db failed",ErrorCode:"0003"}}

	ErrorRedis =ErrResponse{HttpCode:401,Error:Err{Error:"connect Redis failed",ErrorCode:"0004"}}

	//流量控制
	ErrorLimitRate =ErrResponse{HttpCode:http.StatusTooManyRequests,Error:Err{Error:"Status Too Many Requests",ErrorCode:"0005"}}

)


type Err struct {
    Error string `json:"error"`
    ErrorCode string `json:"error_code"`

}

type ErrResponse struct {
	HttpCode int `json:"http_code"`
	Error Err
}

