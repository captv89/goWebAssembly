wasm:
	GOOS=js GOARCH=wasm go build -o assets/json.wasm cmd/wasm/main.go

run:
	go run cmd/wasm/main.go

copywasmjs:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/Go-Projects/goWebAssembly/assets

serve:
	go run cmd/server/main.go

buildnserve:
	make wasm
	make serve
