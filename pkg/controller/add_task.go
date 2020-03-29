package controller

import (
	"github.com/umutkocasarac/task-operator/pkg/controller/task"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, task.Add)
}
