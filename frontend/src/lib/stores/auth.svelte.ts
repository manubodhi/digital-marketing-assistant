import { apiClient } from '$lib/api/client';

export interface User {
	id: string;
	email: string;
	fullName: string;
	isEmailVerified: boolean;
	totpEnabled: boolean;
}

export class AuthStore {
	user = $state<User | null>(null);
	accessToken = $state<string | null>(null);
	isLoading = $state<boolean>(false);
	errorMessage = $state<string | null>(null);

	isAuthenticated = $derived<boolean>(this.user !== null && this.accessToken !== null);

	constructor() {
		// Restore token from localStorage if in browser
		if (typeof window !== 'undefined') {
			const savedToken = localStorage.getItem('dma_access_token');
			const savedUser = localStorage.getItem('dma_user');
			if (savedToken && savedUser) {
				try {
					this.accessToken = savedToken;
					this.user = JSON.parse(savedUser);
					apiClient.setToken(savedToken);
				} catch {
					this.logout();
				}
			}
		}
	}

	async register(email: string, password: string, fullName: string): Promise<boolean> {
		this.isLoading = true;
		this.errorMessage = null;

		try {
			const res = await apiClient.register(email, password, fullName);
			this.setSession(res.access_token, {
				id: res.user.id,
				email: res.user.email,
				fullName: res.user.full_name,
				isEmailVerified: res.user.is_email_verified,
				totpEnabled: res.user.totp_enabled
			});
			return true;
		} catch (err) {
			this.errorMessage = (err as Error).message;
			return false;
		} finally {
			this.isLoading = false;
		}
	}

	async login(email: string, password: string): Promise<boolean> {
		this.isLoading = true;
		this.errorMessage = null;

		try {
			const res = await apiClient.login(email, password);
			this.setSession(res.access_token, {
				id: res.user.id,
				email: res.user.email,
				fullName: res.user.full_name,
				isEmailVerified: res.user.is_email_verified,
				totpEnabled: res.user.totp_enabled
			});
			return true;
		} catch (err) {
			this.errorMessage = (err as Error).message;
			return false;
		} finally {
			this.isLoading = false;
		}
	}

	logout(): void {
		this.user = null;
		this.accessToken = null;
		this.errorMessage = null;
		apiClient.setToken(null);
		if (typeof window !== 'undefined') {
			localStorage.removeItem('dma_access_token');
			localStorage.removeItem('dma_user');
		}
	}

	private setSession(token: string, user: User): void {
		this.accessToken = token;
		this.user = user;
		apiClient.setToken(token);
		if (typeof window !== 'undefined') {
			localStorage.setItem('dma_access_token', token);
			localStorage.setItem('dma_user', JSON.stringify(user));
		}
	}
}

export const authStore = new AuthStore();
