package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "iHome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) ReturnData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserController) Register() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)

	var req = make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp["errno"] = "4001"
		resp["errmsg"] = "register failed"
		return
	}
	//mobile:"111"
	//password:"111"
	//sms_code:"111"
	o := orm.NewOrm()
	user := User{}
	user.Mobile = req["mobile"].(string)
	user.Password_hash = req["password"].(string)
	user.Name = user.Mobile
	id, err := o.InsertOrUpdate(&user)
	if err != nil {
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}
	beego.Info("register success id = ", id)
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)

	c.SetSession("name", user.Name)
}
