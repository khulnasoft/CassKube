// Copyright KhulnaSoft, Ltd.
// Please see the included license file for details.

package controller

import (
	"github.com/khulnasoft/casskube/operator/pkg/controller/cassandradatacenter"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cassandradatacenter.Add)
}
