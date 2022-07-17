local-http:
	INTERFACE=http air
local-http-internal:
	INTERFACE=http-internal air
build:
	go build -o ./src ./src/main.go
start:
	./src/main
start-http:
	INTERFACE=http-internal air
start-http-internal:
	INTERFACE=http-internal air
