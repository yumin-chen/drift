# Drift

**Tagline:** Deterministic, Air-Gap-Ready Diagrams-as-Code

Drift is a powerful, self-contained engine that transforms D2 and Mermaid definitions into beautiful, consistent SVG diagrams. It is designed for security-conscious, air-gapped environments and supports both static file rendering and dynamic data inputs.

---

## Key Features

- **Deterministic Rendering:** Identical input (static file or dynamic AST) produces the exact same SVG every time.
- **Air-Gap Ready:** All dependencies, including fonts and layout engines, are vendored. No network calls are made at build or run time.
- **Mermaid Compatibility:** The preprocessor maps Mermaid Flowcharts, Class Diagrams, and ER Diagrams to a canonical D2 AST.
- **Dynamic Data Layer:** Consumes generic dynamic inputs (JSON, API responses, etc.) that can be mapped to a D2 AST by the WASM/JS layer.
- **Stateless Operation:** Drift only renders the AST it is given at any moment, without maintaining subscriptions, streams, or polling.
- **CLI & WASM Support:** Use Drift as a standalone CLI for offline rendering or as a WASM module in any browser environment.

---

## Supported Diagram Types

### v0.1 (MVP)

| Mermaid Type / Diagram           | Status    | Notes / Mapping to D2                                           |
| -------------------------------- | --------- | --------------------------------------------------------------- |
| Flowcharts (graph TD/LR/BT/RL)   | ✅ MVP    | Fully deterministic, DAG layout                                 |
| Class Diagrams                   | ✅ MVP    | Nodes = classes, edges = relationships, optional clusters for namespaces |
| ER Diagrams                      | ✅ MVP    | Nodes = entities, edges = relationships with optional cardinality, clusters optional |

### Future Roadmap

| Mermaid Type / Diagram           | Status    | Notes                                                           |
| -------------------------------- | --------- | --------------------------------------------------------------- |
| State Diagrams                   | ❌ Future | Requires extensions for loops, nested states, and conditional transitions. |
| Gantt / Timeline                 | ❌ Future | Requires timeline-aware metadata and sequencing.                |
| Sequence Diagrams                | ❌ Future | Requires additional preprocessor and layout logic for sequential messages. |

---

## Usage

### CLI

Render a local diagram file:

```bash
drift render my-diagram.d2 --output my-diagram.svg
drift render my-flowchart.mermaid --output my-flowchart.svg --engine ELK
```

### Browser (WASM)

```javascript
import drift from 'drift-wasm';

const mermaidSyntax = `
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
`;

const svgOutput = drift.render(mermaidSyntax);
document.getElementById('diagram-container').innerHTML = svgOutput;
```

---

## Architecture

Drift’s architecture is a self-contained pipeline: `Parser → Layout → Renderer → Output`.

1.  **Preprocessor:** Converts Mermaid syntax (Flowchart, Class, ER) into a canonical D2 Abstract Syntax Tree (AST).
2.  **Dynamic Data Layer (JS/WASM):** The JS/WASM layer is responsible for transforming dynamic data (from JSON, APIs, etc.) into a D2-compatible AST before passing it to the renderer.
3.  **Renderer:** Takes a D2 AST, selects a layout engine (DAGRE, TALA, or ELK), and generates an SVG.
4.  **Output:** The final SVG is delivered via the CLI (for files) or the WASM module (for browser rendering).

---

## Directory Structure

```
drift/
├── src/
│   ├── cli/
│   ├── wasm/
│   ├── preprocessor/
│   ├── renderer/
│   └── assets/
├── examples/
├── third_party/   # vendored D2 source
├── docs/
├── tests/
├── LICENSE
└── README.md
```

---

## Roadmap

- **v0.1 (MVP):** Flowchart, Class, and ER diagram support. Stable CLI and WASM rendering pipelines. Full air-gap compliance.
- **v0.2:** Formalize the dynamic data integration layer with generic JSON/API source examples.
- **v0.3+:** Implement preprocessor extensions for State, Gantt, and Sequence diagrams.
- **v0.4+:** Investigate timeline-aware rendering and optional SVG interactivity.

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
