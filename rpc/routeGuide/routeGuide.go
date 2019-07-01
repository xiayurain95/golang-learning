package routeGuide

import (
	"net"

	grpc "google.golang.org/grpc"
)

func RouteGuideRpc() {
	rpcServer := grpc.NewServer()
	RegisterRouteGuideServer(rpcServer, new(RouteGuideImp))
	lis, err := net.Listen("tcp", ":9996")
	if err != nil {
		return
	}
	rpcServer.Serve(lis)
}
