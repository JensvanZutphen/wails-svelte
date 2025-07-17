<script lang="ts">
	import { Star } from '@lucide/svelte';
	import { bounceOut } from 'svelte/easing';
	import { scale } from 'svelte/transition';
	import type { WelcomeStateProps } from '../shared/types.js';

	let { isVisible }: WelcomeStateProps = $props();

	// Svelte 5 runes
	let mounted = $state(false);

	// Effect to handle mounting
	$effect(() => {
		mounted = true;
	});
</script>

{#if isVisible && mounted}
	<!-- Welcome State -->
	<div class="mx-auto w-full max-w-xs text-center sm:max-w-sm" in:scale={{ delay: 900, duration: 600, easing: bounceOut }}>
		<div class="rounded-2xl border border-white/20 bg-white/60 p-4 shadow-xl backdrop-blur-sm sm:rounded-3xl sm:p-6">
			<div class="mb-3 sm:mb-4">
				<Star class="mx-auto h-8 w-8 animate-bounce fill-current text-yellow-400 sm:h-12 sm:w-12" />
			</div>
			<h3 class="mb-2 text-lg font-bold text-gray-800 sm:text-xl">Ready for Adventure?</h3>
			<p class="text-sm text-gray-600 sm:text-base">
				Click the button above to discover your first Pokémon and begin your journey!
			</p>
		</div>
	</div>
{/if}