package request

// User register structure
type Register struct {
	Name          string `form:"name" json:"name" uri:"name"`
	Account       string `form:"account" json:"account" uri:"account" binding:"required"`
	Password      string `form:"password" json:"password" uri:"password" binding:"required"`
	AuthorityId   string `form:"type" json:"type" uri:"type" gorm:"default:3"`
	ParentID      string `form:"parent_id" json:"parentID" uri:"parent_id"`
	Email         string `form:"email" json:"email" uri:"email"`
	Phone         string `form:"phone" json:"phone" uri:"phone"`
	AuthorityName string `form:"authority_name" uri:"authority_name" json:"authority_name"`
}

type Login struct {
	Account  string `json:"account"`  //账号
	Password string `json:"password"` //密码
}
