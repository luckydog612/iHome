package controllers

import "github.com/astaxie/beego"

type SessionController struct {
	beego.Controller
}

func (c *SessionController) ReturnData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c SessionController) GetSession() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	resp["errno"] = "4001"
	resp["errmsg"] = "query failed"
}
