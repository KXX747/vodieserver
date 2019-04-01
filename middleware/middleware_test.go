package middleware

import (
	"testing"
	"fmt"
)

func TestMain(m *testing.M)  {

	m.Run()
}

//创建uuid
func TestCreateUuid(t *testing.T) {

	uuid:=CreateUuid()

	fmt.Println(uuid.String())
}

//redis 存string数据
func TestRedisSet(t *testing.T) {

	error:=RedisSet("zhangsan",[]byte("bog"))

	if error!=nil {

		fmt.Println(error.Error())
	}else {
		fmt.Println("添加成功")
	}

}

//redis 获取string
func TestRedisGet(t *testing.T) {
	value,err:=RedisGet("zhangsan")

	if err!=nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(value))
}