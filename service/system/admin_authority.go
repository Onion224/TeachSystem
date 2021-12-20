package system

import (
	"server/global"
	"server/model"
)

type AuthorityService struct{}

func (s *AuthorityService) CreateAuthority(auth model.User) (err error, authority model.User) {
	err = global.GVA_DB.Create(&auth).Error
	return err, auth
}
