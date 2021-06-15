# waynga-grpc

###Installation instructions:
1. Install protoc compiler using instructions here : https://grpc.io/docs/protoc-installation/
- Specifically follow the "Install pre-compiled binaries" section
2. Follow the pre-req section here https://grpc.io/docs/languages/go/quickstart/ to install go plugins


###Updating proto
If you are using "Any" type in proto , then make sure you use the following option for compiling your proto file
`protoc --proto_path=/usr/local/include --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative entityservice/entity.proto`

The first proto_path points to where you have saved the "include" folder contents from #1 in the installation
The second proto_path points to the actual proto file location that you are trying to compile

