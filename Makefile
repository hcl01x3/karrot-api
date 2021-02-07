PROGNAME := karrot-api
GQLGEN := github.com/99designs/gqlgen

help:
	@echo "help"

clean:
	@rm -f $(PROGNAME)

bootstrap:
	@go mod tidy

generate:
	@go generate ./*
	@go run $(GQLGEN) generate

run:
	@go run ./cmd/$(PROGNAME) --env-file .env

test:
	@go test -v

build:
	@go build ./cmd/$(PROGNAME)
