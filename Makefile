backend-image:
	docker build --tag asthtc-backend backend

.PHONY: backend-container
backend-container:
	docker run -it -p 80:80 asthtc-backend