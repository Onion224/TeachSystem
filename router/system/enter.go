package system

import (
	"server/service"
)

//RouterGroup里注册我自己的Router
type RouterGroup struct {
	ManageRouter ManageRouter
	EffectRouter EffectRouter
	BaseRouter   BaseRouter
	UserRouter   UserRouter
}

var RouterGroupApp = new(RouterGroup)

var (
	userService   = service.ServiceGroupApp.SystemService.UserService
	manageService = service.ServiceGroupApp.SystemService.ManageService
	effectService = service.ServiceGroupApp.SystemService.EffectService
)
