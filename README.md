# AI Digital Marketing Assistant

[![Architecture: SvelteKit + Go](https://img.shields.io/badge/Stack-SvelteKit%20%2B%20Go%201.25-blue.svg)](docs/tech_stack_decision.md)
[![Database: PostgreSQL + sqlc](https://img.shields.io/badge/Database-PostgreSQL%20%2B%20sqlc-336791.svg)](docs/database_schema.md)
[![License: Proprietary](https://img.shields.io/badge/License-Proprietary-red.svg)](#)

A high-performance, multi-tenant digital marketing SaaS platform engineered to automate organic creative production (festival & daily promotional posters, copywriting, reels), paid ad campaigns (Meta, Google, WhatsApp), and social media scheduling for businesses, agencies, and retail brands.

---

## 1. What This Platform Does

Small businesses and digital agencies struggle with the daily churn of marketing asset production and cross-platform ad execution. This platform serves as an autonomous co-pilot:

- **AI Creative Studio:** Generate high-impact promotional posters (e.g. Diwali sales, product launches, seasonal offers) with programmatic branding (vector logos, brand colors, typography) and auto-generated copy.
- **Multi-Channel Paid Ads:** Launch paid ad campaigns directly into **Meta Ads** (Facebook & Instagram), **Google Ads** (Search & PMax), and **Click-to-WhatsApp Ads** with automated budget controls and audience targeting.
- **Organic Social Scheduler:** Auto-publish and schedule visual creatives across Instagram Business and Facebook Pages via a unified content calendar.
- **Multi-Tenant Workspace Management:** Agency-ready tenant isolation allowing teams to manage multiple client brand kits, assets, and ad accounts from a single dashboard.

---

## 2. Core Architecture & Tech Stack

The architecture is built on the principles of **lightweight footprint, minimal runtime overhead, and strict type safety**.

```mermaid
flowchart LR
    subgraph Frontend [Frontend: SvelteKit (Svelte 5)]
        UI[Widget-Based UI] --> Store[Svelte 5 Runes Stores]
        Store --> APIClient[Typed OpenAPI Client]
    end

    subgraph Backend [Backend: Go 1.25+]
        APIClient -->|HTTP / JSON| Router[go-chi Router]
        Router --> MW[JWT & Workspace Auth Middleware]
        MW --> Service[Domain Business Services]
        Service --> SQLc[sqlc Type-Safe Queries]
    end

    subgraph Persistence [Database]
        SQLc --> DB[(PostgreSQL)]
    end
```

| Layer | Technology | Architectural Invariants |
| :--- | :--- | :--- |
| **Frontend** | **SvelteKit (Svelte 5)** + Tailwind CSS | **Widget-Based Segregation:** UI widgets (`src/lib/components/`) never call `fetch` directly. All state and mutations reside in `.svelte.ts` stores using native Runes (`$state`, `$derived`). |
| **Backend** | **Go (Golang 1.25+)** | **Pragmatic Layered by Domain:** Modules live in `internal/<domain>/` (`handler.go`, `service.go`, `repository.go`). Explicit constructor injection (`NewService(...)`). Zero reflection DI containers. |
| **Database** | **PostgreSQL** + **`sqlc`** + **`pgx/v5`** | **Compile-Time Persistence:** Raw SQL string concatenation is prohibited. Pure SQL is written in `sql/queries/*.sql` and compiled to type-safe Go functions via `sqlc`. |
| **Contracts** | **OpenAPI 3.1** + `openapi-typescript` | **Single Source of Truth:** `api-spec/openapi.yaml` defines all endpoints. TypeScript types are generated via `npm run codegen`. |

---

## 3. Engineering Documentation Map (OKF)

The project follows the **Optimized Knowledge Framework (OKF)** for dense, machine-traversable specifications:

| Doc ID | Specification Document | Description |
| :--- | :--- | :--- |
| **INDEX** | [docs/README.md](docs/README.md) | Master Navigation Index & Invariant Map |
| **DOC-001** | [docs/tech_stack_decision.md](docs/tech_stack_decision.md) | Architectural Decision Record (ADR): SvelteKit + Go Selection Rationale |
| **DOC-002** | [docs/architecture_decisions.md](docs/architecture_decisions.md) | System Invariants, Svelte 5 State Segregation & `sqlc` Specification |
| **DOC-003** | [docs/GRAPH.md](docs/GRAPH.md) | Topological Knowledge Graph & Entity-to-File Node Traversal Registry |
| **DOC-004** | [docs/ui_ux_design_system.md](docs/ui_ux_design_system.md) | Design System, Dark/Light Theme Tokens, Typography Scale & Widget States |
| **DOC-005** | [docs/system_design_hld_lld.md](docs/system_design_hld_lld.md) | High-Level Design (HLD) & Low-Level Design (LLD) |
| **DOC-006** | [docs/database_schema.md](docs/database_schema.md) | Production PostgreSQL DDL Schema, Foreign Keys & Indexes |
| **DOC-007** | [docs/milestones_wbs.md](docs/milestones_wbs.md) | Work Breakdown Structure (WBS) & Task-Level Roadmap (Phases 1–6) |

---

## 4. Autonomous Agent Team & Quality Standards

The codebase enforces strict quality guardrails maintained by 7 specialized roles defined in [AGENTS.md](AGENTS.md) and [.agents/agents/](.agents/agents/):

1. **ui-ux-designer:** Design system, theme tokens, WCAG accessibility, user flow.
2. **frontend-dev:** SvelteKit, Svelte 5 runes, widget segregation, zero fetch in UI.
3. **backend-dev:** Go 1.25+, PostgreSQL, `sqlc`, `pgx/v5`, domain isolation.
4. **qa-engineer:** API contract validation, unit tests, regression suites.
5. **security-analyst:** Multi-tenant isolation audit, JWT auth, RBAC.
6. **code-reviewer:** Invariant enforcement, Go/Svelte best practices review.
7. **product-owner:** Scope governance, milestone tracking against requirements.

### Day-1 Testing Mandate (Definition of Done)
* **Backend:** Every Go service method requires corresponding unit tests in `*_test.go` using mock repositories.
* **Frontend:** Every Svelte store requires unit tests in `*.test.ts` using Vitest.
* **Workflow:** Code $\rightarrow$ QA Contract/Unit Test $\rightarrow$ Security Audit $\rightarrow$ Reviewer Sign-off $\rightarrow$ Git Commit.

---

## 5. Current Project Status

- [x] **Milestone 1:** Specifications, Theme & Architecture ([DOC-007](docs/milestones_wbs.md))
- [ ] **Milestone 2:** Scaffolding, Tooling & Database Migrations *(In Progress)*
- [ ] **Milestone 3:** Authentication & Multi-Tenant Workspaces (End-to-End)
- [ ] **Milestone 4:** Brand Kit Studio & Asset Management
- [ ] **Milestone 5:** AI Creative Engine (Posters, Banners, Copy)
- [ ] **Milestone 6:** Organic Social Scheduler & Paid Ad Launcher
