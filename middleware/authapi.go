package middleware

import (
	"net/http"
)

/**
api接口请求校验
 */


 var HEADER_FIELD_SESSION ="X-SESSSION-ID"


//校验拦截api-auth
func ValidDateApi(r *http.Request) bool{

	if r.URL.Path == "/user/create"||r.URL.Path =="/user/find" {
		return true
	}

	//sid := r.Header.Get(HEADER_FIELD_SESSION)
	//if len(sid)==0 {
	//	return false
	//}
	//value,err:=RedisGet(sid)
	//if err!=nil||len(value)==0 {
	//	return false
	//}

	return true
}
