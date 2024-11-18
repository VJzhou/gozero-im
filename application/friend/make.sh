# make rpc 
goctl rpc protoc friend.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# make api
goctl api go --api friend.api --dir . --home /Users/zhouweijie/go-workspace/im/template

#format api
goctl api format friend.api --dir . 