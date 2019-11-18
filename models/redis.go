package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var RedisCon cache.Cache

func init() {
	var err error
	RedisCon, err = cache.NewCache("redis", `{"key":"ihome","conn":":6379","dbNum":"0","password":"123456"}`)
	if err != nil {
		beego.Error("get redis connection failed, err: ", err)
	}
}
