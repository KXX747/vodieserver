package models

import (
	"errors"
	"github.com/KXX747/vodieserver/middleware"
	"encoding/json"
	"github.com/KXX747/vodieserver/defs"
	"fmt"
)

//数据处理层


//创建用户
func(u *UserInfo) CreateUserModles() error {
	//
	 findUserParmas :=&ReuqestUserParmas{u.UserName}
	userInfo,error:=findUserParmas.FindUser()
	if error!=nil {
		return error
	}

	if userInfo.UserName !="" {
		return errors.New("请选择其他用户名称")
	}

	fmt.Println("engineConn.Insert  ",u)
	_, error=engineConn.Insert(u)
	if error!=nil {
		return errors.New(defs.ErrorDB.Error.Error)
	}

	v,_:=json.Marshal(u)

	fmt.Println("usercreate  RedisSet  name =",userInfo.UserName," <><>V<><> ",v)
	error=middleware.RedisSet(u.UserName,v)
	if error!=nil {
		return errors.New(defs.ErrorRedis.Error.Error)
	}
	return nil
}

//查询用户
func(u *ReuqestUserParmas)  FindUser() (UserInfo,error) {
	//
	var userInfo UserInfo
	fmt.Println("userInfo = ",userInfo)
	_,err:=engineConn.Where("user_name=?",u.UserName).Get(&userInfo)
	if err!=nil {
		return userInfo,err
	}

	err =userInfo.CheckValidator()
	if err!=nil {
		return userInfo,err
	}

	return userInfo,nil
}

//修改用户信息
//开启事务提交数据
func(u *UserInfo)  UpdateUser() error {
	//
	session := engineConn.NewSession()

	// add Begin() before any
	err := session.Begin()

	_,err=engineConn.Where("user_name=?",u.UserName).Update(u)
	if err!=nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

//删除用户信息
func(u *ReuqestUserParmas)  DeleteUser() error {

	_,err:=engineConn.Delete(&UserInfo{UserName:u.UserName})

	if err!=nil {
		return err
	}
	return nil
}