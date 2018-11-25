.PHONY: grpc-go-source
grpc-go-source: go-tools-service-source go-reverse-proxy-source

.PHONY: go-tools-service
go-tools-service-source:
	protoc -Iproto/toolsService \
	-I/usr/local/include \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:proto/toolsService \
	tools.proto

.PHONY: rest-reverse-proxy
go-reverse-proxy-source:
	protoc -Iproto/toolsService \
	-I/usr/local/include  \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/toolsService/tools.yaml:proto/toolsService \
	tools.proto

	# swagger definitions too
	protoc -Iproto/toolsService \
	-I/usr/local/include \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:proto/toolsService \
	tools.proto

.PHONY: images
images: backend-image reverse-proxy-image

.PHONY: backend-image
backend-image:
	# note: backend docker image needs backend/ and proto/
	docker build --tag asthtc-backend -f backend/Dockerfile .

.PHONY: reverse-proxy-image
reverse-proxy-image:
	# note: reverse proxy docker image needs reverse_proxy/ and proto/
	docker build --tag asthtc-reverse-proxy -f reverse_proxy/Dockerfile .

.PHONY: reverse-proxy-run
reverse-proxy-run:
	docker run -it -p 8080:8080 asthtc-reverse-proxy

.PHONY: backend-container
backend-run:
	docker run -it -p 80:80 asthtc-backend