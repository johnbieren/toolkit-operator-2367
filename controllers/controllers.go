package controllers

import (
	"github.com/Troy876/toolkit-operator-2367/controllers/bar"
	"github.com/Troy876/toolkit-operator-2367/controllers/foo"
	"github.com/redhat-appstudio/operator-toolkit/controller"
)

var EnabledControllers = []controller.Controller{
	&bar.Controller{},
	&foo.Controller{},
}
