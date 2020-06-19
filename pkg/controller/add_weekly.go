package controller

import (
	"github.com/cloudnative-id/community-operator/pkg/controller/weekly"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, weekly.AddManager)
	AddToDispatcherFuncs = append(AddToDispatcherFuncs, weekly.AddDispatcher)
}
