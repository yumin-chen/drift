# ADR 0008: Future-Proofing / Optional Features

**Status:** Proposed

## Context
Advanced diagram types (timeline-aware, interactive SVGs) are desirable but not critical for MVP.

## Decision
- Timeline-aware diagrams, interactive SVGs, and advanced Mermaid types are **optional future features**.
- Architecture allows independent evolution of preprocessor, renderer, and layout engines.

## Consequences
- MVP remains achievable.
- Extensions can be added without affecting foundational components.
- Requires careful planning to maintain determinism when features are added.
