//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}

func jsonWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments"
		}
		input := args[0].String()
		pretty, err := prettyJson(input)
		if err != nil {
			return err.Error()
		}
		return pretty
	})
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}
