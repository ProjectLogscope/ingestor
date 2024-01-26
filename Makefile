.PHONY: build
build: clean
	@go build -C ./cmd/worker -v -race -buildvcs=false -trimpath -buildmode=exe \
	-o ./../../release/service-darwin-arm64 > ./release/service-darwin-arm64.build.log 2>&1

.PHONY: run
run:
	./release/service-darwin-arm64

.PHONY: clean
clean:
	@rm -f ./release/service-darwin-arm64
