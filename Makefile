.PHONY: build
build:
	GO111MODULE=on go build ./cmd/zsh-alias-matcher

.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: clean
clean:
	rm -f zsh-alias-matcher
