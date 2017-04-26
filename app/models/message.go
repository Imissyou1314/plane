package models

import (
	"time"
)

/** 消息实体 */
type Message struct {
	Id         int64     // MessageId 消息ID
	PlaneId    int64     // userid接收对象ID
	ToUserId   int64     // userid发送对象ID
	FromUserId int64     // planeID飞机对象ID
	IsRead     bool      `orm:"default(false)"`              // boole 是否已经被读取
	Context    string    `orm:"type(text)"`                  // 消息内容
	ImageUrl   string    `orm:"size(20)"`                    // 消息图片链接
	SendTime   time.Time `orm:"auto_now_add;type(datetime)"` // 发送时间
	ReadTime   time.Time `orm:"type(datetime)"`              // 接收时间
	Address    string    `orm:"size(20)"`                    // 该飞机所在的地方
	Log        float64   // 经度
	Don        float64   // 维度
}
