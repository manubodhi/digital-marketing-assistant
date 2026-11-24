<script lang="ts">
	import { authStore } from '$lib/stores/auth.svelte';
	import { workspaceStore } from '$lib/stores/workspace.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Input from '$lib/components/ui/Input.svelte';

	interface Props {
		onSuccess?: () => void;
	}

	let { onSuccess }: Props = $props();

	let email = $state('');
	let password = $state('');

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		const success = await authStore.login(email, password);
		if (success) {
			await workspaceStore.loadWorkspaces();
			if (onSuccess) {
				onSuccess();
			}
		}
	}
</script>

<form onsubmit={handleSubmit} class="space-y-4 w-full">
	{#if authStore.errorMessage}
		<div class="p-3 text-xs rounded-lg bg-destructive/10 text-destructive border border-destructive/20">
			{authStore.errorMessage}
		</div>
	{/if}

	<Input
		type="email"
		label="Work Email"
		placeholder="name@agency.com"
		bind:value={email}
		required
	/>

	<Input
		type="password"
		label="Password"
		placeholder="••••••••"
		bind:value={password}
		required
	/>

	<Button type="submit" variant="primary" size="lg" loading={authStore.isLoading} disabled={authStore.isLoading}>
		<span>Sign In</span>
	</Button>
</form>
