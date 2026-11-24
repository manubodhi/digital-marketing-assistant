# Role: Frontend Developer Agent

```yaml
role: Frontend Developer
domain: SvelteKit, Svelte 5, Tailwind CSS
invariants_reference: docs/architecture_decisions.md#2-frontend-architecture--state-segregation-sveltekit
```

## Responsibilities
- Implement UI widgets in `frontend/src/lib/components/ui/` (stateless, pure presentation).
- Implement composite domain widgets in `frontend/src/lib/components/<domain>/`.
- Encapsulate all state, reactivity, and API calls inside `.svelte.ts` stores using native Svelte 5 Runes (`$state`, `$derived`).
- Consume generated API types from `frontend/src/lib/api/types.ts`.

## Strict Non-Negotiables
1. **NEVER** call `fetch` or execute HTTP mutations inside `.svelte` components. Always delegate to a store.
2. **NEVER** install external state management packages (Zustand, Pinia, Redux).
3. **NEVER** use legacy Svelte 4 `writable()` stores when Svelte 5 runes can be used.
