package router

import "dogrod-web-service/router/blog"

type RouterGroup struct {
	Blog blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
