package models

/** User model */
type User struct {
	Id       int    // UserId 用户ID
	UserName string `orm:"size(20)"` // 用户名
	Password string `orm:"size(20)"` // 密码
	Accout   string `orm:"size(20)"` // 账号
	ImageUrl string `orm:"size(20)"` // 头像图片地址
	Email    string `orm:"size(20)"` // 邮箱
	Phone    string `orm:"size(20)"` // 用户电话
	Address  string `orm:"size(20)"` // 用户地址
	Status   int    // 用户状态
}
