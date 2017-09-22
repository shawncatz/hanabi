DOCKER="shawncatz/dashotv"

run:
	gin run

deps:
	dep ensure -update
	go get -u github.com/codegangsta/gin
	go get -u github.com/pressly/sup/cmd/sup

linux:
	GOOS=linux GOARCH=amd64 go build -o dashotv-api

docker:
	docker build -t $(DOCKER) .

docker-run:
	docker run --net=host --name test --rm $(DOCKER)

docker-push:
	docker push $(DOCKER)

deploy:
	sup prod deploy