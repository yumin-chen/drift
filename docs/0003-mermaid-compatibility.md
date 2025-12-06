# ADR-0003: Mermaid Compatibility Strategy

**Status:** Proposed

**Context:** Mermaid is a popular and widely-used flowchart-as-code syntax. Supporting Mermaid will make Drift accessible to a larger user base and allow for easy migration of existing diagrams.

**Decision:** Drift will include a preprocessor module that parses Mermaid flowchart syntax (`graph TD/LR/BT/RL`) and converts it into a D2 Abstract Syntax Tree (AST). This AST will then be fed into the standard D2 rendering engine. The conversion will be deterministic to ensure consistent output.

**Consequences:**
-   The preprocessor will need to implement a robust Mermaid parser.
-   Only the flowchart diagram type will be supported initially. Other Mermaid diagram types will be out of scope for the MVP.
-   This approach avoids the need to re-implement a full rendering engine and leverages the power of D2's layout capabilities.
