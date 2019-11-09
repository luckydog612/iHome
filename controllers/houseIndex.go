package controllers

import "github.com/astaxie/beego"

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController) ReturnData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *HouseIndexController) GetHouseIndex() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	resp["errno"] = "4001"
	resp["errmsg"] = "query failed"
}
