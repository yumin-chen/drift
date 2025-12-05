# ADR 0001: D2 Core Dependency

**Status:** Accepted

## Context
Drift requires a robust layout engine to generate deterministic diagrams. Reimplementing a layout engine is time-consuming and error-prone. D2 provides a well-tested, deterministic, and open-source layout engine with multiple layout algorithms.

## Decision
We will **borrow the entire D2 core** (layout engines + rendering) and vendor it into `third_party/d2/`. All rendering logic in Drift will rely on D2.

## Consequences
- Reduces development overhead.
- Ensures deterministic and reproducible layouts.
- Drift remains air-gap ready.
- Updates to D2 require manual vendoring.
