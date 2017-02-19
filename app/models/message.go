package models

import (
	"time"
)

/** 消息实体 */
type Message struct {
	Id         int       // MessageId 消息ID
	PlanId     int       // userid接收对象ID
	toUserId   int       // userid发送对象ID
	fromUserId int       // planeID飞机对象ID
	IsRead     bool      // boole 是否已经被读取
	Context    string    `orm:"type(text)"`                  // 消息内容
	ImageUrl   string    `orm:"size(20)"`                    // 消息图片链接
	SendTime   time.Time `orm:"auto_now_add;type(datetime)"` // 发送时间
	ReadTime   time.Time `orm:"type(datetime)"`              // 接收时间
	Address    string    `orm:"size(20)"`                    // 该飞机所在的地方
	log        float64   // 经度
	don        float64   // 维度
}
