
FILES = $(shell find . -type f -name '*.go')


gofmt:
	@actools gofmt -w $(FILES)
	@actools gofmt -r '&α{} -> new(α)' -w $(FILES)
