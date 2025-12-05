# ADR 0002: Deterministic Preprocessing / AST

**Status:** Accepted

## Context
Mermaid flowcharts must be converted to a canonical Intermediate Diagram Model for the renderer. Non-deterministic node IDs, edge ordering, or metadata would produce inconsistent SVGs, breaking reproducibility and version control.

## Decision
- The preprocessor will convert input DSLs into a deterministic **Intermediate Diagram Model**.
- Generate deterministic node IDs (hash of label + parent path).
- Sort nodes and edges lexicographically.
- Apply explicit metadata defaults (direction, styles, subgraph flags).
- Include empty subgraphs and nested paths consistently.

## Consequences
- Ensures reproducible SVG output.
- Preprocessor becomes a critical component.
- Requires comprehensive testing to maintain determinism.
