protoc greet/greetpb/greet.proto --go-grpc_out=. greet/greetpb/greet.proto

protoc -I src/ --go-grpc_out=src/ src/simple/simple.proto