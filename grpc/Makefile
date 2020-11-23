
all:./service.pb.go

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.
