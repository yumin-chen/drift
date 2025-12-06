# ADR-0005: Timeline-Aware Diagrams

**Status:** Future

**Context:** Flowcharts and diagrams often represent processes that change over time. Standard diagramming tools lack a native understanding of timelines, making it difficult to visualize evolution, versioning, or sequential states.

**Decision:** As a future feature, Drift will explore the concept of "timeline awareness." This would involve extending the D2 syntax to support metadata that defines how diagram elements appear, disappear, or change at different points in a timeline. The renderer would then be able to generate interactive SVGs or a sequence of diagrams that visualize this evolution.

**Consequences:**
-   This will require significant research and development, including potential modifications to the D2 core.
-   The user interface (CLI and WASM) will need to be updated to support timeline-based rendering.
-   This feature would provide a unique and powerful capability not found in other diagramming tools.
