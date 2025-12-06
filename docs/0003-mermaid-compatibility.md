# ADR-0003: Mermaid Compatibility Strategy

**Status:** Superseded

**Context:** The initial proposal was to support Mermaid as a primary input format to make Drift accessible to a wider user base. The goal was to parse Mermaid syntax directly and convert it into a D2 AST.

**Decision:** This strategy has been superseded. Drift's primary input is now a normalized internal graph provided by upstream MotionPlatform modules. While those modules may themselves parse Mermaid or other formats, Drift's core responsibility is to deterministically render the graph it is given, not to parse various external syntaxes.

**Consequences:**
-   The preprocessor's role shifts from a Mermaid parser to a graph normalizer that produces a canonical D2 AST.
-   This simplifies Drift's focus, concentrating its efforts on deterministic rendering and layout, leaving the complexities of parsing external formats to other, more specialized components within MotionPlatform.
