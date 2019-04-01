package models

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestUserInfo_CreateUserModles(t *testing.T) {

	userInfo:=UserInfo{0,"张三","xiaoer","2017-07-10 13:21:22","0"}

	b ,_:=json.Marshal(userInfo)



	var userinfos UserInfo

	json.Unmarshal(b,&userinfos)

	fmt.Println(userinfos)
}