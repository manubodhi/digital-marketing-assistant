<script lang="ts">
	import { workspaceStore } from '$lib/stores/workspace.svelte';
	import { Building2, ChevronDown, Plus, Check } from 'lucide-svelte';

	let isOpen = $state(false);
	let isCreating = $state(false);
	let newWorkspaceName = $state('');

	function toggleDropdown() {
		isOpen = !isOpen;
	}

	function handleSelect(id: string) {
		workspaceStore.selectWorkspace(id);
		isOpen = false;
	}

	async function handleCreate(e: SubmitEvent) {
		e.preventDefault();
		if (!newWorkspaceName.trim()) return;
		const created = await workspaceStore.createWorkspace(newWorkspaceName);
		if (created) {
			newWorkspaceName = '';
			isCreating = false;
			isOpen = false;
		}
	}
</script>

<div class="relative inline-block text-left">
	<button
		type="button"
		onclick={toggleDropdown}
		class="inline-flex items-center gap-2 px-3 py-1.5 text-xs font-medium rounded-lg border border-border bg-card hover:bg-muted text-foreground transition-colors cursor-pointer"
	>
		<Building2 class="w-3.5 h-3.5 text-muted-foreground" />
		<span class="max-w-[140px] truncate">
			{workspaceStore.currentWorkspace?.name ?? 'Select Workspace'}
		</span>
		<ChevronDown class="w-3.5 h-3.5 text-muted-foreground" />
	</button>

	{#if isOpen}
		<div
			class="absolute left-0 mt-1.5 w-60 rounded-xl border border-border bg-popover text-popover-foreground shadow-lg z-50 p-1.5 space-y-1"
		>
			<div class="px-2 py-1 text-[10px] font-semibold uppercase tracking-wider text-muted-foreground">
				Workspaces
			</div>

			<div class="max-h-48 overflow-y-auto space-y-0.5">
				{#each workspaceStore.workspaces as ws (ws.id)}
					<button
						type="button"
						onclick={() => handleSelect(ws.id)}
						class="w-full flex items-center justify-between px-2.5 py-1.5 text-xs rounded-md text-left transition-colors hover:bg-muted cursor-pointer {ws.id ===
						workspaceStore.currentWorkspace?.id
							? 'bg-muted font-medium text-foreground'
							: 'text-muted-foreground'}"
					>
						<span class="truncate">{ws.name}</span>
						{#if ws.id === workspaceStore.currentWorkspace?.id}
							<Check class="w-3.5 h-3.5 text-primary shrink-0" />
						{/if}
					</button>
				{/each}
			</div>

			<div class="border-t border-border pt-1">
				{#if !isCreating}
					<button
						type="button"
						onclick={() => (isCreating = true)}
						class="w-full flex items-center gap-2 px-2.5 py-1.5 text-xs rounded-md text-muted-foreground hover:text-foreground hover:bg-muted transition-colors cursor-pointer"
					>
						<Plus class="w-3.5 h-3.5" />
						<span>New Workspace</span>
					</button>
				{:else}
					<form onsubmit={handleCreate} class="p-1 space-y-2">
						<input
							type="text"
							placeholder="Workspace name"
							bind:value={newWorkspaceName}
							class="w-full px-2 py-1 text-xs bg-card border border-input rounded focus-visible:outline-hidden focus-visible:ring-1 focus-visible:ring-ring"
							required
						/>
						<div class="flex items-center justify-end gap-1.5">
							<button
								type="button"
								onclick={() => (isCreating = false)}
								class="px-2 py-1 text-[10px] rounded hover:bg-muted text-muted-foreground cursor-pointer"
							>
								Cancel
							</button>
							<button
								type="submit"
								class="px-2 py-1 text-[10px] bg-primary text-primary-foreground rounded hover:bg-primary/90 cursor-pointer font-medium"
							>
								Create
							</button>
						</div>
					</form>
				{/if}
			</div>
		</div>
	{/if}
</div>
