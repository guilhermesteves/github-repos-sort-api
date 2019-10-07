.PHONY: build clean deploy

build:
	dep ensure -v
    GOOS=linux go build -ldflags="-s -w" -o build/like cmd/like/main.go
    GOOS=linux go build -ldflags="-s -w" -o build/repos cmd/repos/main.go

clean:
	rm -rf ./build ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose