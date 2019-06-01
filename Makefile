
FILES = $(shell find . -type f -name '*.go')


.PHONY: testdata

gofmt:
	@actools gofmt -w $(FILES)
	@actools gofmt -r '&α{} -> new(α)' -w $(FILES)

testdata:
	actools protoc -I/opt/googleapis -I. --go_out=plugins=grpc,paths=source_relative:. testdata/foo/foo.proto
	actools protoc -I/opt/googleapis -I. --include_imports --include_source_info --descriptor_set_out=testdata/foo/foo.pb testdata/foo/foo.proto

test:
	actools go install ./cmd/validate-proto-http
	actools run go validate-proto-http
