<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Sparkles, Zap, Shuffle } from '@lucide/svelte';
	import { bounceOut } from 'svelte/easing';
	import { fade, fly, scale } from 'svelte/transition';
	import type { HeroSectionProps } from '../shared/types.js';

	let {
		title = 'PokéDiscover',
		subtitle = 'Embark on a magical journey to discover amazing Pokémon from across the universe',
		buttonText = 'Discover New Pokémon',
		onDiscoverClick
	}: HeroSectionProps = $props();

	let mounted = $state(false);
	let loading = $derived($effect.pending() > 0);

	$effect(() => {
		mounted = true;
	});
</script>

<!-- Hero Section -->
<div class="mb-6 text-center sm:mb-8">
	{#if mounted}
		<div class="mb-3 inline-flex items-center gap-1 sm:mb-4 sm:gap-2" in:scale={{ duration: 800, easing: bounceOut }}>
			<Sparkles class="h-4 w-4 text-pink-500 sm:h-6 sm:w-6" />
			<h1 class="bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text text-2xl font-bold text-transparent sm:text-4xl lg:text-5xl">
				{title}
			</h1>
			<Sparkles class="h-4 w-4 text-pink-500 sm:h-6 sm:w-6" />
		</div>
		<p class="mx-auto mb-4 max-w-2xl text-sm text-gray-600 sm:mb-6 sm:text-lg lg:text-xl" in:fade={{ delay: 300, duration: 600 }}>
			{subtitle}
		</p>

		<div class="relative inline-block" in:fly={{ y: 30, delay: 600, duration: 600, easing: bounceOut }}>
			<Button
				onclick={onDiscoverClick}
				class="relative transform rounded-full bg-gradient-to-r from-blue-500 to-purple-600 px-4 py-2 text-sm font-semibold text-white shadow-lg transition-all duration-300 hover:scale-105 hover:from-blue-600 hover:to-purple-700 hover:shadow-xl sm:px-6 sm:py-3 sm:text-lg {loading ? 'animate-pulse' : ''}"
				disabled={loading}
			>
				<div class="flex items-center gap-1 sm:gap-2">
					<Shuffle class="h-4 w-4 sm:h-5 sm:w-5 {loading ? 'animate-spin' : ''}" />
					<span class="hidden sm:inline">{loading ? 'Discovering...' : buttonText}</span>
					<span class="sm:hidden">{loading ? 'Finding...' : 'Discover'}</span>
					<Zap class="h-4 w-4 text-yellow-300 sm:h-5 sm:w-5" />
				</div>
			</Button>

			<div class="absolute inset-0 -z-10 rounded-full bg-gradient-to-r from-blue-500 to-purple-600 opacity-30 blur"></div>
		</div>
	{/if}
</div>