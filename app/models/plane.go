package models

import (
	"time"
)

/** Plane飞机实体*/
type Plane struct {
	Id            int64     // PlaneID 飞机ID
	PlaneName     string    `orm:"size(200)"`                   // PlaneName 飞机名
	ImageUrl      string    `orm:"size(200)"`                   // 飞机图标
	UserId        int64     `orm:"size(100)"`                   // userid   用户ID
	PlayTime      time.Time `orm:"type(datatime)"`              //  飞行时间
	BeginTime     time.Time `orm:"auto_now_add;type(datetime)"` //  开始时间
	BeginAddress  string    `orm:"size(200)"`                   //  开始地点
	ResultAddress string    `orm:"size(200)"`                   //  结束地点
	Distance      float64   `orm:"size(500)"`                   //  两点间的距离
	PlaneSpeed    float64   `orm:"size(500)"`                   //  PlaneSpeed 飞机的速度
	Email         string    `orm:"size(130)"`                   //  用户定义消息结束接收邮箱
	WatchNum      int       `orm:"size(20)"`                    // 被拦截的数次
	IsBack        bool      `orm:"size(2)"`                     // 是否已经被回收
}

type PlaneLog struct {
	Id         int64     // 飞机记录
	LogTime    time.Time `orm:"auto_now_add;type(datetime)"` // 记录时间
	LogAddress string    `orm:"size(20)"`                    // 地区名
	Log        float32   // 经度
	dog        float32   // 维度
}
