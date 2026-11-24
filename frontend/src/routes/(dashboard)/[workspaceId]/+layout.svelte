<script lang="ts">
	import WorkspaceSwitcher from '$lib/components/workspace/WorkspaceSwitcher.svelte';
	import { authStore } from '$lib/stores/auth.svelte';
	import { workspaceStore } from '$lib/stores/workspace.svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import {
		LayoutDashboard,
		Palette,
		Megaphone,
		CalendarDays,
		LogOut,
		Sparkles
	} from 'lucide-svelte';

	let { children } = $props();

	let currentPath = $derived($page.url.pathname);
	let wsId = $derived(workspaceStore.currentWorkspace?.id || $page.params.workspaceId);

	function handleLogout() {
		authStore.logout();
		workspaceStore.clear();
		goto('/login');
	}
</script>

<div class="min-h-screen flex bg-background text-foreground">
	<!-- Left Sidebar -->
	<aside class="w-60 border-r border-border bg-card flex flex-col shrink-0">
		<div class="h-14 px-4 flex items-center gap-2 border-b border-border">
			<div class="w-7 h-7 rounded-lg bg-primary text-primary-foreground flex items-center justify-center font-bold text-xs">
				<Sparkles class="w-4 h-4" />
			</div>
			<span class="font-semibold text-sm tracking-tight">MarketingAI</span>
		</div>

		<!-- Nav Links -->
		<nav class="flex-1 p-3 space-y-1">
			<a
				href="/{wsId}/overview"
				class="flex items-center gap-2.5 px-3 py-2 rounded-lg text-xs font-medium transition-colors {currentPath.includes('/overview')
					? 'bg-muted text-foreground'
					: 'text-muted-foreground hover:bg-muted/50 hover:text-foreground'}"
			>
				<LayoutDashboard class="w-4 h-4" />
				<span>Overview</span>
			</a>

			<a
				href="/{wsId}/brand-kit"
				class="flex items-center gap-2.5 px-3 py-2 rounded-lg text-xs font-medium transition-colors {currentPath.includes('/brand-kit')
					? 'bg-muted text-foreground'
					: 'text-muted-foreground hover:bg-muted/50 hover:text-foreground'}"
			>
				<Palette class="w-4 h-4" />
				<span>Brand Kit</span>
			</a>

			<a
				href="/{wsId}/campaigns"
				class="flex items-center gap-2.5 px-3 py-2 rounded-lg text-xs font-medium transition-colors {currentPath.includes('/campaigns')
					? 'bg-muted text-foreground'
					: 'text-muted-foreground hover:bg-muted/50 hover:text-foreground'}"
			>
				<Megaphone class="w-4 h-4" />
				<span>Paid Ads</span>
			</a>

			<a
				href="/{wsId}/calendar"
				class="flex items-center gap-2.5 px-3 py-2 rounded-lg text-xs font-medium transition-colors {currentPath.includes('/calendar')
					? 'bg-muted text-foreground'
					: 'text-muted-foreground hover:bg-muted/50 hover:text-foreground'}"
			>
				<CalendarDays class="w-4 h-4" />
				<span>Social Calendar</span>
			</a>
		</nav>

		<!-- User Footer -->
		<div class="p-3 border-t border-border flex items-center justify-between">
			<div class="flex flex-col truncate">
				<span class="text-xs font-medium truncate">{authStore.user?.fullName ?? 'Marketer'}</span>
				<span class="text-[10px] text-muted-foreground truncate">{authStore.user?.email ?? ''}</span>
			</div>
			<button
				type="button"
				onclick={handleLogout}
				title="Sign Out"
				class="p-1.5 rounded-md hover:bg-muted text-muted-foreground hover:text-destructive cursor-pointer transition-colors"
			>
				<LogOut class="w-4 h-4" />
			</button>
		</div>
	</aside>

	<!-- Main Stage -->
	<div class="flex-1 flex flex-col min-w-0">
		<!-- Top Bar -->
		<header class="h-14 px-6 border-b border-border bg-card flex items-center justify-between shrink-0">
			<div class="flex items-center gap-4">
				<WorkspaceSwitcher />
			</div>

			<div class="flex items-center gap-2 text-xs text-muted-foreground">
				<span class="inline-block w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
				<span>Connected</span>
			</div>
		</header>

		<!-- Content View -->
		<main class="flex-1 p-8 overflow-y-auto max-w-7xl">
			{@render children()}
		</main>
	</div>
</div>
