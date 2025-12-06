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
	// 1. Determine if the input is Mermaid syntax or a dynamic data structure (e.g., JSON).
	// 2. If Mermaid, parse it to a D2 AST using the preprocessor.
	// 3. If dynamic data, transform it into a D2 AST.
	// 4. Render the D2 AST to SVG using the renderer.
	// 5. Return the SVG string.

	return "Hello from Go WASM!" + input
}

func main() {
	js.Global().Set("Drift", js.ValueOf(map[string]interface{}{
		"render": js.FuncOf(render),
	}))

	// Prevent the Go program from exiting.
	select {}
}
