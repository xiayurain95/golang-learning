package LB

import (
	context "context"
	"flag"
	fmt "fmt"
	"time"

	grpc "google.golang.org/grpc"
)

func helloClient() {
	var (
		serv = flag.String("service", "hello_service", "service name")
		reg  = flag.String("etcd address", "localhost:2379", "address of etcd service")
	)
	flag.Parse()

	r := NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		fmt.Errorf("dial grpc fails")
		return
	}

	ticker := time.NewTicker(time.Second * 1)
	for t := range ticker.C {
		client := NewGreeterClient(conn)
		resp,err:=
	}
}
