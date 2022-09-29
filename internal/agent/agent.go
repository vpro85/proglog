package agent

import (
	"crypto/tls"
	"google.golang.org/grpc"
	"proglog/internal/discovery"
	"proglog/internal/log"
	"sync"
)

type Agent struct {
	Config

	log        *log.Log
	server     *grpc.Server
	membership *discovery.Membership
	replicator *log.Replicator

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
}

type Config struct {
	ServerTLSConfig *tls.Config
	PeerTLSConfig   *tls.Config
}
