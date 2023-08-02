package controllers

import (
	"github.com/redhat-appstudio/operator-toolkit/controller"
)

var EnabledControllers = []controller.Controller{
	&FooReconciler{},
	&BarReconciler{},
}
