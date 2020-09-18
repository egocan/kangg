package common

import (
	"github.com/astaxie/beego"
	"math/rand"
	"path"
	"strconv"
	"time"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Uploads() {

	response := make(map[string]interface{})

	filename := c.GetString("filename")
	f, h, err := c.GetFile("editormd-image-file")
	defer f.Close()
	if err != nil {
		response["message"] = err.Error()
		response["success"] =  0
	}
	if len(filename) ==0 {

		ext := path.Ext(h.Filename)
		filename = time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000)) + ext

		err := c.SaveToFile("editormd-image-file", "static/uploads/"+filename)
		if err != nil {
			response["message"] = err.Error()
			response["success"] =  0
		}else {
			response["success"] =  1
			response["message"] = "Success."
			response["url"] = "/static/uploads/" + filename
		}
	}else{
		response["success"] =  1
		response["message"] = "Success."
		response["url"] = filename
	}

	c.Data["json"] = response
	c.ServeJSON()

}