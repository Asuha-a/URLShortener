.PHONY: setup
setup:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/pb/user/user.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/pb/url/url.proto

.PHONY: clean
clean:
	rm -f api/pb/user/user.pb.go
	rm -f api/pb/user/user_grpc.pb.go
	rm -f api/pb/url/url.pb.go
	rm -f api/pb/url/url_grpc.pb.go

.PHONY: help
help:
	@echo "setup: setup the envs"
	@echo "clean: delete all files created by make"