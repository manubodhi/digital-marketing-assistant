<script lang="ts">
	import LoginForm from '$lib/components/auth/LoginForm.svelte';
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
				<span>Marketing Co-Pilot</span>
			</div>
			<h1 class="text-2xl font-semibold tracking-tight text-foreground">Welcome Back</h1>
			<p class="text-xs text-muted-foreground">Sign in to your digital marketing assistant console</p>
		</div>

		<div class="p-6 rounded-2xl border border-border bg-card shadow-xs">
			<LoginForm onSuccess={handleSuccess} />
		</div>

		<p class="text-center text-xs text-muted-foreground">
			Don't have an account?{' '}
			<a href="/register" class="text-foreground font-medium underline underline-offset-4 hover:text-primary">
				Sign Up
			</a>
		</p>
	</div>
</div>
