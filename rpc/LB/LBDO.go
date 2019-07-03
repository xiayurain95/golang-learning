package LB

import etcd3 "github.com/coreos/etcd/clientv3"

type resolver struct {
	serviceName string
}

type registerBlock struct {
	Prefix     string
	client     *etcd3.Client
	serverKey  string
	stopSignal chan bool
}

type helloImp struct {
}

type watcher struct {
	re            *resolver
	client        *etcd3.Client //maybe need a connecting pool for this,short time ~~ never mind for now
	isInitialized bool
	Prefix        string "PreFix is the etcd prefix(combined with the registerBlock Prefix Feild)"
}

func NewResolver(s string) *resolver {
	return &resolver{serviceName: s}
}

func NewregisterBlock(prefix string, serverKey string, stopSignal chan bool) *registerBlock {
	return &registerBlock{
		Prefix:     prefix,
		serverKey:  serverKey,
		stopSignal: stopSignal,
	}
}
