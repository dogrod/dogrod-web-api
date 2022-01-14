package service

import (
	"dogrod-web-service/service/blog"
	"dogrod-web-service/service/user"
)

type ServiceGroup struct {
	BlogServiceGroup blog.ServiceGroup
	UserServiceGroup user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
