# grpc-contract-test
 A Contract Testing framework for gRPC

# regen the proto
C:\Users\Allen\go\src\GitHub\grpc-contract-test> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\contract\contract.proto