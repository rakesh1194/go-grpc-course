# For Greeting Example
protoc greet/greetpb/greet.proto --go-grpc_out=greet/ greet/greetpb/greet.proto

protoc -I greet/ --go_out=greet/ greet/greetpb/greet.proto


# For Calculator Example
protoc calculator/sumpb/sum.proto --go-grpc_out=calculator/ calculator/sumpb/sum.proto

protoc -I calculator/ --go_out=calculator/ calculator/sumpb/sum.proto