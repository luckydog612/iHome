package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "iHome/models"
	"path"
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
	c.SetSession("mobile", user.Mobile)
	c.SetSession("user_id", id)
}

func (c *UserController) UpdateAvatar() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)

	f, h, err := c.GetFile("avatar")
	if err != nil {
		beego.Error("get avatar file err ", err)
		resp["errno"] = RECODE_REQERR
		resp["errmsg"] = RecodeText(RECODE_REQERR)
		return
	}

	suffix := path.Ext(h.Filename)
	fmt.Println(h.Filename)
	fmt.Println(suffix)
	if suffix != ".gif" && suffix != ".jpg" && suffix != ".png" && suffix != ".jpeg" {
		beego.Error("file format err: ", suffix, " gif jpg png jpeg requested")
		resp["errno"] = RECODE_REQERR
		resp["errmsg"] = RecodeText(RECODE_REQERR)
		return
	}

	buffer := make([]byte, h.Size)
	num, err := f.Read(buffer)
	if err != nil {
		beego.Error("出错啦！", num, err)
	}
	// 暂时不转为哈希
	//avatarName := utils.HashName(buffer)
	//fmt.Println(string(avatarName))
	//fmt.Println(avatarName)
	defer f.Close()
	err = c.SaveToFile("avatar", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if err != nil {
		beego.Error("save avatar file err ", err)
		resp["errno"] = RECODE_REQERR
		resp["errmsg"] = RecodeText(RECODE_REQERR)
		return
	}
	user := User{}
	user_id := c.GetSession("user_id")
	fmt.Println("user_id", user_id)
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("id", user_id).One(&user)
	if err != nil {
		beego.Error("查无此人", user_id)
	}
	avatar_url := fmt.Sprintf("127.0.0.1:8080/static/upload/%s", h.Filename)
	user.Avatar_url = avatar_url
	col, err := o.Update(&user)
	if err != nil {
		beego.Error("col: ", col, "err: ", err)
	}
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	repData := make(map[string]string)
	repData["avatar_url"] = avatar_url
	resp["data"] = repData
}

func (c *UserController) GetUserData() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)

	// 1.从session中获取用户的user_id
	user_id := c.GetSession("user_id")
	user := User{Id: user_id.(int)}
	// 2.通过user_id获取数据库中的user信息
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	resp["data"] = user
}

func (c *UserController) UpdateUserName() {
	var resp = make(map[string]interface{})
	defer c.ReturnData(resp)

	// 1.获取session中的user_id
	user_id := c.GetSession("user_id")
	// 2. 获取前端传过来的消息
	reqBody := make(map[string]string)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody)
	if err != nil {
		beego.Error("unmarshal userName err: ", err)
		resp["errno"] = RECODE_DATAERR
		resp["errmsg"] = RecodeText(RECODE_DATAERR)
		return
	}
	// 3. 更新user_id对应的name
	userName := reqBody["name"]
	o := orm.NewOrm()
	user := User{}
	user.Name = userName
	user.Id = int(user_id.(int64))
	err = o.QueryTable("user").Filter("id", user.Id).One(&user)
	if err != nil {
		beego.Error("query user err: ", err)
		resp["errno"] = RECODE_DBERR
		resp["errmsg"] = RecodeText(RECODE_DBERR)
		return
	}
	user.Name = userName
	col, err := o.Update(&user)
	if err != nil || col <= 0 {
		beego.Error("update userName to db err: ", err)
	}
	// 4. 把session中的name更新
	c.SetSession("name", userName)
	// 5. 将数据返回前端
	resp["errno"] = RECODE_OK
	resp["errmsg"] = RecodeText(RECODE_OK)
	resp["data"] = reqBody
}
