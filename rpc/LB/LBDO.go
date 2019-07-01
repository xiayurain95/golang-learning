package LB

type resolver struct {
	serviceName string
}

type watcher struct {
	re            *resolver
	client        *etcd3.Client
	isInitialized bool
}

func NewResolver(s string) *resolver {
	return &resolver{serviceName: s}
}
