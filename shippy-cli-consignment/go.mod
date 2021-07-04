module github.com/tung238/shippy/shippy-cli-consignment

go 1.16

replace github.com/tung238/shippy/shippy-cli-consignment => ../shippy-cli-consignment

// replace github.com/tung238/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/asim/go-micro/v3 v3.5.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/tung238/shippy/shippy-service-consignment v0.0.0-20210703093814-181004926850 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.25.0 // indirect
)
