.PHONY: all
all:
	CGO_ENABLED=0 go build -trimpath -o unikornctl main.go
