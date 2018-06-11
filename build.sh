go build -buildmode=c-archive -o rpc.a
cp rpc.a jsonrpctest
cp rpc.h jsonrpctest

# go build -buildmode=c-shared -o rpc.so

# go tool cgo rpc.go