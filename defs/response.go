package defs

import (
	"net/http"
	"io"
	"encoding/json"
)


func SendErrorResponse(w http.ResponseWriter,errResp ErrResponse){

	w.WriteHeader(errResp.HttpCode)

	resStr,_:=json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))
}

func SendNormalResponse(w http.ResponseWriter,resp string , code int)  {

	w.WriteHeader(code)
	io.WriteString(w,resp)

}