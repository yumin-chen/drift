# ADR-0002: Dynamic Layout Engine Selection

**Status:** Proposed

**Context:** Different diagram types and complexities benefit from different layout engines. D2 supports multiple layout engines, including DAGRE, TALA, and ELK. Forcing a single engine may produce suboptimal results for some diagrams.

**Decision:** Drift will support dynamic selection of the layout engine at runtime. The user will be able to specify the desired engine via a command-line flag (`--engine`) or a parameter in the WASM `render` function. The default engine will be DAGRE.

**Consequences:**
-   The rendering module will need to include logic to select the appropriate engine.
-   The CLI and WASM interfaces will need to expose the engine selection option.
-   Users will have more control over the final appearance of their diagrams.
