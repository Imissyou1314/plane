package test

import (
	"path/filepath"
	"plane/app/service"
	"runtime"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func AddUserTest() {
	user, _ := service.UserService.AddUser("missyou", "1255418233@qq.com", "123456", 1)
}

func GetUserTest() {
	user, _ := service.UserService.GetUser(1)
}
