# ADR 0004: Dynamic Layout Engine Selection

**Status:** Accepted

## Context
Different diagram types and sizes require different layout engines for optimal performance and readability.

## Decision
Renderer will dynamically select a layout engine based on diagram type and complexity:
- DAGRE → small/simple diagrams
- TALA → tree-like diagrams
- ELK → large, highly nested diagrams

## Consequences
- Optimizes performance and layout quality.
- Engine selection logic must remain deterministic.
- Future diagram types may require updates to selection rules.
