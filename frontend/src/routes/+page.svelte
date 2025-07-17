<script lang="ts">
	import { GetRandomPokemon } from '$lib/wailsjs/go/app/App';
	import { AnimatedBackground } from '../custom/ui/animated-background/index.js';
	import { HeroSection } from '../custom/ui/hero-section/index.js';
	import { PokemonCard } from '../custom/ui/pokemon-card/index.js';
	import { WelcomeState } from '../custom/ui/welcome-state/index.js';
	import type { Pokemon } from '../custom/ui/shared/types.js';
	import { preloadImages } from '../custom/ui/shared/utils.js';

	// Svelte 5 runes with async support
	let pokemonPromise = $state<Promise<Pokemon> | null>(null);
	let animatedBackground: AnimatedBackground;

	async function fetchPokemonWithPreload(): Promise<Pokemon> {
		const pokemon = await GetRandomPokemon();
		
		// Collect all image URLs to preload
		const imageUrls = [pokemon.image_url];
		
		// Add evolution images if they exist
		if (pokemon.species?.evolutions) {
			pokemon.species.evolutions.forEach(evolution => {
				if (evolution.image_url) {
					imageUrls.push(evolution.image_url);
				}
			});
		}
		
		// Preload all images
		await preloadImages(imageUrls);
		
		return pokemon;
	}

	function fetchNewPokemon() {
		// Trigger particle animation in background component
		if (animatedBackground) {
			animatedBackground.triggerParticles();
		}

		pokemonPromise = fetchPokemonWithPreload();
	}

	function handlePokemonChange(pokemon: Pokemon) {
		// Trigger particle animation in background component
		if (animatedBackground) {
			animatedBackground.triggerParticles();
		}

		// Create a promise that preloads images before resolving
		pokemonPromise = (async () => {
			// Collect all image URLs to preload
			const imageUrls = [pokemon.image_url];
			
			// Add evolution images if they exist
			if (pokemon.species?.evolutions) {
				pokemon.species.evolutions.forEach(evolution => {
					if (evolution.image_url) {
						imageUrls.push(evolution.image_url);
					}
				});
			}
			
			// Preload all images
			await preloadImages(imageUrls);
			
			return pokemon;
		})();
	}
</script>

<!-- Main layout with animated background -->
<div class="relative flex h-screen flex-col overflow-hidden">
	<!-- Animated Background Component -->
	<AnimatedBackground bind:this={animatedBackground} />

	<main
		class="container relative mx-auto flex h-full max-w-7xl flex-col justify-center px-3 py-4 sm:px-4 sm:py-6 lg:px-6"
	>
		<!-- Hero Section Component -->
		<HeroSection onDiscoverClick={fetchNewPokemon} />

		<!-- Pokemon Display Area -->
		{#if pokemonPromise}
			{#await pokemonPromise}
				<div class="text-center">
					<PokemonCard pokemon={null} isLoading={true} />
					<p class="mt-2 text-sm text-gray-500 animate-pulse">Preparing your Pokémon...</p>
				</div>
			{:then pokemon}
				<PokemonCard
					pokemon={pokemon}
					isLoading={false}
					onPokemonChange={handlePokemonChange}
				/>
			{:catch error}
				<PokemonCard pokemon={null} isLoading={false} error={error.message} />
			{/await}
		{:else}
			<WelcomeState isVisible={true} />
		{/if}
	</main>
</div>
