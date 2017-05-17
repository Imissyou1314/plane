package models

import "time"

// Video 日志模块
type Video struct {
	Id         int64     // 视屏ID
	UploadTime time.Time `orm:"auto_now_add;type(datetime)"` // 记录时间
	UserId     int64     `orm:"size(100)"`                   // 用户ID
	Ip         string    `orm:"size(20)"`                    // 上传用户IDip
	Drive      string    `orm:"size(20)"`                    // 上传设备
	UserName   string    `orm:"size(20)"`                    // 用户名
	VideoName  string    `orm:"size(140)"`                   // VideoName 视屏名字
	VideoUrl   string    `orm:"size(150)"`                   // VideoUrl 视屏地址
	VideoTag   string    `orm:"size(100)"`                   // VideoTag 视屏标签
	VideoSize  int64     `orm:"size(100)"`                   // VideoSize视屏大小
	VideoType  string    `orm:"size(100)"`                   // VideoType 视屏类型
}
