port = 8089
host = 127.0.0.1


run_local:
	@# $ make run_dev [port=<number>]
	@# Examples:
	@# $ make run_local port=9393
	go run app/main.go -port=$(port) -host=$(host)

run_prod:
	@# $ make run_prod [port=<number>]
	@# Examples:
	@# $ make run_prod port=9393 workers=3
	 make build && ./main --port=$(port) -host=$(host)

build:
	@# $ make build
	go build app/main.go