all: receiver obu agg gate

gate:
	@go build -o bin/gateway ./gateway ;
	@./bin/gateway ;

obu:
	@go build -o bin/obu ./obu ;
	@./bin/obu ;

receiver:
	@go build -o bin/receiver ./data_receiver ;
	@./bin/receiver ;

calc:
	@go build -o bin/calculator ./distance_calculator ;
	@./bin/calculator ;
	
agg:
	@go build -o bin/aggregator ./aggregator ;
	@./bin/aggregator ;

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

.PHONY: obu invoicer