# Project: AI Digital Marketing Assistant

This repository contains a multi-tenant digital marketing SaaS platform automating organic creative production (posters, copy, reels), paid ad campaigns (Meta, Google, WhatsApp), and social scheduling.

---

## 1. Core Architecture Index

| Component | Technology | Architectural Invariants |
| :--- | :--- | :--- |
| **Frontend** | SvelteKit (Svelte 5) + Tailwind CSS | **Widget-Based Segregation:** UI widgets (`src/lib/components/`) MUST NOT call `fetch` directly. All state and API interactions are encapsulated in `.svelte.ts` stores using native Runes (`$state`, `$derived`). |
| **Backend** | Go (1.25+) | **Pragmatic Layered by Domain:** Packages live in `internal/<domain>/` (`handler.go`, `service.go`, `repository.go`). Dependencies MUST be injected via constructor functions (`NewService(...)`). Zero reflection DI containers. |
| **Database** | PostgreSQL + `sqlc` + `pgx/v5` | **Zero-Reflection Persistence:** Raw SQL string concatenation is prohibited. All queries reside in `sql/queries/*.sql` and are compiled to type-safe Go functions via `sqlc`. |
| **Contracts** | OpenAPI 3.1 + `openapi-typescript` | **Single Source of Truth:** `api-spec/openapi.yaml` defines all API endpoints. TypeScript types are generated via `npm run codegen`. |

---

## 2. Documentation Directory Map (OKF)

All deep technical specifications and rationale are documented in the `docs/` directory:
- [docs/README.md](docs/README.md): Master Navigation Index & Invariant Map
- [docs/tech_stack_decision.md](docs/tech_stack_decision.md): **DOC-001** (Stack Selection Rationale)
- [docs/architecture_decisions.md](docs/architecture_decisions.md): **DOC-002** (State Management, sqlc & System Invariants)
- [docs/GRAPH.md](docs/GRAPH.md): **DOC-003** (Topological Knowledge Graph & Node Traversal Registry)

---

## 3. Specialized Agent Roster (`.agents/agents/`)

The team workflow is divided across 7 specialized roles:
1. [ui-ux-designer](.agents/agents/ui-ux-designer.md): Design system, theme tokens, WCAG accessibility, user flow.
2. [frontend-dev](.agents/agents/frontend-dev.md): SvelteKit, Svelte 5 runes, widget segregation, zero fetch in UI.
3. [backend-dev](.agents/agents/backend-dev.md): Go 1.25+, PostgreSQL, `sqlc`, `pgx/v5`, domain isolation.
4. [qa-engineer](.agents/agents/qa-engineer.md): API contract validation, unit tests, regression suites.
5. [security-analyst](.agents/agents/security-analyst.md): Multi-tenant isolation audit, JWT auth, RBAC.
6. [code-reviewer](.agents/agents/code-reviewer.md): Invariant enforcement, Go/Svelte best practices review.
7. [product-owner](.agents/agents/product-owner.md): Scope governance, milestone tracking against requirements.

---

## 4. Rules for AI Agents Working in This Repository

1. **Check Invariants First:** Before generating or editing code, verify compliance with [docs/architecture_decisions.md](docs/architecture_decisions.md#1-system-invariants-non-negotiable-rules).
2. **Day-1 Testing Invariant (Mandatory):** Every feature change MUST include automated tests before it is marked complete:
   - Backend: Every service method MUST have unit tests in `*_test.go` using mock repositories.
   - Frontend: Every `.svelte.ts` store MUST have unit tests in `*.test.ts` using Vitest.
3. **Automated Team Workflow (Definition of Done):**
   - Step 1: `frontend-dev` or `backend-dev` writes the feature code AND its corresponding unit test.
   - Step 2: `qa-engineer` runs the test suite (`go test ./...`, `npm run test`) and verifies contract synchronization.
   - Step 3: `security-analyst` audits the changes for `workspace_id` tenant isolation and auth bounds.
   - Step 4: `code-reviewer` validates against `DOC-001` and `DOC-002` invariants before git commit.
4. **Never Add Redundant State Libraries:** Do not install Pinia, Zustand, Redux, or Svelte 4 store wrappers. Use Svelte 5 Runes in `.svelte.ts` files.
5. **Never Use GORM or Raw SQL Concatenation:** All database queries must be written in standard SQL inside `sql/queries/` and compiled with `sqlc`.
6. **Maintain Contract Synchronization:** If a backend endpoint is modified, update `api-spec/openapi.yaml` and regenerate the frontend types.
