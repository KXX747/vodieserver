package controllers

import (
	"testing"
)



func TestCreateUsers(t *testing.T) {

	/**
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8080/user/find", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-auth-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBpZCI6IjM1MDEwMjE5ODIwMzEwNjQ0MSIsImlzcyI6IjM1MDEwMjE5ODIwMzEwNjQ0MSJ9.U-vskEa7M8RpgijWMAI0Fn8FFR0IsW6WLg9wxDFOs4s")
	req.Header.Set("x-auth-version", "1.1.222")
	req.Header.Set("x-auth-uid", "1111")


	//req.Write()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	**/
}


func BenchmarkCreateUser(b *testing.B) {

	/**
	for n:=0;n<1000 ;n++  {

		client := &http.Client{}

		req, err := http.NewRequest("POST", "http://localhost:8080/user", strings.NewReader("name=cjb"))
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

	**/
}