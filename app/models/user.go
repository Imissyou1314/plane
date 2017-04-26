package models

type UserStatus int

const (
	_      UserStatus = iota
	Login             //登录状态
	Logout            // 登出状态
	Icon              // 冻结状态
	Remove            // 删除状态
)

/**User model */
type User struct {
	Id       int64  // UserId 用户ID
	UserName string `orm:"size(20)"` // 用户名
	Password string `orm:"size(20)"` // 密码
	Sex      int    //性别
	Salt     string `orm:"size(20)"`           // 加密研制
	Accout   string `orm:"size(20)"`           // 账号
	ImageUrl string `orm:"default();size(20)"` // 头像图片地址
	Email    string `orm:"size(20)"`           // 邮箱
	Phone    string `orm:"size(20)"`           // 用户电话
	Address  string `orm:"size(20)"`           // 用户地址
	Status   int    // 用户状态
}
