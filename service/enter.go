package service

import (
	"dogrod-web-service/service/blog"
	"dogrod-web-service/service/system"
)

type ServiceGroup struct {
	BlogServiceGroup blog.ServiceGroup
	UserServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
