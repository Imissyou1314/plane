package service

import (
	"fmt"
	"net/url"
	"plane/app/models"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer //数据库Orm
	// log 						log.newLog()     // 日志管理器
	tablePrefix     string           //表前缀
	LogService      *logService      //日志服务
	UserService     *userService     //用户服务
	PlaneService    *planeService    //飞机服务
	PlaneLogService *planeLogService // 飞机日记服务
	MessageService  *messageService  //消息服务
)

// ServiceIF 接口
type ServiceIF interface {
	FindOneById(ID int64)
	FindAll()
	DeteleById(ID int64)
	Add()
	UpdataById(ID int64)
}

func init() {

	log.SetFormatter(&log.JSONFormatter{})

	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	tablePrefix = beego.AppConfig.String("db.prefix")

	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset-utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	// orm.RegisterDriver("mysql", orm.mysql)
	log.Info(dsn)
	orm.RegisterDataBase(`default`, "mysql", "root:root@tcp(127.0.0.1:3306)/plane?charset=utf8")
	orm.RegisterDataBase(`miss`, "mysql", dsn)
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterModelWithPrefix(tablePrefix,
		new(models.Log),
		new(models.User),
		new(models.Message),
		new(models.Plane),
		new(models.PlaneLog),
	)

	//自动创建表（建后关闭)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	// 初始化服务对象
	initService()
}

// 初始化服务
func initService() {
	UserService = &userService{}
	LogService = &logService{}
	PlaneLogService = &planeLogService{}
	PlaneService = &planeService{}
	MessageService = &messageService{}
}

// tableName 获取表名
func tableName(name string) string {
	return tablePrefix + name
}

// debug
func debug(v ...interface{}) {
	beego.Debug(v...)
}

// DBVersion  数据库版本
func DBVersion() string {
	var lists []orm.ParamsList
	o.Raw("select version()").ValuesList(&lists)
	return lists[0][0].(string)
}

// ConcatenateError
func concatenateError(err error, stderr string) error {
	if len(stderr) == 0 {
		return err
	}
	return fmt.Errorf("%v: %s", err, stderr)
}

// CheckErr 检查错误类型
func CheckErr(err error) bool {
	if nil == err {
		return true
	}
	return false
}

// NewError 生成异常
func NewError(errMsg string) error {
	return fmt.Errorf("Error Msg: %s", errMsg)
}
