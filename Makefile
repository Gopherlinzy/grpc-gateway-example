protoc:
	protoc --grpc-gateway_out=.  --go_out=.  --openapiv2_out . --openapiv2_opt use_go_templates=true --go-grpc_out=.  ./proto/*.proto