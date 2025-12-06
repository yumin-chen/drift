# Drift — Deterministic Renderer With Dynamic Data Inputs

Drift is fundamentally a deterministic diagram rendering engine. Its primary responsibility is:

`Input → Normalize → Generate D2 → Layout (deterministic) → SVG/PNG/PDF`

However, Drift does not fetch dynamic data. **It consumes dynamic data.** Dynamic data comes from MotionPlatform’s upstream modules (e.g., MotionGraph, MotionPatch, MotionFetch, MotionPolicy).

**NOTE:** Due to limitations in the current environment, the D2 dependency has not been vendored. To make this project fully air-gap ready, you will need to manually vendor the D2 source code into the `third_party/d2` directory.

## Key Characteristics

*   **✔ Deterministic:** Same input graph → same D2 → same layout → identical output.
*   **✔ Dynamic‑Data Compatible:** The graph fed into Drift can be updated in real time by other MotionPlatform modules.
*   **✘ Not stateful:** Drift never maintains subscriptions, streams, or polling. It only renders the graph it is given at that moment, with full determinism.

## How Dynamic Rendering Works in MotionPlatform

1.  **Upstream module updates a graph snapshot.**
    *   Examples: A Git graph updates via `MotionFetch`, a dependency graph updates via `MotionGraph`, a dashboard updates via `MotionMetrics`, or a filtered graph updates via `MotionPolicy`.
2.  **MotionPlatform passes the updated internal graph to Drift.**
3.  **Drift converts that to canonical D2.**
    *   This includes deterministic node/edge sorting, stable IDs, and stable styles.
4.  **The WASM module produces a deterministic layout,** even if the graph is based on changing runtime data, was fetched from an external system, or is re-rendered frequently.

## Key Insight for Documentation

> Drift is **deterministic in rendering, not static in data.**

It is only “static” in the sense that it renders exactly what it is given. It becomes “dynamic” when the input graph is dynamically updated by other modules. This allows for dynamic dashboards without compromising determinism.

## Architecture

Drift's architecture is designed for modularity. The core components include:

-   **Preprocessor:** A normalizer that converts an internal MotionPlatform graph into a D2 Abstract Syntax Tree (AST).
-   **Renderer:** A module that takes a D2 AST, selects a layout engine, and generates an SVG.
-   **CLI:** A command-line interface for testing rendering from local files.
-   **WASM:** A WebAssembly module that exposes Drift's rendering capabilities to the browser for use in MotionPlatform.

![Drift Architecture](docs/architecture.svg)

## Directory Structure

```
drift/
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── third_party/d2/
├── cmd/drift/main.go
├── pkg/preprocessor/parser.go
├── pkg/renderer/renderer.go
├── wasm/drift_wasm.go
├── wasm/build.sh
└── docs/
    ├── architecture.d2
    └── architecture.svg
```

## Architecture Decision Records (ADRs)

-   [ADR-0001: Air-Gap Readiness](./docs/0001-air-gap.md)
-   [ADR-0002: Dynamic Layout Engine Selection](./docs/0002-dynamic-layout-engine.md)
-   [ADR-0003: Mermaid Compatibility Strategy](./docs/0003-mermaid-compatibility.md) (Superseded)
-   [ADR-0004: Browser and WASM Support](./docs/0004-browser-wasm.md)
-   [ADR-0005: Timeline-Aware Diagrams (Future)](./docs/0005-timeline-awareness.md)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
