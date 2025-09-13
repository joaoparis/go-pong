AIR_VERSION := $(shell go list -m -f '{{.Version}}' github.com/air-verse/air)

dev:
	go run github.com/air-verse/air@$(AIR_VERSION)
