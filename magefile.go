// Copyright KhulnaSoft, Ltd.
// Please see the included license file for details.

//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"

	// mage:import jenkins
	"github.com/khulnasoft/casskube/mage/jenkins"
	// mage:import operator
	"github.com/khulnasoft/casskube/mage/operator"
	// mage:import integ
	_ "github.com/khulnasoft/casskube/mage/integ-tests"
	// mage:import lint
	_ "github.com/khulnasoft/casskube/mage/linting"
	// mage:import k8s
	_ "github.com/khulnasoft/casskube/mage/k8s"
)

// Clean all build artifacts, does not clean up old docker images.
func Clean() {
	mg.Deps(operator.Clean)
	mg.Deps(jenkins.Clean)
}
