# ADR 0003: Mermaid Compatibility Layer

**Status:** Accepted

## Context
Many users are familiar with Mermaid. Supporting its flowchart syntax lowers the adoption barrier. However, timeline-aware diagrams and other advanced Mermaid types are complex and optional for MVP.

## Decision
Implement a **preprocessor** to parse Mermaid flowcharts (`graph TD/LR/BT/RL`) and convert them into D2 canonical AST.

## Consequences
- Simplifies migration for existing Mermaid users.
- Relies on deterministic AST rules (ADR 0002).
- Timeline-aware diagrams are not included in MVP but may be supported later.
