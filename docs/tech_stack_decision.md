# DOC-001: Technology Stack Rationale

```yaml
id: DOC-001
type: Architecture-Decision-Record
status: Accepted
domain: Infrastructure & Stack Selection
revision: 1.1
last_updated: 2026-11-24
```

---

## 1. System Context & Requirements
The platform is a multi-tenant digital marketing assistant automating:
- Organic content production (festive/daily posters, copywriting, video rendering).
- Multi-channel paid campaigns (Meta Ads, Google Ads, WhatsApp Click-to-Chat).
- Social scheduling and campaign performance analytics.

Core architectural constraints:
- **Lightweight Runtime Footprint:** Low memory overhead and fast execution across cloud deployments.
- **Pragmatic Simplicity:** Avoid complex distributed abstractions, premature microservices, and reflection-heavy frameworks.
- **Asynchronous Execution:** High concurrency support for cron scheduling, external API synchronization, and webhook ingestion.

---

## 2. Decision Matrix

| Layer | Selected Technology | Primary Rationale | Rejected Alternatives |
| :--- | :--- | :--- | :--- |
| **Frontend** | **SvelteKit (Svelte 5)** | Zero-runtime compilation (~20KB baseline bundle), fine-grained reactivity via Runes, low CPU consumption. | **Next.js / React:** Heavy virtual DOM runtime (~150KB+), hook dependency complexity (`useEffect`), higher re-rendering overhead. |
| **Backend** | **Go (Golang 1.25+)** | Single standalone binary (~15–25MB idle memory), zero external runtime dependencies, native goroutine concurrency for scheduling tasks. | **Node.js / Python:** High memory footprint under heavy loads, runtime drift. <br>**Rust:** High syntactic complexity and slow compilation cycles for rapid API iteration. |
| **Database** | **PostgreSQL** | Robust relational consistency, JSONB support for dynamic ad payload storage, mature ecosystem. | **MongoDB / DynamoDB:** Lacks strict relational constraints needed for multi-tenant isolation. |

---

## 3. Detailed Architectural Justifications

### 3.1 Frontend: SvelteKit with Svelte 5 Runes
1. **Compilation Model:** Svelte transforms declarative markup into direct DOM manipulation code during build time, removing the virtual DOM reconciliation layer.
2. **Reactivity Model:** Svelte 5 Runes (`$state`, `$derived`, `$props`) allow reactive state encapsulation outside UI components inside plain `.svelte.ts` modules without requiring third-party state managers.
3. **Standard Web Primitives:** Native adherence to `Request`, `Response`, and standard `fetch` pipelines reduces vendor abstraction overhead.

### 3.2 Backend: Go Service Architecture
1. **Concurrency Model:** Goroutines require ~2KB stack memory per routine, allowing thousands of simultaneous background tasks (e.g., ad status polling, WhatsApp batch queues) to run concurrently on minimal infrastructure.
2. **Operational Simplicity:** The compiled binary contains all dependencies, enabling deployment inside scratch or distroless Docker containers (<30MB total image size).
3. **Type Rigidity:** Compile-time static analysis catches contract mismatches early without requiring heavy runtime type guards.

---

## 4. Consequences & Operational Boundary
- **Independent Services:** Frontend and backend run as decoupled processes communicating over HTTP/JSON.
- **Contract Synchronization:** API contracts must be explicitly maintained using OpenAPI 3.1 definitions (defined in [DOC-002](file:///Users/manoranjan/Projects/Learning/digital-marketing-assistant/docs/architecture_decisions.md)).
