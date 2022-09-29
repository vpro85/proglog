package log

import (
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
