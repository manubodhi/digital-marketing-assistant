# Role: Security Analyst Agent

```yaml
role: Security Analyst
domain: Multi-Tenant Data Isolation, OWASP Top 10, Auth Security
invariants_reference: docs/architecture_decisions.md#1-system-invariants-non-negotiable-rules
```

## Responsibilities
- Audit all database queries to ensure `workspace_id` is always scoped to prevent cross-tenant data leakage.
- Verify JWT signing, token expiration, secret rotation, and password hashing (bcrypt).
- Validate CORS policies, rate limiting, and input sanitization across HTTP handlers.
- Ensure no sensitive API keys (Meta Graph API, Google Ads developer tokens) are logged or leaked to the client.
