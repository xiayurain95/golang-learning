package LB

import (
	context "context"
	"flag"
	fmt "fmt"
	"log"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"
)

func HelloClient() {
	var (
		serv = flag.String("service", "hello_service", "service name")
		reg  = flag.String("etcd", "http://laocalhoust:2379", "address of etcd service")
	)
	flag.Parse()

	r := NewResolver(*serv, "etcd_naming")
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		log.Printf("dial grpc fails")
		return
	}

	ticker := time.NewTicker(time.Second * 1)
	for t := range ticker.C {
		client := NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &String{Value: "world" + strconv.Itoa(t.Second())})
		if err != nil {
			log.Printf("client Failed")
			fmt.Printf("%v", client)
			return
		}
		fmt.Println(resp.GetValue())
	}
}
