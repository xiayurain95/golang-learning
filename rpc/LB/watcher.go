package LB

import (
	"context"
	"fmt"

	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/naming"
)

func (w *watcher) Close() {
}
func (w *watcher) Next() ([]*naming.Update, error) {
	prefix := fmt.Sprintf("/%s/%s", w.Prefix, w.re.serviceName)
	if !w.isInitialized {
		resp, err := w.client.Get(context.Background(), prefix, etcd3.WithPrefix())
		if err != nil {
		} else {
			w.isInitialized = true
			addr := extarctAddr(resp)
			update := make([]*naming.Update, len(addr))
			for k, v := range addr {
				update[k] = &naming.Update{Op: naming.Add, Addr: v}
			}
			return update, nil
		}
	}

	rch := w.client.Watch(context.Background(), prefix, etcd3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case mvccpb.DELETE:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil

			}
		}
	}

	return nil, nil
}

func extarctAddr(resp *etcd3.GetResponse) []string {
	addr := []string{}

	if resp == nil || resp.Kvs == nil {
		return addr
	}

	for _, v := range resp.Kvs {
		if v.Value != nil {
			addr = append(addr, string(v.Value))
		}
	}

	return addr
}
