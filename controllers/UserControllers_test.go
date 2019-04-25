package controllers

import (
	"testing"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/KXX747/vodieserver/models"
	"time"
	"encoding/json"
	"strconv"
)


var (
	requestCount = 1000
)

var HEADER_FIELD_SESSION ="X-SESSSION-ID"

func TestMain(m *testing.M){

	m.Run()
}

func BenchmarkUserWorkFlow(b *testing.B) {
	for n:=0;n<1000 ; n++ {
		fmt.Println("         ")
		fmt.Println("         ")
		createUser(n)
		findUser(n)
		updateUser(n)
		findUser(n)
		deleteUser(n)
		findUser(n)
	}

}


//
func TestUserWorkFlow(t *testing.T)  {
	t.Run("create",TestCreateUser)
	t.Run("find",TestFindUser)
	t.Run("update",TestUpdateUser)
	t.Run("delete",TestDeleteUser)
}

//创建用户
func TestCreateUser(t *testing.T) {
	createUser(0)

}

//查询用户
func TestFindUser(t *testing.T) {

	findUser(0)
}

//更新用户
func TestUpdateUser(t *testing.T) {

	updateUser(0)
}

//删除用户
func TestDeleteUser(t *testing.T) {

		deleteUser(0)
}

func createUser(position int)  {

	userInfo:=models.UserInfo{}
	userInfo.Id = 0
	UserName:=fmt.Sprintf("%s,%s", "shanghai",  strconv.Itoa(position))
	userInfo.UserName =UserName
	userInfo.Pwd= "888888"
	userInfo.CreateUserDate = time.UTC.String()
	userInfo.IsVip="0"
	b,_:=json.Marshal(userInfo)

	fmt.Println(string(b))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/user/create", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))


}


func findUser(position int){
	client := &http.Client{}
	findUser:=models.ReuqestUserParmas{}
	UserName:=fmt.Sprintf("%s,%s", "shanghai",  strconv.Itoa(position))
	findUser.UserName = UserName
	b ,_:=json.Marshal(findUser)
	req, err := http.NewRequest("POST", "http://localhost:18080/user/find", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))
}

func updateUser(position int)  {

	userInfo:=models.UserInfo{}
	userInfo.Id = 0
	UserName:=fmt.Sprintf("%s,%s", "shanghai",  strconv.Itoa(position))
	userInfo.UserName = UserName
	userInfo.Pwd= "888888"
	userInfo.CreateUserDate = time.UTC.String()
	userInfo.IsVip="1"
	b,_:=json.Marshal(userInfo)

	fmt.Println(string(b))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/user/update", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION, userInfo.UserName)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))


}

func deleteUser(position int)  {

	client := &http.Client{}
	findUser:=models.ReuqestUserParmas{}
	UserName:=fmt.Sprintf("%s,%s", "shanghai",  strconv.Itoa(position))
	findUser.UserName = UserName
	b ,_:=json.Marshal(findUser)
	req, err := http.NewRequest("POST", "http://localhost:8080/user/delete", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION, findUser.UserName)


	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))
}