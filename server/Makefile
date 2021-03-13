DOCKER="shawncatz/hanabi-server"

run:
	gin run

deps:
	go mod tidy
	go get -u github.com/codegangsta/gin
	#go get -u github.com/pressly/sup/cmd/sup

linux:
	GOOS=linux GOARCH=amd64 go build -o hanabi-server

docker:
	docker build -t $(DOCKER) .

docker-run:
	docker run --net=host --name test --rm $(DOCKER)

docker-push:
	docker push $(DOCKER)
