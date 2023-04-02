.PHONY: build test clean deploy start-local

GO_CODE_DIR := server
OUT_DIR := out


build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $(OUT_DIR)/get_user ${GO_CODE_DIR}/cmd/get_user.go

test:
	go test -v ./${GO_CODE_DIR}/...

clean:
	rm -rf ./$(OUT_DIR)

deploy: clean build
	sls deploy --verbose

start-local:
	sls offline
