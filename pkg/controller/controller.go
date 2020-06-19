package controller

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"github.com/cloudnative-id/community-operator/dispatcher"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error
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

func AddToDispatcher(disp dispatcher.Dispatcher) {
	for _, f := range AddToDispatcherFuncs {
		f(disp)
	}
}
