package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model"
)

type UserService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func (userService *UserService) Register(u model.User) (err error, userInter model.User) {
	//定义一个User对象
	var user model.User
	// 判断用户名是否注册
	if !errors.Is(global.GVA_DB.Where("account = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("账号已经注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.ID = uuid.NewV4()
	//u.Account = u.Account
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) Login(u *model.User) (err error, userInter *model.User) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("account = ? AND password = ?", u.Account, u.Password).First(&user).Error
	return err, &user
}
