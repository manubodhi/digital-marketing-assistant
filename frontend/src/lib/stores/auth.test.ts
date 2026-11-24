import { describe, it, expect, beforeEach, vi } from 'vitest';
import { AuthStore } from './auth.svelte';
import { apiClient } from '$lib/api/client';

describe('AuthStore (Svelte 5 Runes)', () => {
	let store: AuthStore;

	beforeEach(() => {
		vi.restoreAllMocks();
		store = new AuthStore();
	});

	it('should initialize with null session and not authenticated', () => {
		expect(store.user).toBeNull();
		expect(store.accessToken).toBeNull();
		expect(store.isAuthenticated).toBe(false);
		expect(store.isLoading).toBe(false);
	});

	it('should successfully login and set user session', async () => {
		vi.spyOn(apiClient, 'login').mockResolvedValueOnce({
			access_token: 'mock-jwt-token',
			user: {
				id: 'mock-uuid',
				email: 'test@example.com',
				full_name: 'Test Marketer',
				is_email_verified: true,
				totp_enabled: false
			}
		});

		const success = await store.login('test@example.com', 'password123');

		expect(success).toBe(true);
		expect(store.isAuthenticated).toBe(true);
		expect(store.user?.email).toBe('test@example.com');
		expect(store.accessToken).toBe('mock-jwt-token');
		expect(store.errorMessage).toBeNull();
	});

	it('should handle login error gracefully', async () => {
		vi.spyOn(apiClient, 'login').mockRejectedValueOnce(new Error('Invalid email or password'));

		const success = await store.login('test@example.com', 'wrongpassword');

		expect(success).toBe(false);
		expect(store.isAuthenticated).toBe(false);
		expect(store.errorMessage).toBe('Invalid email or password');
		expect(store.isLoading).toBe(false);
	});

	it('should clear session on logout', async () => {
		vi.spyOn(apiClient, 'login').mockResolvedValueOnce({
			access_token: 'mock-jwt-token',
			user: {
				id: 'mock-uuid',
				email: 'test@example.com',
				full_name: 'Test Marketer',
				is_email_verified: true,
				totp_enabled: false
			}
		});

		await store.login('test@example.com', 'password123');
		expect(store.isAuthenticated).toBe(true);

		store.logout();
		expect(store.isAuthenticated).toBe(false);
		expect(store.user).toBeNull();
		expect(store.accessToken).toBeNull();
	});
});
