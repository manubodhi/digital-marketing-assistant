# Technical Architecture & Standards Registry (OKF Index)

```yaml
id: DOC-INDEX
type: Registry-Index
status: Active
last_updated: 2026-11-24
```

---

## Document Registry

| Document ID | Title | Domain | Status | Key Topics |
| :--- | :--- | :--- | :--- | :--- |
| [**DOC-001**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/tech_stack_decision.md) | Technology Stack Rationale | Infrastructure & Core Stack | `Accepted` | SvelteKit, Go runtime, PostgreSQL, Rationale & Trade-offs |
| [**DOC-002**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/architecture_decisions.md) | System Architecture & Specifications | Architecture, State, Persistence | `Accepted` | Svelte 5 Runes state segregation, Go layer isolation, `sqlc` + `pgx/v5`, OpenAPI contract sharing, Day-1 Testing Invariant |
| [**DOC-003**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/GRAPH.md) | Architectural Knowledge Graph | Topology & Node-Edge Registry | `Accepted` | Graphify traversal paths, Entity mapping (Auth, Workspace, BrandKit) |
| [**DOC-004**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/ui_ux_design_system.md) | UI/UX Design System & Theme Specification | Frontend Presentation, Theming | `Accepted` | Theme tokens (dark/light), typography scale, widget state mandates |
| [**DOC-005**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/system_design_hld_lld.md) | High-Level (HLD) & Low-Level Design (LLD) | System Architecture, Package Internals | `Accepted` | Multi-tenant request lifecycle, Go package internals, response envelopes |
| [**DOC-006**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/database_schema.md) | Database Schema & Entity Relationships | Persistence & PostgreSQL DDL | `Accepted` | Production DDL script, indexes, constraints, ERD |
| [**DOC-007**](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/milestones_wbs.md) | Work Breakdown Structure & Milestones | Project Management & Execution Order | `Active` | Granular tasks (Task 1.1 to 6.5) with acceptance criteria |

---

## Architectural Navigation Matrix

### 1. Invariants & Rules
* [System Invariants (Non-negotiable Rules)](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/architecture_decisions.md#1-system-invariants-non-negotiable-rules)
* [Specialized Team Roster & Workflow Rules](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/AGENTS.md)

### 2. Design System & Theme
* [Color Tokens & Light/Dark Theme](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/ui_ux_design_system.md#2-color-palette--semantic-tokens)
* [Typography & Spacing Standards](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/ui_ux_design_system.md#3-typography-scale)

### 3. System Design & Topology
* [Global System Topology (HLD)](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/system_design_hld_lld.md#11-global-system-topology)
* [Multi-Tenant Request Lifecycle & Security Boundary](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/system_design_hld_lld.md#12-multi-tenant-request-lifecycle--security-boundary)
* [Low-Level Package Internals (LLD)](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/system_design_hld_lld.md#2-low-level-design-lld)

### 4. Database & Persistence
* [PostgreSQL DDL Schema Script](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/database_schema.md#2-production-postgresql-ddl-script)
* [sqlc Compilation Workflow](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/architecture_decisions.md#4-persistence-specification-sqlc--postgresql-pgxv5)

### 5. Task Roadmap & WBS
* [Granular Milestone Tasks](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/milestones_wbs.md#2-granular-task-breakdown-by-milestone)
