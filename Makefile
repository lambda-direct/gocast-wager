.PHONY: build clean deploy

build:
	env GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/flipcoin api/flipcoin/index.go
	env GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/testluck api/testluck/index.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
