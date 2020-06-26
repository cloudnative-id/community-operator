package controller

import (
	"github.com/cloudnative-id/community-operator/pkg/controller/announcement"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, announcement.AddManager)

	// AddToDispatcherFuncs is a list of functions to create controllers and add them to a manager.
	AddToDispatcherFuncs = append(AddToDispatcherFuncs, announcement.AddDispatcher)
}
