# ADR 0006: Diagram Governance

**Status:** Accepted

## Context
To ensure consistency and reproducibility, all diagrams must follow standard practices regarding source files, versioning, and location.

## Decision
- Drift is the official platform for diagrams.
- Source `.d2` files are versioned; generated `.svg` may be committed.
- Inline small diagrams in Markdown; larger diagrams in `docs/`.
- CI/CD validation ensures deterministic rendering.

## Consequences
- Enforces organizational consistency.
- Reduces reliance on external tools (Mermaid, Visio).
- Requires discipline in file management.
