<script lang="ts">
	import { GetRandomPokemon } from '$lib/wailsjs/go/app/App';
	import { AnimatedBackground } from '../custom/ui/animated-background/index.js';
	import { HeroSection } from '../custom/ui/hero-section/index.js';
	import { PokemonCard } from '../custom/ui/pokemon-card/index.js';
	import { WelcomeState } from '../custom/ui/welcome-state/index.js';
	import type { Pokemon } from '../custom/ui/shared/types.js';
	import { preloadImage } from '../custom/ui/shared/utils.js';

	// Svelte 5 runes
	let pokemonId = $state(0);
	let isAnimating = $state(false);
	let animatedBackground: AnimatedBackground;

	function fetchNewPokemon() {
		isAnimating = true;
		pokemonId++;

		// Trigger particle animation in background component
		if (animatedBackground) {
			animatedBackground.triggerParticles();
		}

		setTimeout(() => {
			isAnimating = false;
		}, 600);
	}

	async function getPokemon(): Promise<Pokemon | null> {
		if (pokemonId === 0) return null;
		const pokemon = await GetRandomPokemon();

		// Wait for image to load before resolving
		if (pokemon?.image_url) {
			await preloadImage(pokemon.image_url);
		}

		return pokemon;
	}

	// Debug inspection
	$inspect('Pokemon ID:', pokemonId);
	$inspect('Is Animating:', isAnimating);
</script>



<!-- Main layout with animated background -->
<div class="relative flex h-screen flex-col overflow-hidden">
	<!-- Animated Background Component -->
	<AnimatedBackground bind:this={animatedBackground} />

	<main class="container relative mx-auto flex h-full max-w-7xl flex-col justify-center px-3 py-4 sm:px-4 sm:py-6 lg:px-6">
		<!-- Hero Section Component -->
		<HeroSection onDiscoverClick={fetchNewPokemon} isLoading={isAnimating} />

		<!-- Pokemon Display Area -->
		{#if pokemonId > 0}
			{#await getPokemon()}
				<PokemonCard pokemon={null} isLoading={true} />
			{:then pokemon}
				<PokemonCard pokemon={pokemon} isLoading={false} />
			{:catch error}
				<PokemonCard pokemon={null} isLoading={false} error={error.message || 'Unknown error'} />
			{/await}
		{:else}
			<!-- Welcome State Component -->
			<WelcomeState isVisible={true} />
		{/if}
	</main>
</div>