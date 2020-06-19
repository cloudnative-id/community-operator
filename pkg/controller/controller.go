package controller

import (
	"github.com/cloudnative-id/community-operator/dispatcher"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToDispatcherFuncs is a list of functions to create controllers and add them to a manager.
var AddToDispatcherFuncs []func(dispatcher.Dispatcher)

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil
}

// AddToDispatcher adds all Controllers to the Manager
func AddToDispatcher(disp dispatcher.Dispatcher) {
	for _, f := range AddToDispatcherFuncs {
		f(disp)
	}
}
