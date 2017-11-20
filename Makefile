DOCKER="shawncatz/hanabi"

run:
	gin run

deps:
	go get -u github.com/streadway/amqp
	go get -u github.com/hashicorp/consul/api

linux:
	GOOS=linux GOARCH=amd64 go build -o hanabi

docker:
	docker build -t $(DOCKER) .

docker-run:
	docker run --net=host --name test --rm $(DOCKER)

docker-push:
	docker push $(DOCKER)

deploy:
	sup prod deploy