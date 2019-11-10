package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "iHome/models"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController) ReturnData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AreaController) GetAreas() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)
	// 从数据库中取出数据
	areas := make([]*Area, 0)
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Error("query num = ", num, "err: ", err)
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}
	beego.Info("query num = ", num)
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	resp["data"] = &areas
}
