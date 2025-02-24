// Copyright KhulnaSoft, Ltd.
// Please see the included license file for details.

package apis

import (
	"github.com/khulnasoft/casskube/operator/pkg/apis/cassandra/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1beta1.SchemeBuilder.AddToScheme)
}
