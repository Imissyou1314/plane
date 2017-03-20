package models

import (
	"time"
)

/** 日志  */
type Log struct {
	Id       int       																			// 日志ID
	Time     time.Time `orm:"auto_now_add;type(datetime)"`  // 记录时间
	UserId   int       																			// 用户ID
	IP       string    `orm:"size(20)"`            					// ip
	Address  string    `orm:"size(20)"`            					// 地址
	Drive    string    `orm:"size(20)"`            					// 设备
	UserName string    `orm:"size(20)"`            					// 用户名
	TimeOut  time.Time `orm:"null;type(datetime)"` 					// 离开时间
}
