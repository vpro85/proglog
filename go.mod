module proglog

go 1.13

require (
	github.com/casbin/casbin v1.9.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/memberlist v0.4.0 // indirect
	github.com/hashicorp/raft v1.1.1 // indirect
	github.com/hashicorp/serf v0.8.5
	github.com/stretchr/testify v1.7.0
	github.com/travisjeffery/go-dynaport v1.0.0
	github.com/tysonmote/gommap v0.0.2
	go.opencensus.io v0.23.0
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20220927171203-f486391704dc // indirect
	golang.org/x/sys v0.0.0-20220927170352-d9d178bc13c6 // indirect
	google.golang.org/genproto v0.0.0-20220927151529-dcaddaf36704
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/hashicorp/raft-boltdb => github.com/travisjeffery/raft-boltdb v1.0.0
