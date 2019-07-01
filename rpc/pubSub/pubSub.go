package pubSub

import (
	context "context"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
)

func (p *pubSubImp) GetStream(args *Token, stream PubSubService_GetStreamServer) error {
	if ch, ok := p.m[args.GetValue()]; !ok {
		return nil
	} else {
		for v := range ch {
			if err := stream.Send(&String{Value: v.(string)}); err != nil {
				log.Println(err)
			} else {
				log.Println("[send data] subscibe :", v.(string))
			}
		}
	}
	return nil
}

func (p *pubSubImp) Publish(ctx context.Context, args *String) (*String, error) {
	log.Println("[prePublish]:", args.GetValue())
	p.Publisher.Publish(args.GetValue())
	log.Println("[published]:", args.GetValue())
	return &String{}, nil
}
func (p *pubSubImp) Subscrib(ctx context.Context, args *String) (*Token, error) {

	ch := p.Publisher.SubscribeTopic(func(v interface{}) bool {
		log.Println("[check value] subscribe  :", v)
		if value, ok := v.(string); ok {
			if strings.HasPrefix(value, args.GetValue()) {
				return true
			}
		}
		return false
	})

	log.Println("[suceessed] subscribe :", args.GetValue())

	rand.Seed(time.Now().Unix())
	for {
		num := rand.Uint64()
		if _, ok := p.m[num]; !ok {
			p.m[num] = ch
			return &Token{Value: num}, nil
		}
	}
}

func PubSubRpc() {
	grpcServer := grpc.NewServer()
	RegisterPubSubServiceServer(grpcServer, NewPublisher())
	if lis, err := net.Listen("tcp", ":9996"); err != nil {
		return
	} else {
		grpcServer.Serve(lis)
	}
}

func PubClient() {
	if conn, err := grpc.Dial("localhost:9996", grpc.WithInsecure()); err != nil {
		return
	} else {
		defer conn.Close()

		ch := make(chan struct{}, 2)

		client := NewPubSubServiceClient(conn)
		golangToken, err := client.Subscrib(context.Background(), &String{Value: "golang"})
		if err != nil {
			return
		}
		golang, err := client.GetStream(context.Background(), golangToken)
		if err != nil {
			return
		}

		rustToken, err := client.Subscrib(context.Background(), &String{Value: "rust"})
		if err != nil {
			return
		}
		rust, err := client.GetStream(context.Background(), rustToken)
		if err != nil {
			return
		}
		
		go func() {
			for {
				if rev, err := golang.Recv(); err != nil {
					return
				} else {
					log.Println(rev.GetValue())
					ch <- struct{}{}
				}
			}
		}()
		go func() {
			for {
				if rev, err := rust.Recv(); err != nil {
					return
				} else {
					log.Println(rev.GetValue())
					ch <- struct{}{}
				}
			}
		}()

		client.Publish(context.Background(), &String{Value: "rust:hello,rust"})
		client.Publish(context.Background(), &String{Value: "golang:hello,golang"})
		<-ch
		<-ch
	}
}
