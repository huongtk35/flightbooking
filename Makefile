gen-files:
	protoc -Icommon/proto --go_out=. \
	--go_opt=module=github.com/huongtk35/flightbooking \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/huongtk35/flightbooking \
	common/proto/*.proto
	#go build -o bin/greet/server ./greet/server
	#go build -o bin/greet/client ./greet/client