# Drift

**Tagline:** A deterministic, air-gap-ready flowcharts-as-code engine.

## Overview

Drift is a powerful, extensible tool that allows you to create a variety of diagrams from simple, text-based syntaxes. It is built on top of the excellent D2 engine and is designed to be:

*   **Deterministic:** The same input will always produce the same output.
*   **Air-Gap Ready:** All dependencies are vendored, so no network calls are needed at build time.
*   **Mermaid Compatible:** Supports Mermaid flowchart syntax for easy migration.
*   **Extensible:** A new Intermediate Diagram Model allows for future DSLs and richer diagram semantics without breaking core functionality.
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

The architecture of Drift is designed to be a clean, extensible pipeline that separates concerns between input, processing, and rendering.

**[Input Layer] --> [Preprocessor Layer] --> [Intermediate Diagram Model] --> [Model Mapper] --> [Renderer Layer] --> [Output Layer]**

*   **Input Layer:** Accepts various diagram sources, such as `.d2` files and Mermaid flowcharts.
*   **Preprocessor Layer:** Converts the input DSL into a standardized, deterministic **Intermediate Diagram Model**.
*   **Intermediate Diagram Model:** A superset of the D2 AST that can capture richer semantics, such as timeline data or UML constructs.
*   **Model Mapper:** Maps the Intermediate Diagram Model to either the standard D2 AST for backward-compatible rendering or to an Extended Renderer for advanced diagram types.
*   **Renderer Layer:** Uses powerful layout engines like ELK, TALA, and DAGRE to convert the AST into a visual representation.
*   **Output Layer:** Produces the final SVG, PNG, or other output formats.

This layered design ensures that new diagram types and DSLs can be added in the future without disrupting the core rendering engine. A more detailed diagram of the architecture can be found in `docs/architecture.svg`.

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
*   [0009-extensible-architecture.md](docs/adr/0009-extensible-architecture.md)
