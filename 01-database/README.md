# protocol buffers database
database storing tasks using protocol buffers and go


## Installation

pre-requisite: protocol compiler v3.x is installed

Run in workspace:
```shell
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto

cd $GOPATH/src/todo
protoc --go_out=. todo.proto 

cd $GOPATH
go install .src/cmd/todo
go install proto_buf_db/cmd/todo
go install proto_buf_db/...
```

## Running
todo

  

