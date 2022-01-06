package service

import "dogrod-web-service/service/blog"

type ServiceGroup struct {
	BlogServiceGroup blog.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
