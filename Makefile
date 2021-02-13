swagger:
	swagger generate spec -o ./swagger.yaml --scan-models
docker:
	echo docker build -t goservice .
	docker run -p 9090:9090 goservice
