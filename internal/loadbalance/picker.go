package loadbalance

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

var _ base.PickerBuilder = (*Picker)(nil)

type Picker struct {
	mu        sync.Mutex
	leader    balancer.SubConn
	followers []balancer.SubConn
	current   uint64
}

func (p *Picker) Build(buildInfo base.PickerBuildInfo) balancer.Picker {
	p.mu.Lock()
	defer p.mu.Unlock()
	var followers []balancer.SubConn
	for sc, scInfo := range buildInfo.ReadySCs {
		isLeader := scInfo.Address.Attributes.Value("is_leader").(bool)
		if isLeader {
			p.leader = sc
			continue
		}
		followers = append(followers, sc)
	}
	p.followers = followers
	return p
}

func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	//TODO implement me
	panic("implement me")
}
