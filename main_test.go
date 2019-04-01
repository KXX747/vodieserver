package main

import (
	"testing"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

//testing.M test的初始化主函数 需要m.Run()
func TestMain(m *testing.M)  {
	fmt.Println("test main ")
	m.Run()
}


//Test  go test /go test-v
func TestPrint(t *testing.T)  {
	userName :=GetUserName()
	fmt.Println(userName)

	if "2"!=userName {
		t.Error("error")
	}
}

func TestPrintT(t *testing.T)  {
	t.SkipNow()
	fmt.Println("TestPrintT ok")
}

func TestRequestPostCheck(t *testing.T)  {


	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8080/check", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-auth-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6IjM1MDEwMjE5ODIwMzEwNjQ0MSIsImlzcyI6IjM1MDEwMjE5ODIwMzEwNjQ0MSJ9.U-vskEa7M8RpgijWMAI0Fn8FFR0IsW6WLg9wxDFOs4s")
	req.Header.Set("x-auth-version", "1.1.222")
	req.Header.Set("x-auth-uid", "1111")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}



//test Run运行子函数
func TestSubtestsPrint(t *testing.T)  {

	t.Run("a1", func(t *testing.T) {
		fmt.Println("a1")
	})

	t.Run("TestPrintT",TestPrintT)

}


//benchmark函数一般以benchmark开头
//bencgmark的case一般会跑b.N次，而且每次执行都是如此
//run =  go test -bench=.
func BenchmarkGetUserName(b *testing.B) {
	for n:=0;n<b.N ; n++ {
		userName :=GetUserName()
		fmt.Println(n,"  ",userName)
	}
}


//	//
func BenchmarkRequestPostCheck(b *testing.B){
	for n:=0;n<b.N ;n++  {
		client := &http.Client{}

		req, err := http.NewRequest("POST", "http://localhost:8080/check", strings.NewReader("name=cjb"))
		if err != nil {
			// handle error
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("x-auth-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6IjM1MDEwMjE5ODIwMzEwNjQ0MSIsImlzcyI6IjM1MDEwMjE5ODIwMzEwNjQ0MSJ9.U-vskEa7M8RpgijWMAI0Fn8FFR0IsW6WLg9wxDFOs4s")
		req.Header.Set("x-auth-version", "1.1.222")
		req.Header.Set("x-auth-uid", "1111")

		resp, err := client.Do(req)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		fmt.Println(string(body))
	}
}




