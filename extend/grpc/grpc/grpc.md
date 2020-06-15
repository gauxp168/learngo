## grpc

#### 安装gRPC和Protobuf
     go get github.com/golang/protobuf/proto
     go get google.golang.org/grpc（无法使用，用如下命令代替）
     git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
     git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
     git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
     go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
     git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
     cd $GOPATH/src/
     go install google.golang.org/grpc
     go get github.com/golang/protobuf/protoc-gen-go
     上面安装好后，会在GOPATH/bin下生成protoc-gen-go.exe
     但还需要一个protoc.exe，windows平台编译受限，很难自己手动编译，直接去网站下载一个，地址：https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0 ，同样放在GOPATH/bin下