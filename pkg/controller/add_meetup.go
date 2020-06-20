package controller

import (
	"github.com/cloudnative-id/community-operator/pkg/controller/meetup"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, meetup.AddManager)

	// AddToDispatcherFuncs is a list of functions to create controllers and add them to a manager.
	AddToDispatcherFuncs = append(AddToDispatcherFuncs, meetup.AddDispatcher)
}
