<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		type?: 'button' | 'submit' | 'reset';
		variant?: 'primary' | 'secondary' | 'destructive' | 'outline' | 'ghost';
		size?: 'sm' | 'md' | 'lg';
		disabled?: boolean;
		loading?: boolean;
		onclick?: (event: MouseEvent) => void;
		children?: Snippet;
	}

	let {
		type = 'button',
		variant = 'primary',
		size = 'md',
		disabled = false,
		loading = false,
		onclick,
		children
	}: Props = $props();

	const variantClasses = {
		primary: 'bg-primary text-primary-foreground hover:bg-primary/90 shadow-xs',
		secondary: 'bg-secondary text-secondary-foreground hover:bg-secondary/80 border border-border',
		destructive: 'bg-destructive text-destructive-foreground hover:bg-destructive/90',
		outline: 'border border-border bg-transparent hover:bg-muted text-foreground',
		ghost: 'hover:bg-muted text-foreground'
	};

	const sizeClasses = {
		sm: 'h-8 px-3 text-xs rounded-md',
		md: 'h-9 px-4 text-sm rounded-lg',
		lg: 'h-11 px-6 text-sm font-medium rounded-lg'
	};
</script>

<button
	{type}
	disabled={disabled || loading}
	{onclick}
	class="inline-flex items-center justify-center font-medium transition-colors focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring disabled:opacity-50 disabled:pointer-events-none cursor-pointer {variantClasses[variant]} {sizeClasses[size]}"
>
	{#if loading}
		<svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-current" fill="none" viewBox="0 0 24 24">
			<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
			<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
		</svg>
	{/if}
	{#if children}
		{@render children()}
	{/if}
</button>
