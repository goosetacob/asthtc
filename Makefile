.PHONY: go-rpc
go-rpc:
	protoc -I api tool.proto --go_out=plugins=grpc:api

backend-image:
	# note: backend docker image needs backend/ and api/
	docker build --tag asthtc-backend -f backend/Dockerfile .

.PHONY: backend-container
backend-container:
	docker run -it -p 80:80 asthtc-backend