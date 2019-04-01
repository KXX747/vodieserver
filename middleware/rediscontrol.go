package middleware

import (
	"github.com/gomodule/redigo/redis"

	"time"
	"fmt"
)

/**
reids 连接相关工具
 */

const redisValueIsNotExist  ="redigo: nil returned"

 var connPool *redis.Pool

func RedisConn() redis.Conn {

	return connPool.Get()
}

func init() {
	server := "127.0.0.1:6379"
	password := ""

	connPool = newPool(server, password)
}


func newPool(server, password string) *redis.Pool {
	return &redis.Pool{

		IdleTimeout:240*time.Second,
		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",server)
			if err!=nil {
				return  nil,err
			}

			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}



/**
	k/v
 */
func RedisSet(key string ,value []byte) (error) {
	c, err := connPool.Dial()
	if err != nil {
		return err
	}
	_, err = c.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil

}

/**
 get k value
 */
func RedisGet(key string)(value []byte,err error) {
	c, err := connPool.Dial()
	if err != nil {
		return nil, err
	}
	val, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		//
		if redisValueIsNotExist ==  err.Error(){
			return nil, nil

		}else {
			return nil, err
		}
	} else {
		return val, nil
	}
}



/**
 根据key -修改对应的value值
 */
func RedisRemoveValue(key string) error {

	c, err := connPool.Dial()
	if err != nil {
		return err
	}

	_,err =c.Do("DEL", key)
	if err != nil {
		return err
	}

	return nil
}


/**
// 设置过期时间为24小时
  24*3600
   2*365*24*3600
 */
func RedisExpipeValue (key string) error{
	c, err := connPool.Dial()
	if err != nil {
		return err
	}

	n,err:=c.Do("EXPIRE", key, 24*3600)
	if err != nil {
		return err
	}

	fmt.Println("RedisExpipeValue  ",n)

	return nil
}


/**
// 设置过期时间为24小时
 */
func RedisExpipeValueToTime (key string,expipeTime int64) error{
	c, err := connPool.Dial()
	if err != nil {
		return err
	}

	_,errs :=c.Do("EXPIRE", key, expipeTime)
	if errs != nil {
		return errs
	}
	return nil
}

/**
删除指定的key
 */
func RedisDelKey(key string) error {
	c, err := connPool.Dial()
	if err != nil {
		return err
	}
	_,errs :=c.Do("DEL", key)
	if errs != nil {
		return errs
	}
	return nil
}

