<script lang="ts">
	interface Props {
		type?: string;
		label?: string;
		placeholder?: string;
		value?: string;
		error?: string;
		required?: boolean;
		disabled?: boolean;
		oninput?: (event: Event) => void;
	}

	let {
		type = 'text',
		label,
		placeholder,
		value = $bindable(''),
		error,
		required = false,
		disabled = false,
		oninput
	}: Props = $props();
</script>

<div class="flex flex-col space-y-1.5 w-full text-left">
	{#if label}
		<label class="text-xs font-medium text-foreground">
			{label}
			{#if required}<span class="text-destructive">*</span>{/if}
		</label>
	{/if}

	<input
		{type}
		{placeholder}
		bind:value
		{required}
		{disabled}
		{oninput}
		class="w-full px-3 py-2 text-sm bg-card border rounded-lg transition-colors focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring disabled:opacity-50 disabled:cursor-not-allowed {error
			? 'border-destructive focus-visible:ring-destructive'
			: 'border-input hover:border-border'}"
	/>

	{#if error}
		<span class="text-xs text-destructive">{error}</span>
	{/if}
</div>
