# ADR-0001: Air-Gap Readiness

**Status:** Proposed

**Context:** Enterprise and security-conscious environments require tools that can operate without network connectivity. This "air-gapped" requirement prevents security vulnerabilities and ensures build determinism.

**Decision:** Drift will be designed to be fully air-gap ready. All third-party dependencies, including the D2 core engine, will be vendored directly within the repository. The build process will not require any network calls to fetch dependencies or other resources.

**Consequences:**
-   The repository size will be larger due to vendored dependencies.
-   Updating dependencies will be a manual process of replacing the vendored code.
-   Builds will be faster and more reliable, as they are not subject to network availability.
