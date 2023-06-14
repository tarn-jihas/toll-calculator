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



.PHONY: obu invoicer