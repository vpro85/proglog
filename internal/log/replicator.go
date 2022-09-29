package log

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	api "proglog/api/v1"
	"sync"
)

type Replicator struct {
	DialOptions []grpc.DialOption
	LocalServer api.LogClient

	logger  *zap.Logger
	mu      sync.Mutex
	servers map[string]chan struct{}
	closed  bool
	close   chan struct{}
}

func (r *Replicator) Join(name, addr string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.init()
	if r.closed {
		return nil
	}
	if _, ok := r.servers[name]; ok {
		// already replicated so skip
		return nil
	}
	r.servers[name] = make(chan struct{})

	go r.replicate(addr, r.servers[name])
	return nil
}

func (r *Replicator) replicate(addr string, leave chan struct{}) {
	cc, err := grpc.Dial(addr, r.DialOptions...)
	if err != nil {
		r.logError(err, "failed to dial", addr)
		return
	}
	defer cc.Close()

	client := api.NewLogClient(cc)

	ctx := context.Background()
	stream, err := client.ConsumeStream(ctx, &api.ConsumeRequest{
		Offset: 0,
	})
	if err != nil {
		r.logError(err, "failed to consume", addr)
		return
	}
	records := make(chan *api.Record)
	go func() {
		for {
			recv, err := stream.Recv()
			if err != nil {
				r.logError(err, "failed to receive", addr)
				return
			}
			records <- recv.Record
		}
	}()
	for {
		select {
		case <-r.close:
			return
		case <-leave:
			return
		case record := <-records:
			_, err = r.LocalServer.Produce(ctx, &api.ProduceRequest{
				Record: record,
			})
			if err != nil {
				r.logError(err, "failed to produce", addr)
				return
			}
		}
	}
}
