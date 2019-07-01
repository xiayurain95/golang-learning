package main

import rpc "learningGo/rpc/hello"

func main() {
	go rpc.HelloRpc()
	rpc.ChannelClient()
}
