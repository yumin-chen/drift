# ADR-0004: Browser and WASM Support

**Status:** Proposed

**Context:** Many applications, such as wikis, documentation generators, and collaborative platforms, require in-browser diagram rendering. Providing a WebAssembly (WASM) module is the most effective way to deliver Drift's capabilities to the browser.

**Decision:** Drift will be compiled to WASM to provide a `render(string) → SVG` function that can be called from JavaScript. A dedicated Go file (`wasm/drift_wasm.go`) will serve as the entry point for the WASM build, and a `build.sh` script will be provided to automate the compilation process.

**Consequences:**
-   The core rendering logic must be self-contained and not rely on OS-specific features.
-   The WASM module will need to be published to a package registry (e.g., npm) for easy consumption.
-   This will enable a wide range of web-based integrations for Drift.
