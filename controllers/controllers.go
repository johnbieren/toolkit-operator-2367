package controllers

import (
	"github.com/redhat-appstudio/operator-toolkit-example/controllers/bar"
	"github.com/redhat-appstudio/operator-toolkit-example/controllers/foo"
	"github.com/redhat-appstudio/operator-toolkit/controller"
)

var EnabledControllers = []controller.Controller{
	&bar.Controller{},
	&foo.Controller{},
}
