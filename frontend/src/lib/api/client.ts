import type { paths } from './types';

export class APIClient {
	private baseUrl: string;
	private token: string | null = null;
	private workspaceId: string | null = null;

	constructor(baseUrl = 'http://localhost:8080/api/v1') {
		this.baseUrl = baseUrl;
	}

	setToken(token: string | null) {
		this.token = token;
	}

	setWorkspaceId(workspaceId: string | null) {
		this.workspaceId = workspaceId;
	}

	private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
		const headers: Record<string, string> = {
			'Content-Type': 'application/json',
			...(options.headers as Record<string, string>)
		};

		if (this.token) {
			headers['Authorization'] = `Bearer ${this.token}`;
		}

		if (this.workspaceId) {
			headers['X-Workspace-ID'] = this.workspaceId;
		}

		const res = await fetch(`${this.baseUrl}${endpoint}`, {
			...options,
			headers
		});

		const json = await res.json();

		if (!res.ok || !json.success) {
			throw new Error(json.error?.message || `Request failed with status ${res.status}`);
		}

		return json.data as T;
	}

	async register(email: string, password: string, fullName: string) {
		return this.request<{
			access_token: string;
			user: {
				id: string;
				email: string;
				full_name: string;
				is_email_verified: boolean;
				totp_enabled: boolean;
			};
		}>('/auth/register', {
			method: 'POST',
			body: JSON.stringify({ email, password, fullName })
		});
	}

	async login(email: string, password: string) {
		return this.request<{
			access_token: string;
			user: {
				id: string;
				email: string;
				full_name: string;
				is_email_verified: boolean;
				totp_enabled: boolean;
			};
		}>('/auth/login', {
			method: 'POST',
			body: JSON.stringify({ email, password })
		});
	}

	async listWorkspaces() {
		return this.request<
			Array<{
				id: string;
				name: string;
				slug: string;
				role: string;
				created_at: string;
			}>
		>('/workspaces');
	}

	async createWorkspace(name: string, slug?: string) {
		return this.request<{
			id: string;
			name: string;
			slug: string;
			role: string;
			created_at: string;
		}>('/workspaces', {
			method: 'POST',
			body: JSON.stringify({ name, slug: slug || '' })
		});
	}
}

export const apiClient = new APIClient();
