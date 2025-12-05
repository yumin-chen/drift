# ADR 0009: Extensible Architecture

**Status:** Accepted

## Context
The initial Drift architecture was tightly coupled to the D2 AST. To support richer diagram types (e.g., timelines, UML) and new input DSLs (e.g., PlantUML) without breaking backward compatibility, a more extensible, layered architecture is required.

## Decision
We will adopt a multi-layered architecture centered around a new **Intermediate Diagram Model (IDM)**.

**[Input Layer] -> [Preprocessor Layer] -> [Intermediate Diagram Model] -> [Model Mapper] -> [Renderer Layer] -> [Output Layer]**

1.  **Intermediate Diagram Model (IDM):** A superset of the D2 AST that captures additional semantics, such as lifelines, timeline data, and UML constructs. This model serves as the standard, internal representation for all diagrams, regardless of their source format.

2.  **Preprocessor Layer:** Responsible for converting various input DSLs (like Mermaid or PlantUML) into the IDM. Each supported DSL will have its own preprocessor, ensuring that the core engine only has to deal with the standardized IDM.

3.  **Model Mapper:** A new component that inspects the IDM and decides how to render it.
    *   If the IDM contains only D2-compatible features, it will be mapped directly to the D2 AST for standard rendering.
    *   If the IDM contains extended features, it will be routed to an **Extended Renderer Adapter** that can handle the richer semantics.

## Consequences
- **Clean Separation of Concerns:** The architecture clearly separates input parsing from diagram rendering.
- **Extensibility:** New DSLs can be supported by simply adding a new preprocessor. Richer diagram types can be supported by extending the IDM and implementing a corresponding renderer adapter.
- **Backward Compatibility:** Standard D2 and Mermaid diagrams will continue to be rendered by the core D2 engine, ensuring no regressions.
- **Increased Complexity:** The introduction of the IDM and Model Mapper adds complexity to the system.
- **Potential for Divergence:** The Extended Renderer must be carefully managed to maintain the deterministic guarantees of the core engine.

## System Strengths
- Supports both legacy diagrams and future richer diagrams.
- Deterministic and reproducible.
- Air-gap ready and CI/CD friendly.

## Risks and Considerations
- The preprocessor layer may become a bottleneck if too many DSLs are supported.
- Good documentation will be required for developers who wish to write extensions or plugins.
