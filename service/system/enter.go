package system

//RouterGroup里注册我自己的Router
type ServiceGroup struct {
	UserService
	ManageService
	EffectService
	AuthorityService
}
