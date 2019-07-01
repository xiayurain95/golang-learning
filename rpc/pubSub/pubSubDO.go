package pubSub

import (
	"time"

	"github.com/docker/docker/pkg/pubsub"
)

type pubSubImp struct {
	Publisher *pubsub.Publisher
	m         map[uint64]chan interface{}
}

func NewPublisher() *pubSubImp {
	return &pubSubImp{Publisher: pubsub.NewPublisher(1000*time.Millisecond, 10),
		m: make(map[uint64]chan interface{})}
}
