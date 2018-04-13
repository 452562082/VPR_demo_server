package routers

import (
	"asv_demo_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/uploadFile", &controllers.MainController{}, "post:UploadFile")
	beego.Router("/downloadFile", &controllers.MainController{}, "get:DownloadFile")
	beego.Router("/removeFile", &controllers.MainController{}, "get:RemoveFile")
	//beego.Router("/download", &controllers.MainController{}, "get:Download")
}
