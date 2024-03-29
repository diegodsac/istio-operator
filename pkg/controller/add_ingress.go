package controller

import (
	"istio-operator/pkg/controller/ingress"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ingress.Add)
}
