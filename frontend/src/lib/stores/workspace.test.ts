import { describe, it, expect, beforeEach, vi } from 'vitest';
import { WorkspaceStore } from './workspace.svelte';
import { apiClient } from '$lib/api/client';

describe('WorkspaceStore (Svelte 5 Runes)', () => {
	let store: WorkspaceStore;

	beforeEach(() => {
		vi.restoreAllMocks();
		store = new WorkspaceStore();
	});

	it('should initialize empty without selected workspace', () => {
		expect(store.workspaces).toHaveLength(0);
		expect(store.currentWorkspace).toBeNull();
		expect(store.hasWorkspaces).toBe(false);
	});

	it('should load workspaces and auto-select the first one', async () => {
		vi.spyOn(apiClient, 'listWorkspaces').mockResolvedValueOnce([
			{
				id: 'ws-1',
				name: 'Workspace One',
				slug: 'workspace-one',
				role: 'owner',
				created_at: '2026-01-01'
			},
			{
				id: 'ws-2',
				name: 'Workspace Two',
				slug: 'workspace-two',
				role: 'member',
				created_at: '2026-01-02'
			}
		]);

		await store.loadWorkspaces();

		expect(store.workspaces).toHaveLength(2);
		expect(store.hasWorkspaces).toBe(true);
		expect(store.currentWorkspace?.id).toBe('ws-1');
	});

	it('should create workspace, prepend to list, and select it', async () => {
		vi.spyOn(apiClient, 'createWorkspace').mockResolvedValueOnce({
			id: 'ws-new',
			name: 'Brand Agency',
			slug: 'brand-agency',
			role: 'owner',
			created_at: '2026-01-03'
		});

		const ws = await store.createWorkspace('Brand Agency');

		expect(ws).not.toBeNull();
		expect(store.workspaces).toHaveLength(1);
		expect(store.currentWorkspace?.name).toBe('Brand Agency');
	});

	it('should allow switching between workspaces', async () => {
		store.workspaces = [
			{ id: 'ws-1', name: 'Alpha', slug: 'alpha', role: 'owner', createdAt: '' },
			{ id: 'ws-2', name: 'Beta', slug: 'beta', role: 'admin', createdAt: '' }
		];

		store.selectWorkspace('ws-2');
		expect(store.currentWorkspace?.name).toBe('Beta');

		store.selectWorkspace('ws-1');
		expect(store.currentWorkspace?.name).toBe('Alpha');
	});
});
