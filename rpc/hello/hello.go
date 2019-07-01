package hello

import (
	context "context"
	"io"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func (p *helloServerImp) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello" + args.GetValue()}
	return reply, nil
}
func (p *helloServerImp) Channel(stream HelloMessage_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return nil
		}
	}
}
func HelloRpc() {
	grpcServer := grpc.NewServer()
	RegisterHelloMessageServer(grpcServer, new(helloServerImp))
	if lis, err := net.Listen("tcp", ":9996"); err != nil {
		return
	} else {
		grpcServer.Serve(lis)
	}
}

func ChannelClient() {
	if conn, err := grpc.Dial("localhost:9996", grpc.WithInsecure()); err != nil {
		return
	} else {
		defer conn.Close()

		client := NewHelloMessageClient(conn)
		stream, err := client.Channel(context.Background())
		if err != nil {
			return
		}

		ch := make(chan struct{}, 2)
		go func() {
			for {
				reply, err := stream.Recv()
				if err != nil {
					log.Println("%v", err)
					return
				}
				log.Println(reply.GetValue())
				ch <- struct{}{}
			}
		}()
		stream.Send(&String{Value: "xiayu"})
		stream.Send(&String{Value: "abc"})
		<-ch
		<-ch

	}
}

func HelloClient() {
	if conn, err := grpc.Dial("localhost:9996", grpc.WithInsecure()); err != nil {
		return
	} else {
		defer conn.Close()

		client := NewHelloMessageClient(conn)
		reply, err := client.Hello(context.Background(), &String{Value: "xiayu"})
		if err != nil {
			return
		}
		log.Println(reply.GetValue())
	}
}
