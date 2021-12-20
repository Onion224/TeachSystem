package system

import (
	"server/global"
	"server/model"
	"server/model/common/request"
)

type AuthorityService struct{}

func (s *AuthorityService) CreateAuthority(auth model.Authority) (err error, authority model.Authority) {
	err = global.GVA_DB.Create(&auth).Error
	return err, auth
}

func (s *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Authority{})
	err = db.Where("parent_id=?", "0").Count(&total).Error
	var authority []model.Authority
	err = db.Limit(limit).Offset(offset).Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = s.findChildrenAuthority(&authority[k])
		}

	}
	return err, authority, total
}

func (authorityService *AuthorityService) findChildrenAuthority(authority *model.Authority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id=?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err

}
