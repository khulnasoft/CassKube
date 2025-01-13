// Copyright KhulnaSoft, Ltd.
// Please see the included license file for details.

package oplabels

const (
	ManagedByLabel             = "app.kubernetes.io/managed-by"
	ManagedByLabelValue        = "casskube"
	ManagedByLabelDefunctValue = "cassandra-operator"
)

func AddManagedByLabel(m map[string]string) {
	m[ManagedByLabel] = ManagedByLabelValue
}

func AddDefunctManagedByLabel(m map[string]string) {
	m[ManagedByLabel] = ManagedByLabelDefunctValue
}

func HasManagedByCassandraOperatorLabel(m map[string]string) bool {
	v, ok := m[ManagedByLabel]
	return ok && v == ManagedByLabelValue
}
