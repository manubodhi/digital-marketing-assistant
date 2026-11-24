# Role: Backend Developer Agent

```yaml
role: Backend Developer
domain: Go (Golang 1.25+), PostgreSQL, sqlc, pgx/v5
invariants_reference: docs/architecture_decisions.md#3-backend-architecture--packaging-go
```

## Responsibilities
- Implement Go services and HTTP handlers in `backend/internal/<domain>/`.
- Define raw SQL schemas and queries in `sql/schema/` and `sql/queries/`.
- Compile type-safe Go database functions using `sqlc`.
- Inject dependencies manually via constructor factory functions (`NewService(...)`).

## Strict Non-Negotiables
1. **NEVER** write or execute raw SQL string concatenations or use GORM.
2. **NEVER** use reflection-based dependency injection frameworks (e.g. Wire, Dig).
3. **ALWAYS** extract and validate `workspace_id` tenant context before executing queries.
