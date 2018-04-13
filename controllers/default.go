package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"os"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) UploadFile() {
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  fmt.Sprintf("GetFile failed: %v", err),
		}
	}
	path := "/home/asv_demo_file/" + h.Filename
	f.Close()
	this.SaveToFile("file", path)

	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}
	this.ServeJSON(false)
	return
}

func (this *MainController) DownloadFile() {
	filename := this.Ctx.Request.URL.Query().Get("filename")
	this.Ctx.Output.Download("/home/asv_demo_file/" + filename)
	//this.Data["json"] = map[string]interface{}{
	//	"code": 0,
	//	"msg":  "success",
	//}
	//this.ServeJSON(false)
}

func (this *MainController) RemoveFile() {
	filename := this.Ctx.Request.URL.Query().Get("filename")
	os.Remove("/home/asv_demo_file/" + filename)
	this.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}
	this.ServeJSON(false)
}
//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}
