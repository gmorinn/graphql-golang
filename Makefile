ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

DIR=$(notdir $(shell pwd))
export DIR

init:
	@echo -e "\n\t🔑\n"
	@go run $(GOROOT)/src/crypto/tls/generate_cert.go --host $(API_DOMAIN)

gen:
	@echo -e "\n\t🧠\n"
	@go run github.com/99designs/gqlgen generate

sql:
	@echo -e "\n\t🧠\n"
	@sqlc generate

dev:
	@echo -e "\n\t💣\n"
	@go run server.go

.PHONY: init gen dev sql

