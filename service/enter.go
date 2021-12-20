package service

import "server/service/system"

type ServiceGroup struct {
	SystemService system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
