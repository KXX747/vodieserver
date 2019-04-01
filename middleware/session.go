package middleware

import "sync"

/**
session 缓存用户的信息和redis类似
 */

 type SimpleSession struct {
 	UserName string
 	TTL int64
 }


 var sessionMap *sync.Map
