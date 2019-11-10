package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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
