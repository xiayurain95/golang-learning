package LB

import (
	"errors"
	"fmt"
	"strings"

	etcd3 "github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/naming"
)

func (re *resolver) Reslove(target string) (naming.Watcher, error) {
	if re.serviceName == "" {
		return nil, errors.New("no service name provided")
	}

	client, err := etcd3.New(etcd3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		return nil, fmt.Errorf("creat etcd3 client failed, %s", err.Error())
	}
	return watcher{re: re, client: client}, nil
}
