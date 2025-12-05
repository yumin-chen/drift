# ADR 0005: CLI and WASM Interfaces

**Status:** Accepted

## Context
Users need both offline CLI rendering for files and browser/WASM rendering for interactive or in-browser previews.

## Decision
- CLI: `drift render <input_file> -o <output_file>`
- WASM: `Drift.render(string) → SVG`

## Consequences
- CLI and WASM share core preprocessor and renderer logic.
- Enables integration in CI/CD pipelines and documentation previews.
