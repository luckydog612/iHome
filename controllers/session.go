package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "iHome/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) ReturnData(resp map[string]interface{}) {
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *SessionController) GetSessionData() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	type User struct {
		Name string `json:"name"`
	}
	user := User{}
	resp["errno"] = RECODE_DBERR
	resp["errmsg"] = RecodeText(RECODE_DBERR)

	name := c.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = RECODE_OK
		resp["errmsg"] = RecodeText(RECODE_OK)
		resp["data"] = user
	}
}

func (c *SessionController) DeleteSessionData() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	c.DelSession("name")
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
}

func (c *SessionController) Login() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)

	// 1.获取传回数据
	var req = make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp["errno"] = RECODE_REQERR
		resp["errmsg"] = RecodeText(RECODE_REQERR)
		return
	}
	mobile := req["mobile"].(string)
	password := req["password"].(string)
	// 2.校验数据
	if mobile == "" || password == "" {
		resp["errno"] = RECODE_REQERR
		resp["errmsg"] = RecodeText(RECODE_REQERR)
		return
	}
	user := User{Mobile: mobile}
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("mobile", user.Mobile).One(&user)
	if err != nil {
		beego.Error("login query db err: ", err)
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}

	if user.Password_hash != password {
		resp["errno"] = RECODE_DATAERR
		resp["errmsg"] = RecodeText(RECODE_DATAERR)
		return
	}
	// 3.设置session
	c.SetSession("name", user.Name)
	c.SetSession("mobile", user.Mobile)
	c.SetSession("user_id", user.Id)

	// 4.返回数据
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
}
