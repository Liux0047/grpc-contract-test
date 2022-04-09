module github.com/Liux0047/grpc-contract-test

go 1.17

require (
	github.com/google/go-cmp v0.5.5
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

require golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/liux0047/grpc-contract-test v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20220325170049-de3da57026de // indirect
	golang.org/x/sys v0.0.0-20220325203850-36772127a21f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220324131243-acbaeb5b85eb // indirect
)

replace github.com/liux0047/grpc-contract-test => ./
