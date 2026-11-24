import { apiClient } from '$lib/api/client';

export interface Workspace {
	id: string;
	name: string;
	slug: string;
	role: string;
	createdAt: string;
}

export class WorkspaceStore {
	workspaces = $state<Workspace[]>([]);
	currentWorkspace = $state<Workspace | null>(null);
	isLoading = $state<boolean>(false);
	errorMessage = $state<string | null>(null);

	hasWorkspaces = $derived<boolean>(this.workspaces.length > 0);

	constructor() {
		// Restore active workspace ID from localStorage if available
		if (typeof window !== 'undefined') {
			const saved = localStorage.getItem('dma_active_workspace');
			if (saved) {
				try {
					this.currentWorkspace = JSON.parse(saved);
					if (this.currentWorkspace) {
						apiClient.setWorkspaceId(this.currentWorkspace.id);
					}
				} catch {
					localStorage.removeItem('dma_active_workspace');
				}
			}
		}
	}

	async loadWorkspaces(): Promise<void> {
		this.isLoading = true;
		this.errorMessage = null;

		try {
			const list = await apiClient.listWorkspaces();
			this.workspaces = list.map((w) => ({
				id: w.id,
				name: w.name,
				slug: w.slug,
				role: w.role,
				createdAt: w.created_at
			}));

			// Auto-select first workspace if none currently selected
			if (!this.currentWorkspace && this.workspaces.length > 0) {
				this.selectWorkspace(this.workspaces[0].id);
			}
		} catch (err) {
			this.errorMessage = (err as Error).message;
		} finally {
			this.isLoading = false;
		}
	}

	selectWorkspace(id: string): void {
		const found = this.workspaces.find((w) => w.id === id);
		if (found) {
			this.currentWorkspace = found;
			apiClient.setWorkspaceId(found.id);
			if (typeof window !== 'undefined') {
				localStorage.setItem('dma_active_workspace', JSON.stringify(found));
			}
		}
	}

	async createWorkspace(name: string, slug?: string): Promise<Workspace | null> {
		this.isLoading = true;
		this.errorMessage = null;

		try {
			const created = await apiClient.createWorkspace(name, slug);
			const ws: Workspace = {
				id: created.id,
				name: created.name,
				slug: created.slug,
				role: created.role,
				createdAt: created.created_at
			};

			this.workspaces = [ws, ...this.workspaces];
			this.selectWorkspace(ws.id);
			return ws;
		} catch (err) {
			this.errorMessage = (err as Error).message;
			return null;
		} finally {
			this.isLoading = false;
		}
	}

	clear(): void {
		this.workspaces = [];
		this.currentWorkspace = null;
		this.errorMessage = null;
		apiClient.setWorkspaceId(null);
		if (typeof window !== 'undefined') {
			localStorage.removeItem('dma_active_workspace');
		}
	}
}

export const workspaceStore = new WorkspaceStore();
