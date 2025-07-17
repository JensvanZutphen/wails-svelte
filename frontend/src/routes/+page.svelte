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
	let currentPokemon = $state<Pokemon | null>(null);
	let pokemonError = $state<string | undefined>(undefined);
	let animatedBackground: AnimatedBackground;

	async function fetchNewPokemon() {
		isAnimating = true;
		pokemonError = undefined;

		// Trigger particle animation in background component
		if (animatedBackground) {
			animatedBackground.triggerParticles();
		}

		try {
			const pokemon = await GetRandomPokemon();
			currentPokemon = pokemon;
			pokemonId++;
		} catch (error) {
			pokemonError = error instanceof Error ? error.message : 'Unknown error occurred';
			currentPokemon = null;
		} finally {
			setTimeout(() => {
				isAnimating = false;
			}, 600);
		}
	}

	// Debug inspection
	$inspect('Pokemon ID:', pokemonId);
	$inspect('Is Animating:', isAnimating);
	$inspect('Current Pokemon:', currentPokemon);
	$inspect('Pokemon Error:', pokemonError);
	$inspect('Should Show Pokemon Card:', pokemonId > 0 || currentPokemon);
	$inspect('Should Show Welcome:', !(pokemonId > 0 || currentPokemon));
</script>



<!-- Main layout with animated background -->
<div class="relative flex h-screen flex-col overflow-hidden">
	<!-- Animated Background Component -->
	<AnimatedBackground bind:this={animatedBackground} />

	<main class="container relative mx-auto flex h-full max-w-7xl flex-col justify-center px-3 py-4 sm:px-4 sm:py-6 lg:px-6">
		<!-- Hero Section Component -->
		<HeroSection onDiscoverClick={fetchNewPokemon} isLoading={isAnimating} />

		<!-- Pokemon Display Area -->
		{#if currentPokemon && !isAnimating}
			<PokemonCard 
				pokemon={currentPokemon} 
				isLoading={false} 
				error={pokemonError} 
			/>
		{:else if isAnimating}
			<PokemonCard 
				pokemon={null} 
				isLoading={true} 
				error={undefined} 
			/>
		{:else}
			<WelcomeState isVisible={true} />
		{/if}
	</main>
</div>