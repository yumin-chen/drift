//go:build wasm && js

package main

import (
	"syscall/js"
)

// render is the exported function that will be called from JavaScript.
func render(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return ""
	}
	input := args[0].String()

	// TODO: Implement the full rendering pipeline:
	// 1. Deserialize the input string into a MotionGraph.
	// 2. Normalize the MotionGraph to a D2 AST.
	// 3. Render the D2 AST to SVG.
	// 4. Return the SVG string.

	return "Hello from Go WASM!" + input
}

func main() {
	js.Global().Set("Drift", js.ValueOf(map[string]interface{}{
		"render": js.FuncOf(render),
	}))

	// Prevent the Go program from exiting.
	select {}
}
