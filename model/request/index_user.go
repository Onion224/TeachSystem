package request

// User register structure
type Register struct {
	Name     string `form:"name" json:"name" uri:"name"`
	Account  string `form:"account" json:"account" uri:"account" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	Type     int    `form:"type" json:"type" uri:"type" gorm:"default:3"`
	Email    string `form:"email" json:"email" uri:"email"`
	Phone    string `form:"phone" json:"phone" uri:"phone"`
}

type Login struct {
	Account  string `json:"account"`  //账号
	Password string `json:"password"` //密码
}
