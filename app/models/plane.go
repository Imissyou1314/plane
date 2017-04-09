package models

import (
	"time"
)

/** Plane飞机实体*/
type Plane struct {
	Id            int       																		// PlaneID 飞机ID
	PlaneName     string    `orm:"size(20)"` 										// PlaneName 飞机名
	ImageUrl      string    `orm:"size(20)"` 										// 飞机图标
	UserId        int       																		// userid   用户ID
	PlayTime      time.Time `orm:"type(datatime)"`              //  飞行时间
	BeginTime     time.Time `orm:"auto_now_add;type(datetime)"` //  开始时间
	BeginAddress  string    `orm:"size(20)"`                    //  开始地点
	ResultAddress string    `orm:"size(20)"`                    //  结束地点
	Email         string    `orm:"size(30)"`                    //  用户定义消息结束接收邮箱
	WatchNum      int       																		// 被拦截的数次
	IsBack        bool      																		// 是否已经被回收
}

type PlaneLog struct {
	Id         int       																			// 飞机记录
	LogTime    time.Time `orm:"auto_now_add;type(datetime)"` 	// 记录时间
	LogAddress string    `orm:"size(20)"`                    	// 地区名
	Log        float32   																			// 经度
	dog        float32   																			// 维度
}
