package LB

// code of https://colobu.com/2017/03/25/grpc-naming-and-load-balance/
import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	etcd3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

var Prefix = "etcd_naming"
var client etcd3.Client
var serverKey string
var stopSignal = make(chan bool, 1)

func Register(name string, host string, port int, target string, interval time.Duration, ttl int) error {
	serviceValue := fmt.Sprintf("%s:%d", host, port)
	serviceKey := fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceValue)
	var err error
	client, err := etcd3.New(etcd3.Config{Endpoints: strings.Split(target, ",")})
	if err != nil {
		fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
	}
	go func() {
		ticker := time.NewTicker(interval)
		for {
			resp, _ := client.Grant(context.TODO(), int64(ttl))
			_, err := client.Get(context.Background(), serviceKey)
			if err != nil {
				if err == rpctypes.ErrKeyNotFound {
					if _, err := client.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
						log.Printf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
					}
				} else {
					log.Printf("grpclb: service '%s' connect to etcd3 failed: %s", name, err.Error())
				}
			} else {
				if _, err := client.Put(context.Background(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
					log.Printf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
				}
			}

			select {
			case <-stopSignal:
				return
			case <-ticker.C:
			}
		}
	}()
	return nil
}
func Unregister() error {
	stopSignal <- true
	stopSignal = make(chan bool, 1)
	var err error
	if _, err = client.Delete(context.Background(), serverKey); err != nil {

		log.Printf("grpclb: deregister '%s' failed: %s", serverKey, err.Error())
	} else {
		log.Printf("grpclb: deregister '%s' ok.", serverKey)
	}
	return err
}
