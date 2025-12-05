package main

import "syscall/js"

// Render is the WASM entrypoint for rendering a diagram
func Render(this js.Value, args []js.Value) interface{} {
	// TODO: Implement WASM rendering logic
	return "TODO: Implement WASM rendering"
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("Drift", js.ValueOf(map[string]interface{}{
		"render": js.FuncOf(Render),
	}))
	<-c
}
