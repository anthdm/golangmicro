# A microservice example written in Golang

# Installing protobuffer

### Linux

```
sudo apt install -y protobuf-compiler
```

### MacOS

```
brew install protobuff
```

### GRPC and Protobuffer package dependencies

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

NOTE: You should add the `protoc-gen-go-grpc` to your PATH

```
PATH="${PATH}:${HOME}/go/bin"

```

### Running the service

```
make run
```
