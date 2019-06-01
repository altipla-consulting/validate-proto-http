
FILES = $(shell find . -type f -name '*.go')

SRC_PROTOS = $(shell find testdata cmd -type f -name *.proto)
COMPILED_PROTOS := $(SRC_PROTOS:.proto=.pb.go)


.PHONY: testdata

gofmt:
	@actools gofmt -w $(FILES)
	@actools gofmt -r '&α{} -> new(α)' -w $(FILES)

test:
	actools go test ./...

protos: $(COMPILED_PROTOS)
	@gofmt -w $(FILES)
	@gofmt -r '&α{} -> new(α)' -w $(FILES)

%.pb.go: %.proto
	actools protoc -I/opt/googleapis -I. --go_out=plugins=grpc,paths=source_relative:. $<
	actools protoc -I/opt/googleapis -I. --include_imports --include_source_info --descriptor_set_out=$(<:.proto=.pb) $<
