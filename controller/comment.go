package controller

import "github.com/tomocy/mvc/controller"

type Comment struct {
	*controller.Base
}

func NewComment() controller.Controller {
	return &Comment{
		Base: controller.NewBase(),
	}
}
