.PHONY: build
build:
	GO111MODULE=on go build -o build/zsh-alias-matcher ./cmd/zsh-alias-matcher

.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: clean
clean:
	rm -f build/*
