package LB

import (
	context "context"
	"flag"
	fmt "fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc "google.golang.org/grpc"
)

func (p *helloImp) SayHello(ctx context.Context, str *String) (*String, error) {
	fmt.Printf("%v:Receive is %s\n", time.Now(), str.GetValue())
	return &String{Value: "hello:" + str.GetValue()}, nil
}

func HelloServer() {

	var (
		serv = flag.String("service", "hello_service", "service name")
		port = flag.Int("port", 9996, "port of the service")
		reg  = flag.String("etcd address", "localhost:2379", "address of etcd service")
	)
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Errorf("Listen to socket 9996 failed")
		return
	}

	registerIns := NewregisterBlock("etcd_naming", *serv, make(chan bool))
	if err := registerIns.Register(*serv, "127.0.0.1", *port, *reg, time.Second*10, 15); err != nil {
		fmt.Errorf("Rigster service fails")
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		registerIns.Unregister()
		os.Exit(1)
	}()

	fmt.Println("starting service hello")
	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, new(helloImp))
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Errorf("grpcServer.Server failed")
		return
	}

}
