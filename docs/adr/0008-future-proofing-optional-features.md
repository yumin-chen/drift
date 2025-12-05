# ADR 0008: Future-Proofing / Optional Features

**Status:** Accepted

## Context
Advanced diagram types (timeline-aware, interactive SVGs) are desirable but not critical for MVP. A clear distinction between the initial scope and future extensions is needed to manage development effort.

## Decision
The following table defines the scope for the Minimum Viable Product (MVP) and outlines planned future extensions. The architecture, centered around the Intermediate Diagram Model, is designed to accommodate these future features without requiring a full rewrite.

| Layer / Component              | MVP Scope                  | Future Scope                                            |
| ------------------------------ | -------------------------- | ------------------------------------------------------- |
| **Input Layer**                | .d2 + Mermaid flowcharts   | PlantUML, custom DSLs                                   |
| **Preprocessor Layer**         | Mermaid → Intermediate Model | Other preprocessors for new DSLs                        |
| **Intermediate Diagram Model** | Basic D2 AST subset        | Superset for timelines, UML, statecharts                |
| **Model Mapper**               | Identity mapping (D2 AST)  | Map rich Intermediate Model → D2 AST / Extended Renderer |
| **Renderer Layer**             | D2 AST → SVG (ELK, TALA)   | Interactive / timeline-aware rendering                  |
| **Output Layer**               | SVG, PNG                   | Interactive HTML / enriched Markdown                    |

## Consequences
- MVP remains achievable and focused.
- Extensions can be added incrementally without affecting foundational components.
- The Intermediate Diagram Model is the key enabler for future enhancements.
- Requires careful planning to maintain determinism when new features are added.
