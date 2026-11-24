<script lang="ts">
	import { workspaceStore } from '$lib/stores/workspace.svelte';
	import { authStore } from '$lib/stores/auth.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { Sparkles, Palette, Megaphone, TrendingUp } from 'lucide-svelte';
</script>

<div class="space-y-8">
	<!-- Header -->
	<div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
		<div>
			<h1 class="text-2xl font-semibold tracking-tight text-foreground">
				Workspace Overview
			</h1>
			<p class="text-xs text-muted-foreground mt-1">
				Active tenant: <span class="font-medium text-foreground">{workspaceStore.currentWorkspace?.name ?? 'Loading...'}</span>
				({workspaceStore.currentWorkspace?.role ?? 'member'})
			</p>
		</div>

		<div class="flex items-center gap-3">
			<Button variant="outline" size="sm">
				<span>Export Report</span>
			</Button>
			<Button variant="primary" size="sm">
				<Sparkles class="w-3.5 h-3.5 mr-1.5" />
				<span>Create Poster</span>
			</Button>
		</div>
	</div>

	<!-- Metric Cards Grid -->
	<div class="grid grid-cols-1 md:grid-cols-4 gap-4">
		<div class="p-5 rounded-xl border border-border bg-card shadow-xs">
			<span class="text-xs text-muted-foreground font-medium">Brand Identity Status</span>
			<div class="mt-2 flex items-baseline justify-between">
				<span class="text-xl font-semibold text-foreground">Ready</span>
				<span class="text-xs text-emerald-500 font-medium">100% Configured</span>
			</div>
		</div>

		<div class="p-5 rounded-xl border border-border bg-card shadow-xs">
			<span class="text-xs text-muted-foreground font-medium">Active Ad Campaigns</span>
			<div class="mt-2 flex items-baseline justify-between">
				<span class="text-xl font-semibold text-foreground">0</span>
				<span class="text-xs text-muted-foreground">Draft mode</span>
			</div>
		</div>

		<div class="p-5 rounded-xl border border-border bg-card shadow-xs">
			<span class="text-xs text-muted-foreground font-medium">Scheduled Posts</span>
			<div class="mt-2 flex items-baseline justify-between">
				<span class="text-xl font-semibold text-foreground">0</span>
				<span class="text-xs text-muted-foreground">This week</span>
			</div>
		</div>

		<div class="p-5 rounded-xl border border-border bg-card shadow-xs">
			<span class="text-xs text-muted-foreground font-medium">Leads Captured</span>
			<div class="mt-2 flex items-baseline justify-between">
				<span class="text-xl font-semibold text-foreground">0</span>
				<span class="text-xs text-muted-foreground">WhatsApp / Forms</span>
			</div>
		</div>
	</div>

	<!-- Quick Action Banner -->
	<div class="p-6 rounded-2xl border border-border bg-card flex flex-col md:flex-row items-start md:items-center justify-between gap-6">
		<div class="space-y-1">
			<h3 class="text-base font-semibold text-foreground">Next Step: Configure Brand Kit</h3>
			<p class="text-xs text-muted-foreground max-w-xl">
				Set up your client's logo, primary and secondary brand hex colors, and brand voice to start generating promotional posters automatically.
			</p>
		</div>

		<a href="/{workspaceStore.currentWorkspace?.id}/brand-kit">
			<Button variant="primary" size="md">
				<Palette class="w-4 h-4 mr-1.5" />
				<span>Configure Brand Kit</span>
			</Button>
		</a>
	</div>
</div>
