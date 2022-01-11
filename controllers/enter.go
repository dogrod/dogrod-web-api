package controllers

import "dogrod-web-service/controllers/blog"

type ControllerGroup struct {
	BlogControllerGroup blog.ControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)
