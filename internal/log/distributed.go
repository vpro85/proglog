package log

import (
	"github.com/hashicorp/raft"
	"os"
	"path/filepath"
)

type DistributedLog struct {
	config Config
	log    *Log
	raft   *raft.Raft
}

func NewDistributedLog(dataDir string, config Config) (*DistributedLog, error) {
	l := &DistributedLog{
		config: config,
	}
	if err := l.setupLog(dataDir); err != nil {
		return nil, err
	}
	if err := l.setupRaft(dataDir); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *DistributedLog) setupLog(dataDir string) error {
	logDir := filepath.Join(dataDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	var err error
	l.log, err = NewLog(logDir, l.config)
	return err
}
