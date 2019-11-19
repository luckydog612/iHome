package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "iHome/models"
	"iHome/utils"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController) ReturnData(resp map[string]interface{}) {
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AreaController) GetAreas() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	// 如果在缓存中
	rAreas := RedisCon.Get("areas")
	if rAreas != nil {
		beego.Info("have got areas info from redis", rAreas)
		//data,_ := json.Marshal(string(rAreas.([]uint8)))
		resp["errno"] = RECODE_OK
		resp["errmsg"] = RecodeText(RECODE_OK)
		resp["data"] = utils.StringToAreas(string(rAreas.([]uint8)))
		return
	}
	// 从数据库中取出数据
	areas := make([]Area, 0)
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Error("query num = ", num, "err: ", err)
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}
	// 存储到redis中
	areaData, _ := json.Marshal(areas)
	err = RedisCon.Put("areas", areaData, time.Second*3600)
	if err != nil {
		beego.Error("put areas info to redis failed, err: ", err)
	}
	beego.Info("query num = ", num)
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	resp["data"] = areas
}
