<script lang="ts">
	import RegisterForm from '$lib/components/auth/RegisterForm.svelte';
	import { workspaceStore } from '$lib/stores/workspace.svelte';
	import { goto } from '$app/navigation';
	import { Sparkles } from 'lucide-svelte';

	function handleSuccess() {
		if (workspaceStore.currentWorkspace) {
			goto(`/${workspaceStore.currentWorkspace.id}/overview`);
		} else {
			goto('/');
		}
	}
</script>

<div class="min-h-screen flex items-center justify-center p-6 bg-background">
	<div class="w-full max-w-sm space-y-6">
		<div class="text-center space-y-2">
			<div class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-medium bg-muted border border-border text-muted-foreground">
				<Sparkles class="w-3.5 h-3.5 text-foreground" />
				<span>Autonomous SaaS</span>
			</div>
			<h1 class="text-2xl font-semibold tracking-tight text-foreground">Create Account</h1>
			<p class="text-xs text-muted-foreground">Set up your multi-tenant workspace</p>
		</div>

		<div class="p-6 rounded-2xl border border-border bg-card shadow-xs">
			<RegisterForm onSuccess={handleSuccess} />
		</div>

		<p class="text-center text-xs text-muted-foreground">
			Already have an account?{' '}
			<a href="/login" class="text-foreground font-medium underline underline-offset-4 hover:text-primary">
				Sign In
			</a>
		</p>
	</div>
</div>
