module github.com/tung238/shippy/shippy-cli-consignment

go 1.16

replace github.com/tung238/shippy/shippy-cli-consignment => ../shippy-cli-consignment
replace github.com/tung238/shippy/shippy-service-consignment => ../shippy-service-consignment

require github.com/asim/go-micro/v3 v3.5.1 // indirect
require github.com/tung238/shippy/shippy-service-consignment v0.0.0