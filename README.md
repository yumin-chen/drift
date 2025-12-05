# Drift

**Tagline:** A deterministic, air-gap-ready flowcharts-as-code engine.

## Overview

Drift is a powerful tool that allows you to create flowcharts and diagrams from a simple, text-based syntax. It is built on top of the excellent D2 engine and is designed to be:

*   **Deterministic:** The same input will always produce the same output.
*   **Air-Gap Ready:** All dependencies are vendored, so no network calls are needed at build time.
*   **Mermaid Compatible:** Supports Mermaid flowchart syntax for easy migration.
*   **Browser Friendly:** Can be compiled to WASM for use in the browser.

## Key Features

*   **Flowcharts as Code:** Define your diagrams in a simple, human-readable syntax.
*   **Multiple Layout Engines:** Choose between DAGRE, TALA, or ELK for diagram layout.
*   **SVG Output:** Renders diagrams to crisp, scalable SVG images.
*   **CLI and WASM Support:** Use Drift from the command line or in the browser.

## Usage Examples

### CLI

Render a D2 file to SVG:

```bash
drift render my-diagram.d2 my-diagram.svg
```

Render a Mermaid file to SVG:

```bash
drift render my-diagram.mermaid my-diagram.svg
```

### Browser (WASM)

```javascript
import drift from 'drift';

const mermaidSyntax = `
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
`;

const svg = drift.render(mermaidSyntax);

document.getElementById('diagram').innerHTML = svg;
```

## Architecture

The architecture of Drift is simple and modular. It consists of a preprocessor that converts Mermaid syntax to a D2 AST, a renderer that uses a layout engine to generate an SVG, and a CLI and WASM interface.

A more detailed diagram of the architecture can be found in `docs/architecture.d2`.

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

## ADRs

Architectural Decision Records are located in the `docs/adr` directory.

*   [0001-d2-core-dependency.md](docs/adr/0001-d2-core-dependency.md)
*   [0002-deterministic-preprocessing-ast.md](docs/adr/0002-deterministic-preprocessing-ast.md)
*   [0003-mermaid-compatibility-layer.md](docs/adr/0003-mermaid-compatibility-layer.md)
*   [0004-dynamic-layout-engine-selection.md](docs/adr/0004-dynamic-layout-engine-selection.md)
*   [0005-cli-and-wasm-interfaces.md](docs/adr/0005-cli-and-wasm-interfaces.md)
*   [0006-diagram-governance.md](docs/adr/0006-diagram-governance.md)
*   [0007-licensing-and-open-source-compliance.md](docs/adr/0007-licensing-and-open-source-compliance.md)
*   [0008-future-proofing-optional-features.md](docs/adr/0008-future-proofing-optional-features.md)
