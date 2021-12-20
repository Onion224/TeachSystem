package response

import (
	"server/model"
)

type SysUserResponse struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	User model.User `json:"user"`
}
