package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
	"github.com/KXX747/vodieserver/models"
	"encoding/json"
	"io/ioutil"
	"github.com/KXX747/vodieserver/defs"
	"fmt"
)

//创建用户
func CreateUser(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {
	body, _ := ioutil.ReadAll(r.Body)
	var userinfo  models.UserInfo
	json.Unmarshal(body,&userinfo)

	fmt.Println(userinfo)

	err :=userinfo.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	//
	error:=userinfo.CreateUserModles()
	if error!=nil {
		io.WriteString(w,error.Error())
		return
	}

	responseData,_:=json.Marshal(userinfo)
	defs.SendNormalResponse(w,string(responseData),200)

}

//查询用户
func FindUser(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {
	body, _ := ioutil.ReadAll(r.Body)
	var findUserParmas  models.ReuqestUserParmas
	json.Unmarshal(body,&findUserParmas)

	err :=findUserParmas.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	userInfo,error:=findUserParmas.FindUser()
	if error!=nil {
		io.WriteString(w,error.Error())
		return
	}

	responseData,_:=json.Marshal(userInfo)
	defs.SendNormalResponse(w,string(responseData),200)

}


//修改用户信息
func UpdateUser(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {
	body, _ := ioutil.ReadAll(r.Body)
	var userinfo  models.UserInfo
	json.Unmarshal(body,&userinfo)

	err :=userinfo.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	error:=userinfo.UpdateUser()
	if error!=nil {
		io.WriteString(w,error.Error())
		return
	}
	responseData,_:=json.Marshal(userinfo)
	defs.SendNormalResponse(w,string(responseData),200)

}

//删除用户信息
func DeleteUser(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {
	body, _ := ioutil.ReadAll(r.Body)
	var deleteUserParmas  models.ReuqestUserParmas
	json.Unmarshal(body,&deleteUserParmas)

	err :=deleteUserParmas.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	error:=deleteUserParmas.DeleteUser()
	if error!=nil {
		io.WriteString(w,error.Error())
		return
	}

	defs.SendNormalResponse(w,deleteUserParmas.UserName+"删除成功",200)

}