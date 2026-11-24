<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import { Sparkles, ArrowRight, ShieldCheck, Zap } from 'lucide-svelte';
</script>

<main class="flex-1 flex flex-col items-center justify-center p-8 max-w-5xl mx-auto text-center space-y-8">
	<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full text-xs font-medium bg-muted border border-border text-muted-foreground">
		<Sparkles class="w-3.5 h-3.5 text-accent-foreground" />
		<span>Autonomous Multi-Tenant Digital Marketing Agent</span>
	</div>

	<h1 class="text-4xl md:text-6xl font-semibold tracking-tight text-foreground max-w-3xl leading-tight">
		Automate Posters, Ads & Social Campaigns at Scale.
	</h1>

	<p class="text-base md:text-lg text-muted-foreground max-w-2xl leading-relaxed">
		Generate on-brand promotional posters, launch paid Meta & Google ad campaigns, and schedule content across platforms with programmatic precision.
	</p>

	<div class="flex flex-wrap items-center justify-center gap-4 pt-4">
		<a href="/login">
			<Button variant="primary" size="lg">
				<span>Open Console</span>
				<ArrowRight class="w-4 h-4 ml-1.5" />
			</Button>
		</a>
		<a href="/register">
			<Button variant="secondary" size="lg">
				<span>Create Free Account</span>
			</Button>
		</a>
	</div>

	<div class="grid grid-cols-1 md:grid-cols-3 gap-6 w-full pt-12 border-t border-border/60 text-left">
		<div class="p-6 rounded-xl border border-border bg-card">
			<Zap class="w-5 h-5 text-primary mb-3" />
			<h3 class="font-semibold text-foreground text-sm mb-1">AI Creative Studio</h3>
			<p class="text-xs text-muted-foreground">Generate festival & daily posters with vector brand logos, colors, and auto-generated copy.</p>
		</div>

		<div class="p-6 rounded-xl border border-border bg-card">
			<ShieldCheck class="w-5 h-5 text-primary mb-3" />
			<h3 class="font-semibold text-foreground text-sm mb-1">Multi-Tenant Isolation</h3>
			<p class="text-xs text-muted-foreground">Agency-grade workspace separation. Manage dozens of client brand identities securely.</p>
		</div>

		<div class="p-6 rounded-xl border border-border bg-card">
			<Sparkles class="w-5 h-5 text-primary mb-3" />
			<h3 class="font-semibold text-foreground text-sm mb-1">Direct Ad Publishing</h3>
			<p class="text-xs text-muted-foreground">One-click push to Meta Marketing API, Google Ads, and WhatsApp Click-to-Chat campaigns.</p>
		</div>
	</div>
</main>
