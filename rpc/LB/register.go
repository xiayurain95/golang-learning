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

func (p *registerBlock) Register(name string, host string, port int, target string, interval time.Duration, ttl int) error {
	serviceValue := fmt.Sprintf("%s:%d", host, port)
	serviceKey := fmt.Sprintf("/%s/%s/%s", p.Prefix, name, serviceValue)
	var err error
	p.client, err = etcd3.New(etcd3.Config{Endpoints: strings.Split(target, ",")})
	if err != nil {
		log.Printf("grpclb: create etcd3 client failed: %v", err)
		return fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
	}
	go func() {
		ticker := time.NewTicker(interval)
		for {
			resp, _ := p.client.Grant(context.TODO(), int64(ttl))
			_, err := p.client.Get(context.Background(), serviceKey)
			if err != nil {
				if err == rpctypes.ErrKeyNotFound {
					if _, err := p.client.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
						log.Printf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
					}
				} else {
					log.Printf("grpclb: service '%s' connect to etcd3 failed: %s", name, err.Error())
				}
			} else {
				if _, err := p.client.Put(context.Background(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
					log.Printf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
				}
			}

			select {
			case <-p.stopSignal:
				return
			case <-ticker.C:
			}
		}
	}()
	return nil
}
func (p *registerBlock) Unregister() error {
	log.Printf("deregister pre-start")
	p.stopSignal <- true
	p.stopSignal = make(chan bool, 1)
	log.Printf("deregister start")
	var err error
	if _, err = p.client.Delete(context.Background(), p.serverKey); err != nil {
		log.Printf("grpclb: deregister '%s' failed: %s", p.serverKey, err.Error())
	} else {
		log.Printf("grpclb: deregister '%s' ok.", p.serverKey)
	}
	return err
}
