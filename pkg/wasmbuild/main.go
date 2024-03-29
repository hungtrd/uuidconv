package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/hungtrd/uuidconv/pkg/uuid"
)

// Build command: GOARCH=wasm GOOS=js go build -o main.wasm main.go

// making a webassembly compatible json function
func uuidConvert() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid number of argument"
		}
		u, err := uuid.NewFromUnknow(args[0].String())
		if err != nil {
			return err.Error()
		}
		resp := map[string]string{
			"uuid":   u.NormalString,
			"base64": u.Base64Encoded,
			"base62": u.Base62Encoded,
		}
		json, err := json.Marshal(resp)
		if err != nil {
			return err.Error()
		}

		return string(json)
	})
}

func uuidNew() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		u := uuid.New()
		resp := map[string]string{
			"uuid":   u.NormalString,
			"base64": u.Base64Encoded,
			"base62": u.Base62Encoded,
		}
		json, err := json.Marshal(resp)
		if err != nil {
			return err.Error()
		}
		return string(json)
	})
}

func registerCallbacks() {
	js.Global().Set("uuidConvert", uuidConvert())
	js.Global().Set("uuidNew", uuidNew())
}

func main() {
	fmt.Println("running go as webassembly")
	registerCallbacks()

	// add the channel and call it, so that our golang program does not close otherwise we will get
	// Uncaught (in promise) Error: Go program has already exited
	<-make(chan bool)
}
